package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ResourceReference represents a reference between resources
type ResourceReference struct {
	SourceType      string `yaml:"sourceType"`
	SourceName      string `yaml:"sourceName"`
	SourceNamespace string `yaml:"sourceNamespace"`
	TargetType      string `yaml:"targetType"`
	TargetName      string `yaml:"targetName"`
	TargetNamespace string `yaml:"targetNamespace"`
	Field           string `yaml:"field"`
	Line            int    `yaml:"line"`
}

// ValidationResult represents the result of cross-reference validation
type ValidationResult struct {
	ValidReferences   []ResourceReference `yaml:"validReferences"`
	BrokenReferences  []ResourceReference `yaml:"brokenReferences"`
	WarningReferences []ResourceReference `yaml:"warningReferences"`
	Summary           ValidationSummary   `yaml:"summary"`
}

// ValidationSummary provides a summary of validation results
type ValidationSummary struct {
	TotalReferences   int `yaml:"totalReferences"`
	ValidReferences   int `yaml:"validReferences"`
	BrokenReferences  int `yaml:"brokenReferences"`
	WarningReferences int `yaml:"warningReferences"`
}

// ResourceValidator handles cross-reference validation
type ResourceValidator struct {
	outputDir string
	resources map[string]map[string]map[string]interface{}
}

// NewResourceValidator creates a new validator instance
func NewResourceValidator(outputDir string) *ResourceValidator {
	return &ResourceValidator{
		outputDir: outputDir,
		resources: make(map[string]map[string]map[string]interface{}),
	}
}

// Validate performs cross-reference validation on all exported resources
func (v *ResourceValidator) Validate() (*ValidationResult, error) {
	// Load all resources
	if err := v.loadResources(); err != nil {
		return nil, fmt.Errorf("failed to load resources: %w", err)
	}

	result := &ValidationResult{
		ValidReferences:   []ResourceReference{},
		BrokenReferences:  []ResourceReference{},
		WarningReferences: []ResourceReference{},
	}

	// Validate different types of references
	v.validateServiceReferences(result)
	v.validateRoleBindingReferences(result)
	v.validateNetworkPolicyReferences(result)
	v.validateIngressReferences(result)
	v.validateHPAReferences(result)
	v.validatePDBReferences(result)

	// Calculate summary
	result.Summary = ValidationSummary{
		TotalReferences:   len(result.ValidReferences) + len(result.BrokenReferences) + len(result.WarningReferences),
		ValidReferences:   len(result.ValidReferences),
		BrokenReferences:  len(result.BrokenReferences),
		WarningReferences: len(result.WarningReferences),
	}

	return result, nil
}

// loadResources loads all YAML resources from the output directory
func (v *ResourceValidator) loadResources() error {
	return filepath.Walk(v.outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-YAML files and directories
		if info.IsDir() || !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Skip reports directory
		if strings.Contains(path, "kalco-reports") {
			return nil
		}

		// Parse the file path to extract namespace, resource type, and filename
		relPath, err := filepath.Rel(v.outputDir, path)
		if err != nil {
			return nil // Skip files we can't parse
		}

		parts := strings.Split(relPath, string(os.PathSeparator))
		if len(parts) < 3 {
			return nil // Skip files not in expected structure
		}

		namespace := parts[0]
		resourceType := parts[1]
		filename := parts[2]

		// Load the YAML content
		content, err := os.ReadFile(path)
		if err != nil {
			return nil // Skip files we can't read
		}

		var resource interface{}
		if err := yaml.Unmarshal(content, &resource); err != nil {
			return nil // Skip files we can't parse
		}

		// Store the resource
		if v.resources[namespace] == nil {
			v.resources[namespace] = make(map[string]map[string]interface{})
		}
		if v.resources[namespace][resourceType] == nil {
			v.resources[namespace][resourceType] = make(map[string]interface{})
		}

		// Extract resource name from filename
		resourceName := strings.TrimSuffix(filename, ".yaml")
		v.resources[namespace][resourceType][resourceName] = resource

		return nil
	})
}

