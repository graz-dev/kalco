package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kalco/pkg/context"
)

// getConfigDir returns the Kalco configuration directory
func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".kalco")
	return configDir, nil
}

// getActiveContext returns the currently active context if available
func getActiveContext() (*context.Context, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	cm, err := context.NewContextManager(configDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create context manager: %w", err)
	}

	return cm.GetCurrentContext()
}

// requireActiveContext ensures that an active context exists before executing a command
// This function will exit the program if no context is active
func requireActiveContext() {
	activeContext, err := getActiveContext()
	if err != nil {
		printError("No active context found")
		printInfo("You must set and activate a context before running this command")
		printInfo("")
		printInfo("Available commands to manage contexts:")
		printInfo("  kalco context set <name> --output <dir> [--kubeconfig <path>]")
		printInfo("  kalco context use <name>")
		printInfo("  kalco context list")
		printInfo("")
		printInfo("Example:")
		printInfo("  kalco context set production --output ./prod-exports --kubeconfig ~/.kube/prod-config")
		printInfo("  kalco context use production")
		printInfo("  kalco export")
		printInfo("")
		os.Exit(1)
	}

	// Context is active, print info
	printInfo(fmt.Sprintf("Using context: %s", colorize(ColorCyan, activeContext.Name)))
	if activeContext.Description != "" {
		printInfo(fmt.Sprintf("   Description: %s", activeContext.Description))
	}
	if activeContext.KubeConfig != "" {
		printInfo(fmt.Sprintf("   Kubeconfig: %s", activeContext.KubeConfig))
	}
	if activeContext.OutputDir != "" {
		printInfo(fmt.Sprintf("   Output Dir: %s", activeContext.OutputDir))
	}
	if len(activeContext.Labels) > 0 {
		labelStrs := make([]string, 0, len(activeContext.Labels))
		for k, v := range activeContext.Labels {
			labelStrs = append(labelStrs, fmt.Sprintf("%s=%s", k, v))
		}
		printInfo(fmt.Sprintf("   Labels: %s", strings.Join(labelStrs, ", ")))
	}
	printSeparator()
}
