package context

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"os/exec"

	"gopkg.in/yaml.v3"
)

// Context represents a Kalco context configuration
type Context struct {
	Name        string            `json:"name" yaml:"name"`
	KubeConfig  string            `json:"kubeconfig" yaml:"kubeconfig"`
	OutputDir   string            `json:"output_dir" yaml:"output_dir"`
	Labels      map[string]string `json:"labels" yaml:"labels"`
	Description string            `json:"description" yaml:"description"`
	CreatedAt   time.Time         `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" yaml:"updated_at"`
}

// ContextManager handles context operations
type ContextManager struct {
	configDir string
	contexts  map[string]*Context
	current   string
}

// NewContextManager creates a new context manager
func NewContextManager(configDir string) (*ContextManager, error) {
	cm := &ContextManager{
		configDir: configDir,
		contexts:  make(map[string]*Context),
	}

	// Ensure config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	// Load existing contexts
	if err := cm.loadContexts(); err != nil {
		return nil, fmt.Errorf("failed to load contexts: %w", err)
	}

	// Load current context
	if err := cm.loadCurrentContext(); err != nil {
		return nil, fmt.Errorf("failed to load current context: %w", err)
	}

	return cm, nil
}

// SetContext creates or updates a context
func (cm *ContextManager) SetContext(name, kubeconfig, outputDir, description string, labels map[string]string) error {
	if name == "" {
		return fmt.Errorf("context name cannot be empty")
	}

	// Create output directory if specified
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory '%s': %w", outputDir, err)
		}

		// Initialize Git repository and create kalco-config.json
		if err := cm.initializeKalcoDirectory(outputDir, name, kubeconfig, labels, description); err != nil {
			return fmt.Errorf("failed to initialize kalco directory: %w", err)
		}
	}

	now := time.Now()
	context := &Context{
		Name:        name,
		KubeConfig:  kubeconfig,
		OutputDir:   outputDir,
		Labels:      labels,
		Description: description,
		UpdatedAt:   now,
	}

	// If context exists, preserve creation time
	if existing, exists := cm.contexts[name]; exists {
		context.CreatedAt = existing.CreatedAt
	} else {
		context.CreatedAt = now
	}

	cm.contexts[name] = context

	// Save contexts
	if err := cm.saveContexts(); err != nil {
		return fmt.Errorf("failed to save contexts: %w", err)
	}

	return nil
}

// GetContext retrieves a context by name
func (cm *ContextManager) GetContext(name string) (*Context, error) {
	context, exists := cm.contexts[name]
	if !exists {
		return nil, fmt.Errorf("context '%s' not found", name)
	}
	return context, nil
}

// ListContexts returns all available contexts
func (cm *ContextManager) ListContexts() map[string]*Context {
	return cm.contexts
}

// UseContext sets the current active context
func (cm *ContextManager) UseContext(name string) error {
	if _, exists := cm.contexts[name]; !exists {
		return fmt.Errorf("context '%s' not found", name)
	}

	cm.current = name

	// Save current context
	if err := cm.saveCurrentContext(); err != nil {
		return fmt.Errorf("failed to save current context: %w", err)
	}

	return nil
}

// GetCurrentContext returns the current active context
func (cm *ContextManager) GetCurrentContext() (*Context, error) {
	if cm.current == "" {
		return nil, fmt.Errorf("no context is currently active")
	}

	return cm.GetContext(cm.current)
}

// DeleteContext removes a context
func (cm *ContextManager) DeleteContext(name string) error {
	if _, exists := cm.contexts[name]; !exists {
		return fmt.Errorf("context '%s' not found", name)
	}

	// If deleting current context, clear it first
	if cm.current == name {
		cm.current = ""
		// Save the cleared current context
		if err := cm.saveCurrentContext(); err != nil {
			return fmt.Errorf("failed to clear current context: %w", err)
		}
	}

	delete(cm.contexts, name)

	// Save contexts
	if err := cm.saveContexts(); err != nil {
		return fmt.Errorf("failed to save contexts: %w", err)
	}

	return nil
}

