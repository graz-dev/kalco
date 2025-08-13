package orphaned

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// OrphanedResource represents an orphaned resource
type OrphanedResource struct {
	Type      string `yaml:"type"`
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Reason    string `yaml:"reason"`
	Details   string `yaml:"details"`
	File      string `yaml:"file"`
}

// OrphanedResult represents the result of orphaned resource detection
type OrphanedResult struct {
	OrphanedResources []OrphanedResource `yaml:"orphanedResources"`
	Summary           OrphanedSummary    `yaml:"summary"`
}

// OrphanedSummary provides a summary of orphaned resource detection
type OrphanedSummary struct {
	TotalOrphanedResources int            `yaml:"totalOrphanedResources"`
	ByType                 map[string]int `yaml:"byType"`
}

// OrphanedDetector handles orphaned resource detection
type OrphanedDetector struct {
	outputDir string
	resources map[string]map[string]map[string]interface{}
}

// NewOrphanedDetector creates a new detector instance
func NewOrphanedDetector(outputDir string) *OrphanedDetector {
	return &OrphanedDetector{
		outputDir: outputDir,
		resources: make(map[string]map[string]map[string]interface{}),
	}
}

// Detect performs orphaned resource detection on all exported resources
func (d *OrphanedDetector) Detect() (*OrphanedResult, error) {
	// Load all resources
	if err := d.loadResources(); err != nil {
		return nil, fmt.Errorf("failed to load resources: %w", err)
	}

	result := &OrphanedResult{
		OrphanedResources: []OrphanedResource{},
	}

	// Detect different types of orphaned resources
	d.detectOrphanedReplicaSets(result)
	d.detectOrphanedPods(result)
	d.detectOrphanedConfigMaps(result)
	d.detectOrphanedSecrets(result)
	d.detectOrphanedServices(result)
	d.detectOrphanedPVCs(result)

	// Calculate summary
	result.Summary = OrphanedSummary{
		TotalOrphanedResources: len(result.OrphanedResources),
		ByType:                 make(map[string]int),
	}

	for _, resource := range result.OrphanedResources {
		result.Summary.ByType[resource.Type]++
	}

	return result, nil
}

// loadResources loads all YAML resources from the output directory
func (d *OrphanedDetector) loadResources() error {
	return filepath.Walk(d.outputDir, func(path string, info os.FileInfo, err error) error {
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
		relPath, err := filepath.Rel(d.outputDir, path)
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
		if d.resources[namespace] == nil {
			d.resources[namespace] = make(map[string]map[string]interface{})
		}
		if d.resources[namespace][resourceType] == nil {
			d.resources[namespace][resourceType] = make(map[string]interface{})
		}

		// Extract resource name from filename
		resourceName := strings.TrimSuffix(filename, ".yaml")
		d.resources[namespace][resourceType][resourceName] = resource

		return nil
	})
}

// detectOrphanedReplicaSets detects ReplicaSets not owned by Deployments
func (d *OrphanedDetector) detectOrphanedReplicaSets(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		replicaSets, ok := resources["ReplicaSet"]
		if !ok {
			continue
		}

		for rsName, rsData := range replicaSets {
			rs := rsData.(map[string]interface{})
			metadata, ok := rs["metadata"].(map[string]interface{})
			if !ok {
				continue
			}

			// Check if this ReplicaSet is owned by a Deployment
			ownerReferences, ok := metadata["ownerReferences"].([]interface{})
			if !ok || len(ownerReferences) == 0 {
				// No owner references - potentially orphaned
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "ReplicaSet",
					Name:      rsName,
					Namespace: namespace,
					Reason:    "No Owner References",
					Details:   "This ReplicaSet has no owner references and may be orphaned",
					File:      fmt.Sprintf("%s/ReplicaSet/%s.yaml", namespace, rsName),
				})
				continue
			}

			// Check if any owner is a Deployment
			hasDeploymentOwner := false
			for _, owner := range ownerReferences {
				ownerMap, ok := owner.(map[string]interface{})
				if !ok {
					continue
				}

				kind, _ := ownerMap["kind"].(string)
				if kind == "Deployment" {
					hasDeploymentOwner = true
					break
				}
			}

			if !hasDeploymentOwner {
				// Has owner references but none are Deployments
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "ReplicaSet",
					Name:      rsName,
					Namespace: namespace,
					Reason:    "No Deployment Owner",
					Details:   "This ReplicaSet is not owned by a Deployment and may be orphaned",
					File:      fmt.Sprintf("%s/ReplicaSet/%s.yaml", namespace, rsName),
				})
			}
		}
	}
}

