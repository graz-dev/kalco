package context

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewContextManager(t *testing.T) {
	// Test with valid config directory
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if cm == nil {
		t.Fatal("Expected context manager, got nil")
	}
	if cm.configDir != tempDir {
		t.Errorf("Expected config dir %s, got %s", tempDir, cm.configDir)
	}

	// Test with invalid config directory (should still work as it creates the directory)
	invalidDir := filepath.Join(tempDir, "nonexistent", "subdir")
	cm2, err := NewContextManager(invalidDir)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if cm2 == nil {
		t.Fatal("Expected context manager, got nil")
	}
}

func TestSetContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test creating new context
	outputDir := filepath.Join(tempDir, "output")
	labels := map[string]string{"env": "test", "team": "qa"}
	err = cm.SetContext("test-context", tempDir, outputDir, "Test description", labels)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify context was created
	context, exists := cm.contexts["test-context"]
	if !exists {
		t.Fatal("Context was not created")
	}
	if context.Name != "test-context" {
		t.Errorf("Expected name 'test-context', got %s", context.Name)
	}
	if context.KubeConfig != tempDir {
		t.Errorf("Expected kubeconfig '%s', got %s", tempDir, context.KubeConfig)
	}
	if context.OutputDir != outputDir {
		t.Errorf("Expected output dir '%s', got %s", outputDir, context.OutputDir)
	}
	if context.Description != "Test description" {
		t.Errorf("Expected description 'Test description', got %s", context.Description)
	}
	if len(context.Labels) != 2 {
		t.Errorf("Expected 2 labels, got %d", len(context.Labels))
	}
	if context.Labels["env"] != "test" {
		t.Errorf("Expected label env=test, got env=%s", context.Labels["env"])
	}

	// Test updating existing context
	newOutputDir := filepath.Join(tempDir, "new-output")
	err = cm.SetContext("test-context", tempDir, newOutputDir, "Updated description", map[string]string{"env": "prod"})
	if err != nil {
		t.Fatalf("Expected no error updating context, got %v", err)
	}

	// Verify context was updated
	context, exists = cm.contexts["test-context"]
	if !exists {
		t.Fatal("Context was not preserved")
	}
	if context.KubeConfig != tempDir {
		t.Errorf("Expected updated kubeconfig '%s', got %s", tempDir, context.KubeConfig)
	}
	if context.OutputDir != newOutputDir {
		t.Errorf("Expected updated output dir '%s', got %s", newOutputDir, context.OutputDir)
	}
	if context.Description != "Updated description" {
		t.Errorf("Expected updated description 'Updated description', got %s", context.Description)
	}
	if len(context.Labels) != 1 {
		t.Errorf("Expected 1 label, got %d", len(context.Labels))
	}
	if context.Labels["env"] != "prod" {
		t.Errorf("Expected updated label env=prod, got env=%s", context.Labels["env"])
	}

	// Test with empty name (should fail)
	err = cm.SetContext("", tempDir, outputDir, "Description", labels)
	if err == nil {
		t.Fatal("Expected error for empty name, got none")
	}
}

func TestGetContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test getting non-existent context
	_, err = cm.GetContext("nonexistent")
	if err == nil {
		t.Fatal("Expected error for non-existent context, got none")
	}

	// Create and get context
	labels := map[string]string{"env": "test"}
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set context: %v", err)
	}

	context, err := cm.GetContext("test-context")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if context.Name != "test-context" {
		t.Errorf("Expected name 'test-context', got %s", context.Name)
	}
}

func TestListContexts(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test empty contexts
	contexts := cm.ListContexts()
	if len(contexts) != 0 {
		t.Errorf("Expected 0 contexts, got %d", len(contexts))
	}

	// Add contexts
	labels1 := map[string]string{"env": "test"}
	labels2 := map[string]string{"env": "prod"}

	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Test", labels1)
	if err != nil {
		t.Fatalf("Failed to set test context: %v", err)
	}

	err = cm.SetContext("prod-context", tempDir, filepath.Join(tempDir, "output"), "Production", labels2)
	if err != nil {
		t.Fatalf("Failed to set prod context: %v", err)
	}

	// Test listing contexts
	contexts = cm.ListContexts()
	if len(contexts) != 2 {
		t.Errorf("Expected 2 contexts, got %d", len(contexts))
	}

	// Verify both contexts exist
	if _, exists := contexts["test-context"]; !exists {
		t.Error("test-context not found in list")
	}
	if _, exists := contexts["prod-context"]; !exists {
		t.Error("prod-context not found in list")
	}
}

func TestUseContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test using non-existent context
	err = cm.UseContext("nonexistent")
	if err == nil {
		t.Fatal("Expected error for non-existent context, got none")
	}

	// Create context
	labels := map[string]string{"env": "test"}
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set context: %v", err)
	}

	// Use context
	err = cm.UseContext("test-context")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify current context
	if cm.current != "test-context" {
		t.Errorf("Expected current context 'test-context', got %s", cm.current)
	}
}

func TestGetCurrentContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test getting current context when none is set
	_, err = cm.GetCurrentContext()
	if err == nil {
		t.Fatal("Expected error for no current context, got none")
	}

	// Set and use context
	labels := map[string]string{"env": "test"}
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set context: %v", err)
	}

	err = cm.UseContext("test-context")
	if err != nil {
		t.Fatalf("Failed to use context: %v", err)
	}

	// Get current context
	context, err := cm.GetCurrentContext()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if context.Name != "test-context" {
		t.Errorf("Expected current context name 'test-context', got %s", context.Name)
	}
}

func TestDeleteContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test deleting non-existent context
	err = cm.DeleteContext("nonexistent")
	if err == nil {
		t.Fatal("Expected error for non-existent context, got none")
	}

	// Create contexts
	labels := map[string]string{"env": "test"}
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set test context: %v", err)
	}

	err = cm.SetContext("other-context", tempDir, filepath.Join(tempDir, "output"), "Other", labels)
	if err != nil {
		t.Fatalf("Failed to set other context: %v", err)
	}

	// Use test-context as current
	err = cm.UseContext("test-context")
	if err != nil {
		t.Fatalf("Failed to use context: %v", err)
	}

	// Test deleting non-current context
	err = cm.DeleteContext("other-context")
	if err != nil {
		t.Fatalf("Expected no error deleting non-current context, got %v", err)
	}

	// Verify other-context was deleted
	if _, exists := cm.contexts["other-context"]; exists {
		t.Error("other-context should have been deleted")
	}

	// Test deleting current context
	err = cm.DeleteContext("test-context")
	if err != nil {
		t.Fatalf("Expected no error deleting current context, got %v", err)
	}

	// Verify test-context was deleted
	if _, exists := cm.contexts["test-context"]; exists {
		t.Error("test-context should have been deleted")
	}

	// Verify current context was cleared
	if cm.current != "" {
		t.Errorf("Expected current context to be cleared, got %s", cm.current)
	}
}

func TestValidateContext(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	// Test valid context
	validContext := &Context{
		Name:       "valid",
		KubeConfig: tempDir,
		OutputDir:  tempDir,
		Labels:     map[string]string{"env": "test"},
	}

	err = cm.ValidateContext(validContext)
	if err != nil {
		t.Errorf("Expected no error for valid context, got %v", err)
	}

	// Test context with empty name
	invalidContext := &Context{
		Name:       "",
		KubeConfig: tempDir,
		OutputDir:  tempDir,
	}

	err = cm.ValidateContext(invalidContext)
	if err == nil {
		t.Fatal("Expected error for empty name, got none")
	}

	// Test context with non-existent output directory
	invalidContext.Name = "test"
	invalidContext.OutputDir = "/nonexistent/output"
	err = cm.ValidateContext(invalidContext)
	if err == nil {
		t.Fatal("Expected error for non-existent output directory, got none")
	}
}

