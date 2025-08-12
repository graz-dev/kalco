package dumper

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	discoveryfake "k8s.io/client-go/discovery/fake"
	kubernetesfake "k8s.io/client-go/kubernetes/fake"
)

func TestNewDumper(t *testing.T) {
	fakeClientset := kubernetesfake.NewSimpleClientset()
	fakeDiscovery := &discoveryfake.FakeDiscovery{}

	d := NewDumper(fakeClientset, fakeDiscovery)
	if d == nil {
		t.Fatal("NewDumper returned nil")
	}

	if d.clientset != fakeClientset {
		t.Error("clientset not set correctly")
	}

	if d.discoveryClient != fakeDiscovery {
		t.Error("discoveryClient not set correctly")
	}
}

func TestCleanupMetadata(t *testing.T) {
	item := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"name":              "test-resource",
				"uid":               "test-uid",
				"resourceVersion":   "123",
				"generation":        "1",
				"creationTimestamp": "2023-01-01T00:00:00Z",
				"managedFields":     []interface{}{"field1"},
				"ownerReferences":   []interface{}{"ref1"},
			},
			"status": map[string]interface{}{
				"phase": "Running",
			},
		},
	}

	cleanupMetadata(item)

	metadata, exists, _ := unstructured.NestedMap(item.Object, "metadata")
	if !exists {
		t.Fatal("metadata should still exist")
	}

	// Check that problematic fields were removed
	if _, exists := metadata["uid"]; exists {
		t.Error("uid should have been removed")
	}
	if _, exists := metadata["resourceVersion"]; exists {
		t.Error("resourceVersion should have been removed")
	}
	if _, exists := metadata["generation"]; exists {
		t.Error("generation should have been removed")
	}
	if _, exists := metadata["creationTimestamp"]; exists {
		t.Error("creationTimestamp should have been removed")
	}
	if _, exists := metadata["managedFields"]; exists {
		t.Error("managedFields should have been removed")
	}
	if _, exists := metadata["ownerReferences"]; exists {
		t.Error("ownerReferences should have been removed")
	}

	// Check that useful fields remain
	if _, exists := metadata["name"]; !exists {
		t.Error("name should not have been removed")
	}

	// Check that status was removed
	if _, exists := item.Object["status"]; exists {
		t.Error("status should have been removed")
	}
}