// detectOrphanedPods detects Pods not owned by any controller
func (d *OrphanedDetector) detectOrphanedPods(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		pods, ok := resources["Pod"]
		if !ok {
			continue
		}

		for podName, podData := range pods {
			pod := podData.(map[string]interface{})
			metadata, ok := pod["metadata"].(map[string]interface{})
			if !ok {
				continue
			}

			// Check if this Pod is owned by any controller
			ownerReferences, ok := metadata["ownerReferences"].([]interface{})
			if !ok || len(ownerReferences) == 0 {
				// Check if it's a static pod (managed by kubelet)
				annotations, ok := metadata["annotations"].(map[string]interface{})
				if ok {
					if _, isStaticPod := annotations["kubernetes.io/config.source"]; isStaticPod {
						continue // Skip static pods
					}
				}

				// Check if it's a mirror pod
				if annotations, ok := metadata["annotations"].(map[string]interface{}); ok {
					if _, isMirrorPod := annotations["kubernetes.io/config.mirror"]; isMirrorPod {
						continue // Skip mirror pods
					}
				}

				// No owner references and not a static/mirror pod - potentially orphaned
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "Pod",
					Name:      podName,
					Namespace: namespace,
					Reason:    "No Controller Owner",
					Details:   "This Pod has no controller owner and may be orphaned",
					File:      fmt.Sprintf("%s/Pod/%s.yaml", namespace, podName),
				})
			}
		}
	}
}

// detectOrphanedConfigMaps detects ConfigMaps not referenced by any Pod/Deployment
func (d *OrphanedDetector) detectOrphanedConfigMaps(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		configMaps, ok := resources["ConfigMap"]
		if !ok {
			continue
		}

		for cmName, _ := range configMaps {
			// Check if this ConfigMap is referenced by any Pod/Deployment
			if !d.isConfigMapReferenced(namespace, cmName) {
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "ConfigMap",
					Name:      cmName,
					Namespace: namespace,
					Reason:    "No References",
					Details:   "This ConfigMap is not referenced by any Pod or Deployment",
					File:      fmt.Sprintf("%s/ConfigMap/%s.yaml", namespace, cmName),
				})
			}
		}
	}
}

// detectOrphanedSecrets detects Secrets not referenced by any Pod/Deployment
func (d *OrphanedDetector) detectOrphanedSecrets(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		secrets, ok := resources["Secret"]
		if !ok {
			continue
		}

		for secretName, _ := range secrets {
			// Check if this Secret is referenced by any Pod/Deployment
			if !d.isSecretReferenced(namespace, secretName) {
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "Secret",
					Name:      secretName,
					Namespace: namespace,
					Reason:    "No References",
					Details:   "This Secret is not referenced by any Pod or Deployment",
					File:      fmt.Sprintf("%s/Secret/%s.yaml", namespace, secretName),
				})
			}
		}
	}
}

// detectOrphanedServices detects Services not referenced by any Pod/Deployment
func (d *OrphanedDetector) detectOrphanedServices(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		services, ok := resources["Service"]
		if !ok {
			continue
		}

		for svcName, _ := range services {
			// Check if this Service is referenced by any Pod/Deployment
			if !d.isServiceReferenced(namespace, svcName) {
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "Service",
					Name:      svcName,
					Namespace: namespace,
					Reason:    "No References",
					Details:   "This Service is not referenced by any Pod or Deployment",
					File:      fmt.Sprintf("%s/Service/%s.yaml", namespace, svcName),
				})
			}
		}
	}
}

