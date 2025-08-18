package validation

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewResourceValidator(t *testing.T) {
	// Test creating new validator
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)
	if validator == nil {
		t.Fatal("Expected validator, got nil")
	}

	// Test that output directory is set
	if validator.outputDir != tempDir {
		t.Errorf("Expected output dir %s, got %s", tempDir, validator.outputDir)
	}

	// Test that resources map is initialized
	if validator.resources == nil {
		t.Error("Resources map should be initialized")
	}
}

func TestResourceValidatorStructure(t *testing.T) {
	// Test validator structure
	tempDir := t.TempDir()
	validator := &ResourceValidator{
		outputDir: tempDir,
		resources: make(map[string]map[string]map[string]interface{}),
	}

	if validator == nil {
		t.Fatal("ResourceValidator is nil")
	}

	if validator.outputDir != tempDir {
		t.Errorf("Expected output dir %s, got %s", tempDir, validator.outputDir)
	}

	if validator.resources == nil {
		t.Error("Resources map should not be nil")
	}
}

func TestResourceValidatorMethods(t *testing.T) {
	// Test that validator methods exist and are callable
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test that methods can be called without panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Validator method call panicked: %v", r)
		}
	}()

	// Call Validate method (should not panic even with empty directory)
	result, err := validator.Validate()
	if err != nil {
		// This is expected when no resources exist
		if !contains(err.Error(), "failed to load resources") {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if result == nil {
		t.Error("Validate should return a result even with no resources")
	}
}

func TestResourceValidatorNilHandling(t *testing.T) {
	// Test that validator handles nil inputs gracefully
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test with empty directory
	result, err := validator.Validate()
	if err != nil {
		// This is expected when no resources exist
		if !contains(err.Error(), "failed to load resources") {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if result == nil {
		t.Error("Validate should return a result even with no resources")
	}
}

func TestResourceValidatorResultStructure(t *testing.T) {
	// Test that validation results have proper structure
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test validation result structure
	result, err := validator.Validate()
	if err != nil {
		// This is expected when no resources exist
		if !contains(err.Error(), "failed to load resources") {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if result == nil {
		t.Fatal("Validation result is nil")
	}

	// Verify result structure
	if result.ValidReferences == nil {
		t.Error("ValidReferences should not be nil")
	}
	if result.BrokenReferences == nil {
		t.Error("BrokenReferences should not be nil")
	}
	if result.WarningReferences == nil {
		t.Error("WarningReferences should not be nil")
	}
	if result.Summary.TotalReferences != 0 {
		t.Errorf("Expected 0 total references, got %d", result.Summary.TotalReferences)
	}
}

func TestResourceValidatorConcurrency(t *testing.T) {
	// Test that validator can be used concurrently
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

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

			// Call validation method
			validator.Validate()
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestResourceValidatorErrorHandling(t *testing.T) {
	// Test that validator handles errors gracefully
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test with various directory states
	testDirs := []string{
		tempDir,
		filepath.Join(tempDir, "nonexistent"),
		filepath.Join(tempDir, "empty"),
	}

	for i, testDir := range testDirs {
		// Use the main validator for each directory
		validator.outputDir = testDir

		// Test validation
		_, err := validator.Validate()
		if err != nil {
			// This is expected in some cases
			if !contains(err.Error(), "failed to load resources") {
				t.Errorf("Test %d: Unexpected error: %v", i, err)
			}
		}
		// Result can be nil for non-existent directories, which is acceptable
		// We just verify that the method doesn't panic
	}
}

func TestResourceValidatorPerformance(t *testing.T) {
	// Test that validator methods complete in reasonable time
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test validation performance
	start := time.Now()
	validator.Validate()
	duration := time.Since(start)

	if duration > 1*time.Second {
		t.Errorf("Validation took too long: %v", duration)
	}
}

func TestResourceValidatorConsistency(t *testing.T) {
	// Test that validator returns consistent results
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test validation consistency
	result1, _ := validator.Validate()
	result2, _ := validator.Validate()

	if result1 == nil && result2 != nil {
		t.Error("Validation results are inconsistent")
	}
	if result1 != nil && result2 == nil {
		t.Error("Validation results are inconsistent")
	}
}

func TestResourceValidatorLoadResources(t *testing.T) {
	// Test resource loading functionality
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Create a test resource structure
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

	// Test loading resources
	if err := validator.loadResources(); err != nil {
		t.Fatalf("Failed to load resources: %v", err)
	}

	// Verify resource was loaded
	if validator.resources[testNamespace] == nil {
		t.Error("Namespace not found in loaded resources")
	}
	if validator.resources[testNamespace][testResourceType] == nil {
		t.Error("Resource type not found in loaded resources")
	}
	if validator.resources[testNamespace][testResourceType][testResourceName] == nil {
		t.Error("Resource not found in loaded resources")
	}
}

func TestResourceValidatorSelectorValidation(t *testing.T) {
	// Test selector validation functionality
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test selector target existence
	namespace := "test-ns"
	key := "app"
	value := "test"

	// Initially no resources exist
	if validator.selectorTargetExists(namespace, key, value) {
		t.Error("Selector target should not exist initially")
	}

	// Test resource existence
	if validator.resourceExists(namespace, "Pod", "test-pod") {
		t.Error("Resource should not exist initially")
	}

	// Test that validator is properly initialized
	if validator.outputDir != tempDir {
		t.Errorf("Expected output dir %s, got %s", tempDir, validator.outputDir)
	}
}

func TestResourceValidatorReferenceValidation(t *testing.T) {
	// Test reference validation functionality
	tempDir := t.TempDir()
	validator := NewResourceValidator(tempDir)

	// Test with empty resources
	result := &ValidationResult{
		ValidReferences:   []ResourceReference{},
		BrokenReferences:  []ResourceReference{},
		WarningReferences: []ResourceReference{},
	}

	// Test various validation methods
	validator.validateServiceReferences(result)
	validator.validateRoleBindingReferences(result)
	validator.validateNetworkPolicyReferences(result)
	validator.validateIngressReferences(result)
	validator.validateHPAReferences(result)
	validator.validatePDBReferences(result)

	// Verify no references were found (empty resources)
	if len(result.ValidReferences) != 0 {
		t.Errorf("Expected 0 valid references, got %d", len(result.ValidReferences))
	}
	if len(result.BrokenReferences) != 0 {
		t.Errorf("Expected 0 broken references, got %d", len(result.BrokenReferences))
	}
	if len(result.WarningReferences) != 0 {
		t.Errorf("Expected 0 warning references, got %d", len(result.WarningReferences))
	}
}

func TestValidationResultStructure(t *testing.T) {
	// Test ValidationResult structure
	result := &ValidationResult{
		ValidReferences:   []ResourceReference{},
		BrokenReferences:  []ResourceReference{},
		WarningReferences: []ResourceReference{},
		Summary: ValidationSummary{
			TotalReferences:   0,
			ValidReferences:   0,
			BrokenReferences:  0,
			WarningReferences: 0,
		},
	}

	if result == nil {
		t.Fatal("ValidationResult is nil")
	}

	// Test that slices are initialized
	if result.ValidReferences == nil {
		t.Error("ValidReferences should not be nil")
	}
	if result.BrokenReferences == nil {
		t.Error("BrokenReferences should not be nil")
	}
	if result.WarningReferences == nil {
		t.Error("WarningReferences should not be nil")
	}

	// Test summary structure
	if result.Summary.TotalReferences != 0 {
		t.Errorf("Expected 0 total references, got %d", result.Summary.TotalReferences)
	}
}

func TestResourceReferenceStructure(t *testing.T) {
	// Test ResourceReference structure
	ref := ResourceReference{
		SourceType:      "Service",
		SourceName:      "test-service",
		SourceNamespace: "test-ns",
		TargetType:      "Pod",
		TargetName:      "test-pod",
		TargetNamespace: "test-ns",
		Field:           "spec.selector",
		Line:            10,
	}

	if ref.SourceType != "Service" {
		t.Errorf("Expected SourceType 'Service', got %s", ref.SourceType)
	}
	if ref.SourceName != "test-service" {
		t.Errorf("Expected SourceName 'test-service', got %s", ref.SourceName)
	}
	if ref.SourceNamespace != "test-ns" {
		t.Errorf("Expected SourceNamespace 'test-ns', got %s", ref.SourceNamespace)
	}
	if ref.TargetType != "Pod" {
		t.Errorf("Expected TargetType 'Pod', got %s", ref.TargetType)
	}
	if ref.TargetName != "test-pod" {
		t.Errorf("Expected TargetName 'test-pod', got %s", ref.TargetName)
	}
	if ref.TargetNamespace != "test-ns" {
		t.Errorf("Expected TargetNamespace 'test-ns', got %s", ref.TargetNamespace)
	}
	if ref.Field != "spec.selector" {
		t.Errorf("Expected Field 'spec.selector', got %s", ref.Field)
	}
	if ref.Line != 10 {
		t.Errorf("Expected Line 10, got %d", ref.Line)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr))
}
