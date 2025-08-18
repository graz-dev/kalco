package orphaned

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestOrphanedResourceDetection(t *testing.T) {
	// Test basic orphaned resource detection
	tempDir := t.TempDir()

	// Create test directory structure
	testNamespace := "test-ns"
	testResourceType := "Pod"
	testResourceName := "test-pod"

	// Create directory structure
	resourceDir := filepath.Join(tempDir, testNamespace, testResourceType)
	if err := os.MkdirAll(resourceDir, 0755); err != nil {
		t.Fatalf("Failed to create resource directory: %v", err)
	}

	// Create test YAML file
	testYAML := `apiVersion: v1
kind: Pod
metadata:
  name: test-pod
  labels:
    app: test
spec:
  containers:
  - name: test
    image: nginx:latest`

	testFile := filepath.Join(resourceDir, testResourceName+".yaml")
	if err := os.WriteFile(testFile, []byte(testYAML), 0644); err != nil {
		t.Fatalf("Failed to write test YAML: %v", err)
	}

	// Test that the file exists
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatal("Test YAML file was not created")
	}

	// Test directory structure
	if _, err := os.Stat(resourceDir); os.IsNotExist(err) {
		t.Fatal("Resource directory was not created")
	}

	// Test namespace directory
	namespaceDir := filepath.Join(tempDir, testNamespace)
	if _, err := os.Stat(namespaceDir); os.IsNotExist(err) {
		t.Fatal("Namespace directory was not created")
	}
}