// detectOrphanedPVCs detects PersistentVolumeClaims not referenced by any Pod
func (d *OrphanedDetector) detectOrphanedPVCs(result *OrphanedResult) {
	for namespace, resources := range d.resources {
		pvcs, ok := resources["PersistentVolumeClaim"]
		if !ok {
			continue
		}

		for pvcName, _ := range pvcs {
			// Check if this PVC is referenced by any Pod
			if !d.isPVCReferenced(namespace, pvcName) {
				result.OrphanedResources = append(result.OrphanedResources, OrphanedResource{
					Type:      "PersistentVolumeClaim",
					Name:      pvcName,
					Namespace: namespace,
					Reason:    "No References",
					Details:   "This PersistentVolumeClaim is not referenced by any Pod",
					File:      fmt.Sprintf("%s/PersistentVolumeClaim/%s.yaml", namespace, pvcName),
				})
			}
		}
	}
}

// isConfigMapReferenced checks if a ConfigMap is referenced by any Pod/Deployment
func (d *OrphanedDetector) isConfigMapReferenced(namespace, cmName string) bool {
	// Check Pods
	if pods, ok := d.resources[namespace]["Pod"]; ok {
		for _, podData := range pods {
			pod := podData.(map[string]interface{})
			if d.podReferencesConfigMap(pod, cmName) {
				return true
			}
		}
	}

	// Check Deployments
	if deployments, ok := d.resources[namespace]["Deployment"]; ok {
		for _, deployData := range deployments {
			deploy := deployData.(map[string]interface{})
			if d.deploymentReferencesConfigMap(deploy, cmName) {
				return true
			}
		}
	}

	return false
}

// isSecretReferenced checks if a Secret is referenced by any Pod/Deployment
func (d *OrphanedDetector) isSecretReferenced(namespace, secretName string) bool {
	// Check Pods
	if pods, ok := d.resources[namespace]["Pod"]; ok {
		for _, podData := range pods {
			pod := podData.(map[string]interface{})
			if d.podReferencesSecret(pod, secretName) {
				return true
			}
		}
	}

	// Check Deployments
	if deployments, ok := d.resources[namespace]["Deployment"]; ok {
		for _, deployData := range deployments {
			deploy := deployData.(map[string]interface{})
			if d.deploymentReferencesSecret(deploy, secretName) {
				return true
			}
		}
	}

	return false
}

// isServiceReferenced checks if a Service is referenced by any Pod/Deployment
func (d *OrphanedDetector) isServiceReferenced(namespace, svcName string) bool {
	// Check Pods
	if pods, ok := d.resources[namespace]["Pod"]; ok {
		for _, podData := range pods {
			pod := podData.(map[string]interface{})
			if d.podReferencesService(pod, svcName) {
				return true
			}
		}
	}

	// Check Deployments
	if deployments, ok := d.resources[namespace]["Deployment"]; ok {
		for _, deployData := range deployments {
			deploy := deployData.(map[string]interface{})
			if d.deploymentReferencesService(deploy, svcName) {
				return true
			}
		}
	}

	return false
}

// isPVCReferenced checks if a PVC is referenced by any Pod
func (d *OrphanedDetector) isPVCReferenced(namespace, pvcName string) bool {
	// Check Pods
	if pods, ok := d.resources[namespace]["Pod"]; ok {
		for _, podData := range pods {
			pod := podData.(map[string]interface{})
			if d.podReferencesPVC(pod, pvcName) {
				return true
			}
		}
	}

	return false
}

