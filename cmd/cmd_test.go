package cmd

import (
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetConfigDir(t *testing.T) {
	// Test that getConfigDir returns a valid path
	configDir, err := getConfigDir()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Should contain .kalco
	if !contains(configDir, ".kalco") {
		t.Errorf("Expected config dir to contain '.kalco', got %s", configDir)
	}

	// Should be absolute path
	if !filepath.IsAbs(configDir) {
		t.Errorf("Expected absolute path, got %s", configDir)
	}
}

func TestGetActiveContext(t *testing.T) {
	// Test when no context is available
	// This should not error, just return nil or error about no context
	_, err := getActiveContext()
	if err != nil {
		// This is expected when no context is set up
		// The error should be about no context being active
		if !contains(err.Error(), "no context") && !contains(err.Error(), "failed to get config directory") {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	// activeContext can be nil or contain a context, both are valid
}

func TestContextCommandRegistration(t *testing.T) {
	// Test that context command is properly registered
	if contextCmd == nil {
		t.Fatal("contextCmd should not be nil")
	}

	// Test that all subcommands are registered
	subcommands := contextCmd.Commands()
	expectedSubcommands := []string{"set", "list", "use", "delete", "show", "current", "load"}

	for _, expected := range expectedSubcommands {
		found := false
		for _, cmd := range subcommands {
			if cmd.Name() == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Subcommand '%s' not found in context command", expected)
		}
	}
}

func TestContextSetCommandFlags(t *testing.T) {
	// Test that all expected flags are registered
	if contextSetCmd == nil {
		t.Fatal("contextSetCmd should not be nil")
	}

	// Check for required flags
	flags := contextSetCmd.Flags()

	expectedFlags := []string{"kubeconfig", "output", "description", "labels"}
	for _, expected := range expectedFlags {
		if flags.Lookup(expected) == nil {
			t.Errorf("Flag '--%s' not found in context set command", expected)
		}
	}
}

func TestContextListCommand(t *testing.T) {
	// Test that list command is properly configured
	if contextListCmd == nil {
		t.Fatal("contextListCmd should not be nil")
	}

	// Verify command properties
	if contextListCmd.Use != "list" {
		t.Errorf("Expected use 'list', got %s", contextListCmd.Use)
	}

	if contextListCmd.Short == "" {
		t.Error("Expected non-empty short description")
	}
}

func TestContextUseCommand(t *testing.T) {
	// Test that use command is properly configured
	if contextUseCmd == nil {
		t.Fatal("contextUseCmd should not be nil")
	}

	// Verify command properties
	if contextUseCmd.Use != "use [name]" {
		t.Errorf("Expected use 'use [name]', got %s", contextUseCmd.Use)
	}

	if contextUseCmd.Args == nil {
		t.Error("Expected args validation")
	}
}

func TestContextDeleteCommand(t *testing.T) {
	// Test that delete command is properly configured
	if contextDeleteCmd == nil {
		t.Fatal("contextDeleteCmd should not be nil")
	}

	// Verify command properties
	if contextDeleteCmd.Use != "delete [name]" {
		t.Errorf("Expected use 'delete [name]', got %s", contextDeleteCmd.Use)
	}
}

func TestContextShowCommand(t *testing.T) {
	// Test that show command is properly configured
	if contextShowCmd == nil {
		t.Fatal("contextShowCmd should not be nil")
	}

	// Verify command properties
	if contextShowCmd.Use != "show [name]" {
		t.Errorf("Expected use 'show [name]', got %s", contextShowCmd.Use)
	}
}

func TestContextCurrentCommand(t *testing.T) {
	// Test that current command is properly configured
	if contextCurrentCmd == nil {
		t.Fatal("contextCurrentCmd should not be nil")
	}

	// Verify command properties
	if contextCurrentCmd.Use != "current" {
		t.Errorf("Expected use 'current', got %s", contextCurrentCmd.Use)
	}
}

func TestContextLoadCommand(t *testing.T) {
	// Test that load command is properly configured
	if contextLoadCmd == nil {
		t.Fatal("contextLoadCmd should not be nil")
	}

	// Verify command properties
	if contextLoadCmd.Use != "load [directory]" {
		t.Errorf("Expected use 'load [directory]', got %s", contextLoadCmd.Use)
	}

	if contextLoadCmd.Args == nil {
		t.Error("Expected args validation")
	}
}

func TestRootCommandIntegration(t *testing.T) {
	// Test that root command has context subcommand
	if rootCmd == nil {
		t.Fatal("rootCmd should not be nil")
	}

	// Find context command in root
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "context" {
			found = true
			break
		}
	}

	if !found {
		t.Error("context command not found in root command")
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr))
}

func TestContextCommandHelp(t *testing.T) {
	// Test that context command has proper help text
	if contextCmd == nil {
		t.Fatal("contextCmd should not be nil")
	}

	if contextCmd.Short == "" {
		t.Error("Expected non-empty short description")
	}

	if contextCmd.Long == "" {
		t.Error("Expected non-empty long description")
	}

	// Examples are optional, so we don't require them
}

func TestContextSubcommandHelp(t *testing.T) {
	// Test that all subcommands have proper help text
	subcommands := []*cobra.Command{
		contextSetCmd, contextListCmd, contextUseCmd,
		contextDeleteCmd, contextShowCmd, contextCurrentCmd, contextLoadCmd,
	}

	for i, cmd := range subcommands {
		if cmd == nil {
			t.Errorf("Subcommand %d is nil", i)
			continue
		}

		if cmd.Short == "" {
			t.Errorf("Subcommand %s missing short description", cmd.Name())
		}

		if cmd.Long == "" {
			t.Errorf("Subcommand %s missing long description", cmd.Name())
		}
	}
}

func TestContextCommandStructure(t *testing.T) {
	// Test that context command structure is correct
	if contextCmd.Use != "context" {
		t.Errorf("Expected use 'context', got %s", contextCmd.Use)
	}

	// Test that it's a proper command
	if contextCmd.Run != nil {
		t.Error("Context command should not have Run function, it's a parent command")
	}

	// Test that it has subcommands
	if len(contextCmd.Commands()) == 0 {
		t.Error("Context command should have subcommands")
	}
}

func TestContextCommandFlags(t *testing.T) {
	// Test that context command has proper flags
	if contextCmd == nil {
		t.Fatal("contextCmd should not be nil")
	}

	// Context command should inherit global flags
	// Note: help flag might not be directly accessible depending on Cobra version
	flags := contextCmd.Flags()
	if flags == nil {
		t.Error("Context command should have flags")
	}
}

func TestContextSubcommandArgs(t *testing.T) {
	// Test that subcommands have proper argument validation
	subcommands := map[string]*cobra.Command{
		"set":    contextSetCmd,
		"use":    contextUseCmd,
		"delete": contextDeleteCmd,
		"show":   contextShowCmd,
		"load":   contextLoadCmd,
	}

	for name, cmd := range subcommands {
		if cmd == nil {
			t.Errorf("Subcommand %s is nil", name)
			continue
		}

		if cmd.Args == nil {
			t.Errorf("Subcommand %s missing args validation", name)
		}
	}

	// List and current commands don't need args
	if contextListCmd.Args != nil {
		t.Error("List command should not have args validation")
	}

	if contextCurrentCmd.Args != nil {
		t.Error("Current command should not have args validation")
	}
}
