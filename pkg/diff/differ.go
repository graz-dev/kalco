package diff

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Resource represents a K8s resource for diffing
type Resource struct {
	Object    *unstructured.Unstructured
	Source    string // "local", "snapshot", "live"
	Timestamp time.Time
	Exists    bool
}

// Differ handles the diff logic
type Differ struct {
}

// NewDiffer creates a new Differ
func NewDiffer() *Differ {
	return &Differ{}
}

// LoadLocalResource parses a local YAML file (supports multi-doc)
func (d *Differ) LoadLocalResource(path string) ([]*Resource, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open local file: %w", err)
	}
	defer f.Close()

	var resources []*Resource
	decoder := yaml.NewDecoder(f)

	for {
		var obj map[string]interface{}
		if err := decoder.Decode(&obj); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}

		if len(obj) == 0 {
			continue
		}

		// Fix types for DeepCopy compatibility (int -> int64)
		obj = fixTypes(obj).(map[string]interface{})

		u := &unstructured.Unstructured{Object: obj}
		resources = append(resources, &Resource{
			Object: u,
			Source: "local",
			Exists: true,
		})
	}

	if len(resources) == 0 {
		return nil, fmt.Errorf("no valid resources found in file")
	}

	return resources, nil
}

func fixTypes(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, val := range v {
			result[k] = fixTypes(val)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, val := range v {
			result[i] = fixTypes(val)
		}
		return result
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case float64:
		// Check if it's actually an integer
		if v == float64(int64(v)) {
			return int64(v)
		}
		return v
	default:
		return v
	}
}

// LoadSnapshotResource loads a resource from the snapshot directory
// Maps GVK from local resource to expected file path
// Logic: <snapshotDir>/<namespace>/<Kind>/<name>.yaml
// For cluster-scoped: <snapshotDir>/_cluster/<Kind>/<name>.yaml
func (d *Differ) LoadSnapshotResource(snapshotDir string, local *unstructured.Unstructured, overrideNamespace string) (*Resource, error) {
	kind := local.GetKind()
	name := local.GetName()
	namespace := local.GetNamespace()

	var path string
	if namespace == "" {
		if overrideNamespace != "" {
			namespace = overrideNamespace
		} else {
			namespace = "default"
		}
	}

	// But wait, cluster resources are stored in _cluster.
	// We need to know if it's namespaced or not to look in the right place?
	// The prompt says: "Leggi il file --target ... per estrarre: Kind, Metadata.Name, Metadata.Namespace (default "default")."
	// Then: "Costruisci il path atteso nello snapshot: "{snapshot-dir}/{namespace}/{Kind}/{name}.yaml"."
	// It doesn't explicitly mention _cluster here, but the user context (previous turns) shows _cluster.
	// We should probably check if the Kind is usually cluster scoped?
	// For now, let's stick to the prompt's instruction: default namespace to "default".
	// If the file is not found there, we could optionally check _cluster if we knew it was cluster scoped.
	// But let's assume valid namespaces for now.

	// Helper to check file existence
	checkPath := func(ns string) (string, bool) {
		p := filepath.Join(snapshotDir, ns, kind, name+".yaml")
		if _, err := os.Stat(p); err == nil {
			return p, true
		}
		return p, false
	}

	// Check typical namespaced path
	path, exists := checkPath(namespace)

	// If not found, and namespace was "default" (maybe it was really cluster scoped?)
	// Try _cluster path
	if !exists && namespace == "default" {
		clusterPath := filepath.Join(snapshotDir, "_cluster", kind, name+".yaml")
		if _, err := os.Stat(clusterPath); err == nil {
			path = clusterPath
			exists = true
		}
	}

	if !exists {
		return &Resource{Exists: false, Source: "snapshot"}, nil
	}

	// Load file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read snapshot: %w", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("failed to parse snapshot YAML: %w", err)
	}

	obj = fixTypes(obj).(map[string]interface{})

	return &Resource{
		Object:    &unstructured.Unstructured{Object: obj},
		Source:    "snapshot",
		Exists:    true,
		Timestamp: info.ModTime(),
	}, nil
}

// DiffResult contains the result of the comparison
type DiffResult struct {
	Diff     string
	IsNew    bool
	IsStale  bool
	HasDrift bool
}

// Diff compares two resources
func (d *Differ) Diff(local, remote *Resource) (*DiffResult, error) {
	if !remote.Exists {
		return &DiffResult{IsNew: true}, nil
	}

	// Normalize
	l := cleanObject(local.Object)
	r := cleanObject(remote.Object)

	// Compare
	diff := cmp.Diff(r, l) // Remote is "base", Local is "new" intention

	hasDrift := diff != ""

	// Check for critical drift if requested?
	// For now just basic diff

	isStale := false
	if remote.Source == "snapshot" {
		if time.Since(remote.Timestamp) > 24*time.Hour {
			isStale = true
		}
	}

	return &DiffResult{
		Diff:     diff,
		IsNew:    false,
		IsStale:  isStale,
		HasDrift: hasDrift,
	}, nil
}

func cleanObject(u *unstructured.Unstructured) map[string]interface{} {
	obj := u.DeepCopy().Object

	// Remove metadata fields we don't care about for intent vs reality check
	if meta, ok := obj["metadata"].(map[string]interface{}); ok {
		delete(meta, "uid")
		delete(meta, "resourceVersion")
		delete(meta, "generation")
		delete(meta, "creationTimestamp")
		delete(meta, "managedFields")
		delete(meta, "ownerReferences")
		delete(meta, "selfLink")
	}

	delete(obj, "status")

	return obj
}
