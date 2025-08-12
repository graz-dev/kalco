package dumper

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// Dumper handles the dumping of Kubernetes resources
type Dumper struct {
	clientset       kubernetes.Interface
	discoveryClient discovery.DiscoveryInterface
	dynamicClient   dynamic.Interface
}

// NewDumper creates a new Dumper instance
func NewDumper(clientset kubernetes.Interface, discoveryClient discovery.DiscoveryInterface) *Dumper {
	// We need to pass the config from the kube package, so we'll create a new function signature
	// For now, we'll create a dummy dynamic client that will be set later
	return &Dumper{
		clientset:       clientset,
		discoveryClient: discoveryClient,
		dynamicClient:   nil, // Will be set by SetDynamicClient
	}
}

// SetDynamicClient sets the dynamic client for the dumper
func (d *Dumper) SetDynamicClient(dynamicClient dynamic.Interface) {
	d.dynamicClient = dynamicClient
}

// DumpAllResources performs the main task of dumping all resources
func (d *Dumper) DumpAllResources(outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Get all server resources
	fmt.Println("  📡 Discovering API resources...")
	resourceLists, err := d.discoveryClient.ServerPreferredResources()
	if err != nil {
		return fmt.Errorf("failed to get server resources: %w", err)
	}
	fmt.Printf("  ✅ Found %d API resource groups\n", len(resourceLists))

	// Get all namespaces for namespaced resources
	fmt.Println("  🏷️  Enumerating namespaces...")
	namespaces, err := d.clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list namespaces: %v", err)
	}
	fmt.Printf("  ✅ Found %d namespaces\n", len(namespaces.Items))

	// Process each resource group
	fmt.Println("  🔄 Processing resource groups...")
	processedGroups := 0
	for _, resourceList := range resourceLists {
		if err := d.processResourceGroup(resourceList, namespaces.Items, outputDir); err != nil {
			fmt.Printf("  ⚠️  Warning: failed to process resource group %s: %v\n", resourceList.GroupVersion, err)
			continue
		}
		processedGroups++
		fmt.Printf("  ✅ Processed resource group %s\n", resourceList.GroupVersion)
	}

	fmt.Printf("  🎯 Successfully processed %d/%d resource groups\n", processedGroups, len(resourceLists))
	return nil
}

// processResourceGroup processes a single API resource group
func (d *Dumper) processResourceGroup(resourceList *metav1.APIResourceList, namespaces []corev1.Namespace, outputDir string) error {
	validResources := 0
	for _, resource := range resourceList.APIResources {
		// Skip subresources
		if strings.Contains(resource.Name, "/") {
			continue
		}
		validResources++

		// Create GVR for the resource
		// Parse GroupVersion to get Group and Version separately
		groupVersion := resourceList.GroupVersion
		var group, version string
		if groupVersion == "v1" {
			group = ""
			version = "v1"
		} else {
			parts := strings.Split(groupVersion, "/")
			if len(parts) == 2 {
				group = parts[0]
				version = parts[1]
			} else {
				group = groupVersion
				version = ""
			}
		}

		gvr := schema.GroupVersionResource{
			Group:    group,
			Version:  version,
			Resource: resource.Name,
		}

		if resource.Namespaced {
			// Handle namespaced resources
			if err := d.dumpNamespacedResources(gvr, resource, namespaces, outputDir); err != nil {
				fmt.Printf("    ⚠️  Warning: failed to dump namespaced resource %s: %v\n", resource.Kind, err)
			}
		} else {
			// Handle cluster-scoped resources
			if err := d.dumpClusterScopedResources(gvr, resource, outputDir); err != nil {
				fmt.Printf("    ⚠️  Warning: failed to dump cluster-scoped resource %s: %v\n", resource.Kind, err)
			}
		}
	}

	if validResources > 0 {
		fmt.Printf("    📊 Processing %d resources in group %s\n", validResources, resourceList.GroupVersion)
	}
	return nil
}