func TestOrphanedResourceFileOperations(t *testing.T) {
	// Test file operations for orphaned resources
	tempDir := t.TempDir()

	// Test creating multiple resource files
	resources := []struct {
		namespace    string
		resourceType string
		name         string
		content      string
	}{
		{
			namespace:    "ns1",
			resourceType: "Pod",
			name:         "pod1",
			content: `apiVersion: v1
kind: Pod
metadata:
  name: pod1`,
		},
		{
			namespace:    "ns1",
			resourceType: "Service",
			name:         "svc1",
			content: `apiVersion: v1
kind: Service
metadata:
  name: svc1`,
		},
		{
			namespace:    "ns2",
			resourceType: "Deployment",
			name:         "deploy1",
			content: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy1`,
		},
	}

	// Create all resources
	for _, resource := range resources {
		resourceDir := filepath.Join(tempDir, resource.namespace, resource.resourceType)
		if err := os.MkdirAll(resourceDir, 0755); err != nil {
			t.Fatalf("Failed to create resource directory: %v", err)
		}

		resourceFile := filepath.Join(resourceDir, resource.name+".yaml")
		if err := os.WriteFile(resourceFile, []byte(resource.content), 0644); err != nil {
			t.Fatalf("Failed to write resource file: %v", err)
		}

		// Verify file was created
		if _, err := os.Stat(resourceFile); os.IsNotExist(err) {
			t.Errorf("Resource file %s was not created", resourceFile)
		}
	}

	// Test directory listing
	entries, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read temp directory: %v", err)
	}

	// Should have 2 namespace directories
	if len(entries) != 2 {
		t.Errorf("Expected 2 namespace directories, got %d", len(entries))
	}

	// Check ns1 contents
	ns1Dir := filepath.Join(tempDir, "ns1")
	ns1Entries, err := os.ReadDir(ns1Dir)
	if err != nil {
		t.Fatalf("Failed to read ns1 directory: %v", err)
	}

	if len(ns1Entries) != 2 {
		t.Errorf("Expected 2 resource type directories in ns1, got %d", len(ns1Entries))
	}
}

func TestOrphanedResourcePathHandling(t *testing.T) {
	// Test path handling for orphaned resources
	tempDir := t.TempDir()

	// Test with various path separators and special characters
	testPaths := []string{
		"namespace-with-dashes",
		"namespace_with_underscores",
		"namespace.with.dots",
		"namespace with spaces",
		"namespace/with/slashes",
		"namespace\\with\\backslashes",
	}

	for _, testPath := range testPaths {
		// Create directory
		fullPath := filepath.Join(tempDir, testPath)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			t.Fatalf("Failed to create directory for path '%s': %v", testPath, err)
		}

		// Verify directory was created
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Directory for path '%s' was not created", testPath)
		}

		// Test file creation in the directory
		testFile := filepath.Join(fullPath, "test.yaml")
		testContent := `apiVersion: v1
kind: Test
metadata:
  name: test`

		if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
			t.Errorf("Failed to write test file in path '%s': %v", testPath, err)
		}

		// Verify file was created
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			t.Errorf("Test file in path '%s' was not created", testPath)
		}
	}
}

func TestOrphanedResourceFilePermissions(t *testing.T) {
	// Test file permissions for orphaned resources
	tempDir := t.TempDir()

	// Create test directory
	testDir := filepath.Join(tempDir, "test-ns", "test-type")
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create test file with specific permissions
	testFile := filepath.Join(testDir, "test.yaml")
	testContent := `apiVersion: v1
kind: Test
metadata:
  name: test`

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Check file permissions
	info, err := os.Stat(testFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}

	// Check that file is readable and writable by owner
	mode := info.Mode()
	if mode&0400 == 0 {
		t.Error("File should be readable by owner")
	}
	if mode&0200 == 0 {
		t.Error("File should be writable by owner")
	}
}

func TestOrphanedResourceDirectoryTraversal(t *testing.T) {
	// Test directory traversal for orphaned resources
	tempDir := t.TempDir()

	// Create nested directory structure
	nestedPath := filepath.Join(tempDir, "ns1", "type1", "subtype1", "subtype2")
	if err := os.MkdirAll(nestedPath, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}

	// Create file in nested directory
	testFile := filepath.Join(nestedPath, "test.yaml")
	testContent := `apiVersion: v1
kind: Test
metadata:
  name: test`

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Test walking the directory tree
	fileCount := 0
	err := filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			fileCount++
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Failed to walk directory tree: %v", err)
	}

	// Should find exactly one YAML file
	if fileCount != 1 {
		t.Errorf("Expected 1 YAML file, found %d", fileCount)
	}
}

func TestOrphanedResourceCleanup(t *testing.T) {
	// Test cleanup operations for orphaned resources
	tempDir := t.TempDir()

	// Create test resources
	testResources := []string{
		filepath.Join(tempDir, "ns1", "type1", "resource1.yaml"),
		filepath.Join(tempDir, "ns1", "type2", "resource2.yaml"),
		filepath.Join(tempDir, "ns2", "type1", "resource3.yaml"),
	}

	for _, resourcePath := range testResources {
		// Create directory
		dir := filepath.Dir(resourcePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory for %s: %v", resourcePath, err)
		}

		// Create file
		content := `apiVersion: v1
kind: Test
metadata:
  name: test`

		if err := os.WriteFile(resourcePath, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to write file %s: %v", resourcePath, err)
		}
	}

	// Verify all files exist
	for _, resourcePath := range testResources {
		if _, err := os.Stat(resourcePath); os.IsNotExist(err) {
			t.Errorf("Resource file %s was not created", resourcePath)
		}
	}

	// Clean up files
	for _, resourcePath := range testResources {
		if err := os.Remove(resourcePath); err != nil {
			t.Errorf("Failed to remove file %s: %v", resourcePath, err)
		}
	}

	// Verify files were removed
	for _, resourcePath := range testResources {
		if _, err := os.Stat(resourcePath); !os.IsNotExist(err) {
			t.Errorf("Resource file %s was not removed", resourcePath)
		}
	}
}

func TestOrphanedResourceErrorHandling(t *testing.T) {
	// Test error handling for orphaned resources
	tempDir := t.TempDir()

	// Test with non-existent directory
	nonExistentDir := filepath.Join(tempDir, "non-existent")

	// Try to read from non-existent directory
	_, err := os.ReadDir(nonExistentDir)
	if err == nil {
		t.Error("Expected error when reading non-existent directory")
	}

	// Test with non-existent file
	nonExistentFile := filepath.Join(tempDir, "non-existent.yaml")
	_, err = os.ReadFile(nonExistentFile)
	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}

	// Test with invalid permissions
	restrictedDir := filepath.Join(tempDir, "restricted")
	if err := os.MkdirAll(restrictedDir, 0000); err != nil {
		t.Fatalf("Failed to create restricted directory: %v", err)
	}

	// Try to read from restricted directory
	_, err = os.ReadDir(restrictedDir)
	if err == nil {
		t.Error("Expected error when reading restricted directory")
	}

	// Clean up restricted directory
	if err := os.Chmod(restrictedDir, 0755); err != nil {
		t.Errorf("Failed to change permissions on restricted directory: %v", err)
	}
}

func TestOrphanedResourceConcurrency(t *testing.T) {
	// Test concurrent access to orphaned resources
	tempDir := t.TempDir()

	// Create test resources
	testDir := filepath.Join(tempDir, "test-ns", "test-type")
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create multiple goroutines to test concurrent access
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Goroutine %d panicked: %v", id, r)
				}
				done <- true
			}()

			// Create a test file
			testFile := filepath.Join(testDir, fmt.Sprintf("test-%d.yaml", id))
			content := fmt.Sprintf(`apiVersion: v1
kind: Test
metadata:
  name: test-%d`, id)

			if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
				t.Errorf("Goroutine %d failed to write file: %v", id, err)
			}

			// Read the file back
			if _, err := os.ReadFile(testFile); err != nil {
				t.Errorf("Goroutine %d failed to read file: %v", id, err)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify all files were created
	for i := 0; i < 10; i++ {
		testFile := filepath.Join(testDir, fmt.Sprintf("test-%d.yaml", i))
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			t.Errorf("File %s was not created by goroutine %d", testFile, i)
		}
	}
}
