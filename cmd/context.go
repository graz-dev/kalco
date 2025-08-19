package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kalco/pkg/context"

	"github.com/spf13/cobra"
)

var (
	contextCmd = &cobra.Command{
		Use:   "context",
		Short: "Manage Kalco contexts",
		Long: `Manage Kalco contexts for different Kubernetes clusters and configurations.

A context defines:
- Kubeconfig file for cluster access
- Output directory for exports
- Dynamic labels for context identification
- Description for context purpose

Examples:
  kalco context set production --kubeconfig ~/.kube/prod-config --output ./prod-exports --labels env=prod,team=platform
  kalco context list
  kalco context use production
  kalco context load ./existing-kalco-export`,
	}

	contextSetCmd = &cobra.Command{
		Use:   "set [name]",
		Short: "Create or update a context",
		Long: `Create or update a Kalco context with the specified configuration.

The context will be saved and can be used for future operations.`,
		Args: cobra.ExactArgs(1),
		RunE: runContextSet,
	}

	contextListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all contexts",
		Long:  `Display all available contexts with their configuration details.`,
		RunE:  runContextList,
	}

	contextUseCmd = &cobra.Command{
		Use:   "use [name]",
		Short: "Switch to a context",
		Long:  `Switch to the specified context. This context will be used for future operations.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runContextUse,
	}

	contextDeleteCmd = &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a context",
		Long:  `Delete the specified context. Cannot delete the currently active context.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runContextDelete,
	}

	contextShowCmd = &cobra.Command{
		Use:   "show [name]",
		Short: "Show context details",
		Long:  `Display detailed information about a specific context.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runContextShow,
	}

	contextCurrentCmd = &cobra.Command{
		Use:   "current",
		Short: "Show current context",
		Long:  `Display information about the currently active context.`,
		RunE:  runContextCurrent,
	}

	contextLoadCmd = &cobra.Command{
		Use:   "load [directory]",
		Short: "Load context from existing kalco directory",
		Long: `Load a context configuration from an existing kalco directory by reading the kalco-config.json file.

This is useful for importing contexts from existing kalco exports or for team collaboration.`,
		Args: cobra.ExactArgs(1),
		RunE: runContextLoad,
	}

	// Flags for context set
	contextKubeConfig  string
	contextOutputDir   string
	contextDescription string
	contextLabels      []string
)

func init() {
	rootCmd.AddCommand(contextCmd)

	// Add subcommands
	contextCmd.AddCommand(contextSetCmd)
	contextCmd.AddCommand(contextListCmd)
	contextCmd.AddCommand(contextUseCmd)
	contextCmd.AddCommand(contextDeleteCmd)
	contextCmd.AddCommand(contextShowCmd)
	contextCmd.AddCommand(contextCurrentCmd)
	contextCmd.AddCommand(contextLoadCmd)

	// Add flags for context set
	contextSetCmd.Flags().StringVar(&contextKubeConfig, "kubeconfig", "", "Path to kubeconfig file")
	contextSetCmd.Flags().StringVar(&contextOutputDir, "output", "", "Output directory for exports")
	contextSetCmd.Flags().StringVar(&contextDescription, "description", "", "Description of the context")
	contextSetCmd.Flags().StringArrayVar(&contextLabels, "labels", []string{}, "Labels in format key=value (can be specified multiple times)")
}

func runContextSet(cmd *cobra.Command, args []string) error {
	name := args[0]

	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	// Parse labels
	labels := make(map[string]string)
	for _, label := range contextLabels {
		parts := strings.SplitN(label, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid label format: %s (expected key=value)", label)
		}
		labels[parts[0]] = parts[1]
	}

	// Set context
	if err := cm.SetContext(name, contextKubeConfig, contextOutputDir, contextDescription, labels); err != nil {
		return fmt.Errorf("failed to set context: %w", err)
	}

	fmt.Printf("Context '%s' set successfully\n", name)
	return nil
}

func runContextList(cmd *cobra.Command, args []string) error {
	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	contexts := cm.ListContexts()
	if len(contexts) == 0 {
		fmt.Println("No contexts found. Use 'kalco context set' to create your first context.")
		return nil
	}

	// Get current context
	current, err := cm.GetCurrentContext()
	currentName := ""
	if err == nil {
		currentName = current.Name
	}

	fmt.Println("Available contexts:")
	fmt.Println()

	for name, ctx := range contexts {
		// Mark current context
		marker := " "
		if name == currentName {
			marker = "*"
		}

		fmt.Printf("%s %s\n", marker, name)
		if ctx.Description != "" {
			fmt.Printf("  Description: %s\n", ctx.Description)
		}
		if ctx.KubeConfig != "" {
			fmt.Printf("  Kubeconfig: %s\n", ctx.KubeConfig)
		}
		if ctx.OutputDir != "" {
			fmt.Printf("  Output Dir: %s\n", ctx.OutputDir)
		}
		if len(ctx.Labels) > 0 {
			labelStrs := make([]string, 0, len(ctx.Labels))
			for k, v := range ctx.Labels {
				labelStrs = append(labelStrs, fmt.Sprintf("%s=%s", k, v))
			}
			fmt.Printf("  Labels: %s\n", strings.Join(labelStrs, ", "))
		}
		fmt.Printf("  Created: %s\n", ctx.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  Updated: %s\n", ctx.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}

	if currentName != "" {
		fmt.Printf("* = current context\n")
	}

	return nil
}

func runContextUse(cmd *cobra.Command, args []string) error {
	name := args[0]

	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	// Use context
	if err := cm.UseContext(name); err != nil {
		return fmt.Errorf("failed to use context: %w", err)
	}

	fmt.Printf("Switched to context '%s'\n", name)
	return nil
}

func runContextDelete(cmd *cobra.Command, args []string) error {
	name := args[0]

	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	// Delete context
	if err := cm.DeleteContext(name); err != nil {
		return fmt.Errorf("failed to delete context: %w", err)
	}

	fmt.Printf("Context '%s' deleted successfully\n", name)
	return nil
}

func runContextShow(cmd *cobra.Command, args []string) error {
	name := args[0]

	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	// Get context
	ctx, err := cm.GetContext(name)
	if err != nil {
		return fmt.Errorf("failed to get context: %w", err)
	}

	// Display context details
	fmt.Printf("Context: %s\n", ctx.Name)
	fmt.Printf("Description: %s\n", ctx.Description)
	fmt.Printf("Kubeconfig: %s\n", ctx.KubeConfig)
	fmt.Printf("Output Directory: %s\n", ctx.OutputDir)

	if len(ctx.Labels) > 0 {
		fmt.Println("Labels:")
		for k, v := range ctx.Labels {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}

	fmt.Printf("Created: %s\n", ctx.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", ctx.UpdatedAt.Format("2006-01-02 15:04:05"))

	return nil
}

func runContextCurrent(cmd *cobra.Command, args []string) error {
	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("no context is currently active. Use 'kalco context use <name>' to switch to a context")
	}

	// Get current context
	current, err := cm.GetCurrentContext()
	if err != nil {
		return fmt.Errorf("no context is currently active. Use 'kalco context use <name>' to switch to a context")
	}

	// Display current context
	fmt.Printf("Current context: %s\n", current.Name)
	fmt.Printf("Description: %s\n", current.Description)
	fmt.Printf("Kubeconfig: %s\n", current.KubeConfig)
	fmt.Printf("Output Directory: %s\n", current.OutputDir)

	if len(current.Labels) > 0 {
		fmt.Println("Labels:")
		for k, v := range current.Labels {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}

	fmt.Printf("Created: %s\n", current.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", current.UpdatedAt.Format("2006-01-02 15:04:05"))

	return nil
}

func runContextLoad(cmd *cobra.Command, args []string) error {
	directory := args[0]

	// Check if directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", directory)
	}

	// Check if it's a kalco directory by looking for kalco-config.json
	configFile := filepath.Join(directory, "kalco-config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' is not a valid kalco directory (missing kalco-config.json)", directory)
	}

	// Read and parse kalco-config.json
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read kalco config file: %w", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse kalco config file: %w", err)
	}

	// Extract context information
	contextName, ok := config["context_name"].(string)
	if !ok || contextName == "" {
		return fmt.Errorf("invalid or missing context_name in kalco-config.json")
	}

	kubeconfig, _ := config["kubeconfig"].(string)
	description, _ := config["description"].(string)

	// Parse labels
	var labels map[string]string
	if labelsData, ok := config["labels"]; ok {
		if labelsMap, ok := labelsData.(map[string]interface{}); ok {
			labels = make(map[string]string)
			for k, v := range labelsMap {
				if strVal, ok := v.(string); ok {
					labels[k] = strVal
				}
			}
		}
	}

	// Get config directory
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Create context manager
	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return fmt.Errorf("failed to create context manager: %w", err)
	}

	// Set context
	if err := cm.SetContext(contextName, kubeconfig, directory, description, labels); err != nil {
		return fmt.Errorf("failed to set context: %w", err)
	}

	fmt.Printf("Context '%s' loaded successfully from '%s'\n", contextName, directory)
	fmt.Printf("   Kubeconfig: %s\n", kubeconfig)
	fmt.Printf("   Output Dir: %s\n", directory)
	if description != "" {
		fmt.Printf("   Description: %s\n", description)
	}
	if len(labels) > 0 {
		labelStrs := make([]string, 0, len(labels))
		for k, v := range labels {
			labelStrs = append(labelStrs, fmt.Sprintf("%s=%s", k, v))
		}
		fmt.Printf("   Labels: %s\n", strings.Join(labelStrs, ", "))
	}

	return nil
}