// validateServiceReferences validates Service selector references
func (v *ResourceValidator) validateServiceReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		services, ok := resources["Service"]
		if !ok {
			continue
		}

		for serviceName, serviceData := range services {
			service := serviceData.(map[string]interface{})
			spec, ok := service["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			selector, ok := spec["selector"].(map[string]interface{})
			if !ok {
				continue
			}

			// Check if selector targets exist
			for key, value := range selector {
				if key == "app" || key == "name" || key == "component" {
					// Look for matching pods/deployments
					if !v.selectorTargetExists(namespace, key, value.(string)) {
						result.BrokenReferences = append(result.BrokenReferences, ResourceReference{
							SourceType:      "Service",
							SourceName:      serviceName,
							SourceNamespace: namespace,
							TargetType:      "Pod/Deployment",
							TargetName:      value.(string),
							TargetNamespace: namespace,
							Field:           fmt.Sprintf("spec.selector.%s", key),
							Line:            0, // We don't track line numbers in this implementation
						})
					} else {
						result.ValidReferences = append(result.ValidReferences, ResourceReference{
							SourceType:      "Service",
							SourceName:      serviceName,
							SourceNamespace: namespace,
							TargetType:      "Pod/Deployment",
							TargetName:      value.(string),
							TargetNamespace: namespace,
							Field:           fmt.Sprintf("spec.selector.%s", key),
							Line:            0,
						})
					}
				}
			}
		}
	}
}

// validateRoleBindingReferences validates RoleBinding subject references
func (v *ResourceValidator) validateRoleBindingReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		roleBindings, ok := resources["RoleBinding"]
		if !ok {
			continue
		}

		for rbName, rbData := range roleBindings {
			rb := rbData.(map[string]interface{})
			subjects, ok := rb["subjects"].([]interface{})
			if !ok {
				continue
			}

			for _, subject := range subjects {
				subj := subject.(map[string]interface{})
				kind, _ := subj["kind"].(string)
				name, _ := subj["name"].(string)
				subjNamespace, _ := subj["namespace"].(string)

				if subjNamespace == "" {
					subjNamespace = namespace
				}

				// Validate ServiceAccount references
				if kind == "ServiceAccount" {
					if !v.resourceExists(subjNamespace, "ServiceAccount", name) {
						result.BrokenReferences = append(result.BrokenReferences, ResourceReference{
							SourceType:      "RoleBinding",
							SourceName:      rbName,
							SourceNamespace: namespace,
							TargetType:      "ServiceAccount",
							TargetName:      name,
							TargetNamespace: subjNamespace,
							Field:           "subjects",
							Line:            0,
						})
					} else {
						result.ValidReferences = append(result.ValidReferences, ResourceReference{
							SourceType:      "RoleBinding",
							SourceName:      rbName,
							SourceNamespace: namespace,
							TargetType:      "ServiceAccount",
							TargetName:      name,
							TargetNamespace: subjNamespace,
							Field:           "subjects",
							Line:            0,
						})
					}
				}

				// Validate User/Group references (these are external, so warnings)
				if kind == "User" || kind == "Group" {
					result.WarningReferences = append(result.WarningReferences, ResourceReference{
						SourceType:      "RoleBinding",
						SourceName:      rbName,
						SourceNamespace: namespace,
						TargetType:      kind,
						TargetName:      name,
						TargetNamespace: subjNamespace,
						Field:           "subjects",
						Line:            0,
					})
				}
			}
		}
	}
}