// dumpNamespacedResources dumps all instances of a namespaced resource across all namespaces
func (d *Dumper) dumpNamespacedResources(gvr schema.GroupVersionResource, resource metav1.APIResource, namespaces []corev1.Namespace, outputDir string) error {
	totalResources := 0
	for _, namespace := range namespaces {
		// Create directory structure: <outputDir>/<namespace>/<resource_kind>
		resourceDir := filepath.Join(outputDir, namespace.Name, resource.Kind)
		if err := os.MkdirAll(resourceDir, 0755); err != nil {
			return fmt.Errorf("failed to create resource directory: %w", err)
		}

		// List all resources of this type in the namespace
		resourceList, err := d.dynamicClient.Resource(gvr).Namespace(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("      ⚠️  Warning: failed to list %s in namespace %s: %v\n", resource.Kind, namespace.Name, err)
			continue
		}

		if len(resourceList.Items) > 0 {
			fmt.Printf("      📁 Namespace %s: %d %s resources\n", namespace.Name, len(resourceList.Items), resource.Kind)
		}

		// Dump each resource instance
		for _, item := range resourceList.Items {
			if err := d.dumpResource(item, resourceDir); err != nil {
				fmt.Printf("        ⚠️  Warning: failed to dump %s %s in namespace %s: %v\n", resource.Kind, item.GetName(), namespace.Name, err)
			}
		}
		totalResources += len(resourceList.Items)
	}

	if totalResources > 0 {
		fmt.Printf("      ✅ Total %s resources dumped: %d\n", resource.Kind, totalResources)
	}
	return nil
}

// dumpClusterScopedResources dumps all instances of a cluster-scoped resource
func (d *Dumper) dumpClusterScopedResources(gvr schema.GroupVersionResource, resource metav1.APIResource, outputDir string) error {
	// Create directory structure: <outputDir>/_cluster/<resource_kind>
	resourceDir := filepath.Join(outputDir, "_cluster", resource.Kind)
	if err := os.MkdirAll(resourceDir, 0755); err != nil {
		return fmt.Errorf("failed to create resource directory: %w", err)
	}

	// List all resources of this type at cluster level
	resourceList, err := d.dynamicClient.Resource(gvr).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list cluster-scoped %s: %w", resource.Kind, err)
	}

	if len(resourceList.Items) > 0 {
		fmt.Printf("      🌐 Cluster-scoped %s: %d resources\n", resource.Kind, len(resourceList.Items))
	}

	// Dump each resource instance
	for _, item := range resourceList.Items {
		if err := d.dumpResource(item, resourceDir); err != nil {
			fmt.Printf("        ⚠️  Warning: failed to dump cluster-scoped %s %s: %v\n", resource.Kind, item.GetName(), err)
		}
	}

	if len(resourceList.Items) > 0 {
		fmt.Printf("      ✅ Total cluster-scoped %s resources dumped: %d\n", resource.Kind, len(resourceList.Items))
	}
	return nil
}

// dumpResource dumps a single resource instance to a YAML file
func (d *Dumper) dumpResource(item unstructured.Unstructured, resourceDir string) error {
	// Clean up metadata fields that are not useful for re-application
	cleanupMetadata(&item)

	// Convert to YAML
	yamlData, err := yaml.Marshal(item.Object)
	if err != nil {
		return fmt.Errorf("failed to marshal resource to YAML: %w", err)
	}

	// Create filename: <resource_name>.yaml
	filename := filepath.Join(resourceDir, item.GetName()+".yaml")

	// Write to file
	if err := os.WriteFile(filename, yamlData, 0644); err != nil {
		return fmt.Errorf("failed to write YAML file: %w", err)
	}

	return nil
}

// cleanupMetadata removes metadata fields that are not useful for re-application
func cleanupMetadata(item *unstructured.Unstructured) {
	metadata, exists, err := unstructured.NestedMap(item.Object, "metadata")
	if !exists || err != nil {
		return
	}

	// Remove fields that are not useful for re-application
	delete(metadata, "uid")
	delete(metadata, "resourceVersion")
	delete(metadata, "generation")
	delete(metadata, "creationTimestamp")
	delete(metadata, "managedFields")
	delete(metadata, "ownerReferences")

	// Update the cleaned metadata
	item.Object["metadata"] = metadata

	// Remove status field entirely
	delete(item.Object, "status")
}