// loadContexts loads contexts from disk
func (cm *ContextManager) loadContexts() error {
	contextsFile := filepath.Join(cm.configDir, "contexts.yaml")

	if _, err := os.Stat(contextsFile); os.IsNotExist(err) {
		return nil // No contexts file yet
	}

	data, err := os.ReadFile(contextsFile)
	if err != nil {
		return fmt.Errorf("failed to read contexts file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cm.contexts); err != nil {
		return fmt.Errorf("failed to unmarshal contexts: %w", err)
	}

	return nil
}

// saveContexts saves contexts to disk
func (cm *ContextManager) saveContexts() error {
	contextsFile := filepath.Join(cm.configDir, "contexts.yaml")

	data, err := yaml.Marshal(cm.contexts)
	if err != nil {
		return fmt.Errorf("failed to marshal contexts: %w", err)
	}

	if err := os.WriteFile(contextsFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write contexts file: %w", err)
	}

	return nil
}

// loadCurrentContext loads the current context from disk
func (cm *ContextManager) loadCurrentContext() error {
	currentFile := filepath.Join(cm.configDir, "current-context")

	if _, err := os.Stat(currentFile); os.IsNotExist(err) {
		return nil // No current context file yet
	}

	data, err := os.ReadFile(currentFile)
	if err != nil {
		return fmt.Errorf("failed to read current context file: %w", err)
	}

	cm.current = string(data)
	return nil
}

// saveCurrentContext saves the current context to disk
func (cm *ContextManager) saveCurrentContext() error {
	currentFile := filepath.Join(cm.configDir, "current-context")

	if err := os.WriteFile(currentFile, []byte(cm.current), 0644); err != nil {
		return fmt.Errorf("failed to write current context file: %w", err)
	}

	return nil
}

// GetConfigDir returns the configuration directory
func (cm *ContextManager) GetConfigDir() string {
	return cm.configDir
}

// ValidateContext validates a context configuration
func (cm *ContextManager) ValidateContext(context *Context) error {
	if context.Name == "" {
		return fmt.Errorf("context name cannot be empty")
	}

	if context.KubeConfig != "" {
		if _, err := os.Stat(context.KubeConfig); os.IsNotExist(err) {
			return fmt.Errorf("kubeconfig file '%s' does not exist", context.KubeConfig)
		}
	}

	if context.OutputDir != "" {
		// Ensure output directory can be created
		dir := context.OutputDir
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("cannot create output directory '%s': %w", dir, err)
		}
	}

	return nil
}

// initializeKalcoDirectory creates the kalco-config.json file and initializes Git repository
func (cm *ContextManager) initializeKalcoDirectory(outputDir, contextName, kubeconfig string, labels map[string]string, description string) error {
	// Create kalco-config.json
	config := map[string]interface{}{
		"context_name": contextName,
		"kubeconfig":   kubeconfig,
		"labels":       labels,
		"description":  description,
		"created_at":   time.Now().Format(time.RFC3339),
		"version":      "1.0",
	}

	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal kalco config: %w", err)
	}

	configFile := filepath.Join(outputDir, "kalco-config.json")
	if err := os.WriteFile(configFile, configData, 0644); err != nil {
		return fmt.Errorf("failed to write kalco config file: %w", err)
	}

	// Initialize Git repository if not already initialized
	gitDir := filepath.Join(outputDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		// Initialize Git repository
		cmd := exec.Command("git", "init")
		cmd.Dir = outputDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to initialize Git repository: %w", err)
		}

		// Add and commit initial files
		cmd = exec.Command("git", "add", ".")
		cmd.Dir = outputDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to add files to Git: %w", err)
		}

		cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Initial kalco context: %s", contextName))
		cmd.Dir = outputDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to commit initial files: %w", err)
		}
	}

	return nil
}