// validateNetworkPolicyReferences validates NetworkPolicy selector references
func (v *ResourceValidator) validateNetworkPolicyReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		networkPolicies, ok := resources["NetworkPolicy"]
		if !ok {
			continue
		}

		for npName, npData := range networkPolicies {
			np := npData.(map[string]interface{})
			spec, ok := np["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			// Check pod selector
			if podSelector, ok := spec["podSelector"].(map[string]interface{}); ok {
				v.validateSelectorReferences(npName, namespace, podSelector, "NetworkPolicy", "spec.podSelector", result)
			}

			// Check ingress rules
			if ingress, ok := spec["ingress"].([]interface{}); ok {
				for _, rule := range ingress {
					if ruleMap, ok := rule.(map[string]interface{}); ok {
						if from, ok := ruleMap["from"].([]interface{}); ok {
							for _, fromItem := range from {
								if fromMap, ok := fromItem.(map[string]interface{}); ok {
									if podSelector, ok := fromMap["podSelector"].(map[string]interface{}); ok {
										v.validateSelectorReferences(npName, namespace, podSelector, "NetworkPolicy", "spec.ingress.from.podSelector", result)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// validateIngressReferences validates Ingress backend references
func (v *ResourceValidator) validateIngressReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		ingresses, ok := resources["Ingress"]
		if !ok {
			continue
		}

		for ingressName, ingressData := range ingresses {
			ingress := ingressData.(map[string]interface{})
			spec, ok := ingress["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			// Check default backend
			if defaultBackend, ok := spec["defaultBackend"].(map[string]interface{}); ok {
				v.validateIngressBackend(ingressName, namespace, defaultBackend, "spec.defaultBackend", result)
			}

			// Check rules
			if rules, ok := spec["rules"].([]interface{}); ok {
				for _, rule := range rules {
					if ruleMap, ok := rule.(map[string]interface{}); ok {
						if http, ok := ruleMap["http"].(map[string]interface{}); ok {
							if paths, ok := http["paths"].([]interface{}); ok {
								for _, path := range paths {
									if pathMap, ok := path.(map[string]interface{}); ok {
										if backend, ok := pathMap["backend"].(map[string]interface{}); ok {
											v.validateIngressBackend(ingressName, namespace, backend, "spec.rules.http.paths.backend", result)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// validateHPAReferences validates HorizontalPodAutoscaler target references
func (v *ResourceValidator) validateHPAReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		hpas, ok := resources["HorizontalPodAutoscaler"]
		if !ok {
			continue
		}

		for hpaName, hpaData := range hpas {
			hpa := hpaData.(map[string]interface{})
			spec, ok := hpa["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			if scaleTargetRef, ok := spec["scaleTargetRef"].(map[string]interface{}); ok {
				kind, _ := scaleTargetRef["kind"].(string)
				name, _ := scaleTargetRef["name"].(string)

				if kind == "Deployment" || kind == "StatefulSet" || kind == "ReplicaSet" {
					if !v.resourceExists(namespace, kind, name) {
						result.BrokenReferences = append(result.BrokenReferences, ResourceReference{
							SourceType:      "HorizontalPodAutoscaler",
							SourceName:      hpaName,
							SourceNamespace: namespace,
							TargetType:      kind,
							TargetName:      name,
							TargetNamespace: namespace,
							Field:           "spec.scaleTargetRef",
							Line:            0,
						})
					} else {
						result.ValidReferences = append(result.ValidReferences, ResourceReference{
							SourceType:      "HorizontalPodAutoscaler",
							SourceName:      hpaName,
							SourceNamespace: namespace,
							TargetType:      kind,
							TargetName:      name,
							TargetNamespace: namespace,
							Field:           "spec.scaleTargetRef",
							Line:            0,
						})
					}
				}
			}
		}
	}
}

// validatePDBReferences validates PodDisruptionBudget target references
func (v *ResourceValidator) validatePDBReferences(result *ValidationResult) {
	for namespace, resources := range v.resources {
		pdbs, ok := resources["PodDisruptionBudget"]
		if !ok {
			continue
		}

		for pdbName, pdbData := range pdbs {
			pdb := pdbData.(map[string]interface{})
			spec, ok := pdb["spec"].(map[string]interface{})
			if !ok {
				continue
			}

			if selector, ok := spec["selector"].(map[string]interface{}); ok {
				v.validateSelectorReferences(pdbName, namespace, selector, "PodDisruptionBudget", "spec.selector", result)
			}
		}
	}
}

// validateSelectorReferences validates selector-based references
func (v *ResourceValidator) validateSelectorReferences(sourceName, namespace string, selector map[string]interface{}, sourceType, field string, result *ValidationResult) {
	for key, value := range selector {
		if key == "app" || key == "name" || key == "component" || key == "tier" {
			if !v.selectorTargetExists(namespace, key, value.(string)) {
				result.BrokenReferences = append(result.BrokenReferences, ResourceReference{
					SourceType:      sourceType,
					SourceName:      sourceName,
					SourceNamespace: namespace,
					TargetType:      "Pod/Deployment",
					TargetName:      value.(string),
					TargetNamespace: namespace,
					Field:           field,
					Line:            0,
				})
			} else {
				result.ValidReferences = append(result.ValidReferences, ResourceReference{
					SourceType:      sourceType,
					SourceName:      sourceName,
					SourceNamespace: namespace,
					TargetType:      "Pod/Deployment",
					TargetName:      value.(string),
					TargetNamespace: namespace,
					Field:           field,
					Line:            0,
				})
			}
		}
	}
}

// validateIngressBackend validates Ingress backend references
func (v *ResourceValidator) validateIngressBackend(ingressName, namespace string, backend map[string]interface{}, field string, result *ValidationResult) {
	serviceName, _ := backend["serviceName"].(string)
	if serviceName != "" {
		if !v.resourceExists(namespace, "Service", serviceName) {
			result.BrokenReferences = append(result.BrokenReferences, ResourceReference{
				SourceType:      "Ingress",
				SourceName:      ingressName,
				SourceNamespace: namespace,
				TargetType:      "Service",
				TargetName:      serviceName,
				TargetNamespace: namespace,
				Field:           field,
				Line:            0,
			})
		} else {
			result.ValidReferences = append(result.ValidReferences, ResourceReference{
				SourceType:      "Ingress",
				SourceName:      ingressName,
				SourceNamespace: namespace,
				TargetType:      "Service",
				TargetName:      serviceName,
				TargetNamespace: namespace,
				Field:           field,
				Line:            0,
			})
		}
	}
}

// selectorTargetExists checks if a selector target exists
func (v *ResourceValidator) selectorTargetExists(namespace, key, value string) bool {
	// Check Pods
	if pods, ok := v.resources[namespace]["Pod"]; ok {
		for _, podData := range pods {
			pod := podData.(map[string]interface{})
			if metadata, ok := pod["metadata"].(map[string]interface{}); ok {
				if labels, ok := metadata["labels"].(map[string]interface{}); ok {
					if labelValue, exists := labels[key]; exists && labelValue == value {
						return true
					}
				}
			}
		}
	}

	// Check Deployments
	if deployments, ok := v.resources[namespace]["Deployment"]; ok {
		for _, deployData := range deployments {
			deploy := deployData.(map[string]interface{})
			if metadata, ok := deploy["metadata"].(map[string]interface{}); ok {
				if labels, ok := metadata["labels"].(map[string]interface{}); ok {
					if labelValue, exists := labels[key]; exists && labelValue == value {
						return true
					}
				}
			}
		}
	}

	// Check StatefulSets
	if statefulSets, ok := v.resources[namespace]["StatefulSet"]; ok {
		for _, ssData := range statefulSets {
			ss := ssData.(map[string]interface{})
			if metadata, ok := ss["metadata"].(map[string]interface{}); ok {
				if labels, ok := metadata["labels"].(map[string]interface{}); ok {
					if labelValue, exists := labels[key]; exists && labelValue == value {
						return true
					}
				}
			}
		}
	}

	// Check ReplicaSets
	if replicaSets, ok := v.resources[namespace]["ReplicaSet"]; ok {
		for _, rsData := range replicaSets {
			rs := rsData.(map[string]interface{})
			if metadata, ok := rs["metadata"].(map[string]interface{}); ok {
				if labels, ok := metadata["labels"].(map[string]interface{}); ok {
					if labelValue, exists := labels[key]; exists && labelValue == value {
						return true
					}
				}
			}
		}
	}

	return false
}

// resourceExists checks if a resource exists
func (v *ResourceValidator) resourceExists(namespace, resourceType, name string) bool {
	if resources, ok := v.resources[namespace]; ok {
		if resourceTypeResources, ok := resources[resourceType]; ok {
			_, exists := resourceTypeResources[name]
			return exists
		}
	}
	return false
}