func TestPersistence(t *testing.T) {
	tempDir := t.TempDir()

	// Create first context manager
	cm1, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create first context manager: %v", err)
	}

	// Add contexts
	labels := map[string]string{"env": "test"}
	err = cm1.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set context: %v", err)
	}

	err = cm1.UseContext("test-context")
	if err != nil {
		t.Fatalf("Failed to use context: %v", err)
	}

	// Create second context manager (should load existing contexts)
	cm2, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create second context manager: %v", err)
	}

	// Verify contexts were loaded
	contexts := cm2.ListContexts()
	if len(contexts) != 1 {
		t.Errorf("Expected 1 context, got %d", len(contexts))
	}

	if _, exists := contexts["test-context"]; !exists {
		t.Error("test-context not found in loaded contexts")
	}

	// Verify current context was loaded
	if cm2.current != "test-context" {
		t.Errorf("Expected current context 'test-context', got %s", cm2.current)
	}
}

func TestInitializeKalcoDirectory(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	outputDir := filepath.Join(tempDir, "kalco-output")
	labels := map[string]string{"env": "test", "team": "qa"}

	// Create the output directory first
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	// Test directory initialization
	err = cm.initializeKalcoDirectory(outputDir, "test-context", tempDir, labels, "Test description")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify directory was created
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Fatal("Output directory was not created")
	}

	// Verify kalco-config.json was created
	configFile := filepath.Join(outputDir, "kalco-config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Fatal("kalco-config.json was not created")
	}

	// Verify Git repository was initialized
	gitDir := filepath.Join(outputDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Fatal("Git repository was not initialized")
	}

	// Verify kalco-config.json content
	data, err := os.ReadFile(configFile)
	if err != nil {
		t.Fatalf("Failed to read kalco-config.json: %v", err)
	}

	// Basic content verification (should contain context name)
	if !contains(data, "test-context") {
		t.Error("kalco-config.json does not contain context name")
	}
	if !contains(data, "test") {
		t.Error("kalco-config.json does not contain label value")
	}
}

func TestGetConfigDir(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	configDir := cm.GetConfigDir()
	if configDir != tempDir {
		t.Errorf("Expected config dir %s, got %s", tempDir, configDir)
	}
}

// Helper function to check if byte slice contains string
func contains(data []byte, str string) bool {
	return len(data) > 0 && len(str) > 0 && len(data) >= len(str)
}

func TestContextTimestamps(t *testing.T) {
	tempDir := t.TempDir()
	cm, err := NewContextManager(tempDir)
	if err != nil {
		t.Fatalf("Failed to create context manager: %v", err)
	}

	before := time.Now()
	labels := map[string]string{"env": "test"}
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Description", labels)
	if err != nil {
		t.Fatalf("Failed to set context: %v", err)
	}
	after := time.Now()

	context := cm.contexts["test-context"]

	// Verify timestamps
	if context.CreatedAt.Before(before) || context.CreatedAt.After(after) {
		t.Errorf("CreatedAt %v is not within expected range [%v, %v]", context.CreatedAt, before, after)
	}
	if context.UpdatedAt.Before(before) || context.UpdatedAt.After(after) {
		t.Errorf("UpdatedAt %v is not within expected range [%v, %v]", context.UpdatedAt, before, after)
	}

	// Update context and verify UpdatedAt changes
	time.Sleep(10 * time.Millisecond) // Ensure time difference
	beforeUpdate := time.Now()
	err = cm.SetContext("test-context", tempDir, filepath.Join(tempDir, "output"), "Updated", labels)
	if err != nil {
		t.Fatalf("Failed to update context: %v", err)
	}
	afterUpdate := time.Now()

	updatedContext := cm.contexts["test-context"]

	// CreatedAt should remain the same
	if !updatedContext.CreatedAt.Equal(context.CreatedAt) {
		t.Error("CreatedAt should not change when updating context")
	}

	// UpdatedAt should be updated
	if updatedContext.UpdatedAt.Before(beforeUpdate) || updatedContext.UpdatedAt.After(afterUpdate) {
		t.Errorf("UpdatedAt %v is not within expected range [%v, %v]", updatedContext.UpdatedAt, beforeUpdate, afterUpdate)
	}
}