// podReferencesConfigMap checks if a Pod references a ConfigMap
func (d *OrphanedDetector) podReferencesConfigMap(pod map[string]interface{}, cmName string) bool {
	spec, ok := pod["spec"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check volumes
	if volumes, ok := spec["volumes"].([]interface{}); ok {
		for _, volume := range volumes {
			vol := volume.(map[string]interface{})
			if configMap, ok := vol["configMap"].(map[string]interface{}); ok {
				if name, _ := configMap["name"].(string); name == cmName {
					return true
				}
			}
		}
	}

	// Check env vars
	if containers, ok := spec["containers"].([]interface{}); ok {
		for _, container := range containers {
			cont := container.(map[string]interface{})
			if env, ok := cont["env"].([]interface{}); ok {
				for _, envVar := range env {
					envMap := envVar.(map[string]interface{})
					if valueFrom, ok := envMap["valueFrom"].(map[string]interface{}); ok {
						if configMapKeyRef, ok := valueFrom["configMapKeyRef"].(map[string]interface{}); ok {
							if name, _ := configMapKeyRef["name"].(string); name == cmName {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

// podReferencesSecret checks if a Pod references a Secret
func (d *OrphanedDetector) podReferencesSecret(pod map[string]interface{}, secretName string) bool {
	spec, ok := pod["spec"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check volumes
	if volumes, ok := spec["volumes"].([]interface{}); ok {
		for _, volume := range volumes {
			vol := volume.(map[string]interface{})
			if secret, ok := vol["secret"].(map[string]interface{}); ok {
				if name, _ := secret["secretName"].(string); name == secretName {
					return true
				}
			}
		}
	}

	// Check env vars
	if containers, ok := spec["containers"].([]interface{}); ok {
		for _, container := range containers {
			cont := container.(map[string]interface{})
			if env, ok := cont["env"].([]interface{}); ok {
				for _, envVar := range env {
					envMap := envVar.(map[string]interface{})
					if valueFrom, ok := envMap["valueFrom"].(map[string]interface{}); ok {
						if secretKeyRef, ok := valueFrom["secretKeyRef"].(map[string]interface{}); ok {
							if name, _ := secretKeyRef["name"].(string); name == secretName {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

// podReferencesService checks if a Pod references a Service
func (d *OrphanedDetector) podReferencesService(pod map[string]interface{}, svcName string) bool {
	// Pods typically don't directly reference Services by name
	// This is more of a runtime relationship
	return false
}

// podReferencesPVC checks if a Pod references a PVC
func (d *OrphanedDetector) podReferencesPVC(pod map[string]interface{}, pvcName string) bool {
	spec, ok := pod["spec"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check volumes
	if volumes, ok := spec["volumes"].([]interface{}); ok {
		for _, volume := range volumes {
			vol := volume.(map[string]interface{})
			if persistentVolumeClaim, ok := vol["persistentVolumeClaim"].(map[string]interface{}); ok {
				if name, _ := persistentVolumeClaim["claimName"].(string); name == pvcName {
					return true
				}
			}
		}
	}

	return false
}

// deploymentReferencesConfigMap checks if a Deployment references a ConfigMap
func (d *OrphanedDetector) deploymentReferencesConfigMap(deploy map[string]interface{}, cmName string) bool {
	spec, ok := deploy["spec"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check template
	if template, ok := spec["template"].(map[string]interface{}); ok {
		return d.podReferencesConfigMap(template, cmName)
	}

	return false
}

// deploymentReferencesSecret checks if a Deployment references a Secret
func (d *OrphanedDetector) deploymentReferencesSecret(deploy map[string]interface{}, secretName string) bool {
	spec, ok := deploy["spec"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check template
	if template, ok := spec["template"].(map[string]interface{}); ok {
		return d.podReferencesSecret(template, secretName)
	}

	return false
}

// deploymentReferencesService checks if a Deployment references a Service
func (d *OrphanedDetector) deploymentReferencesService(deploy map[string]interface{}, svcName string) bool {
	// Deployments typically don't directly reference Services by name
	// This is more of a runtime relationship
	return false
}
