package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommand(t *testing.T) {
	// Test that root command exists and has correct properties
	if rootCmd.Use != "kalco" {
		t.Errorf("Expected root command use to be 'kalco', got '%s'", rootCmd.Use)
	}

	if rootCmd.Short == "" {
		t.Error("Expected root command to have a short description")
	}

	if rootCmd.Long == "" {
		t.Error("Expected root command to have a long description")
	}

	// Test that root command has the expected subcommands
	expectedSubcommands := []string{"context", "export", "version"}
	actualSubcommands := make([]string, 0, len(rootCmd.Commands()))
	for _, cmd := range rootCmd.Commands() {
		actualSubcommands = append(actualSubcommands, cmd.Name())
	}

	for _, expected := range expectedSubcommands {
		found := false
		for _, actual := range actualSubcommands {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected subcommand '%s' not found", expected)
		}
	}

	// Test that root command has persistent flags
	if rootCmd.PersistentFlags().Lookup("kubeconfig") == nil {
		t.Error("Expected root command to have kubeconfig persistent flag")
	}

	if rootCmd.PersistentFlags().Lookup("verbose") == nil {
		t.Error("Expected root command to have verbose persistent flag")
	}

	if rootCmd.PersistentFlags().Lookup("no-color") == nil {
		t.Error("Expected root command to have no-color persistent flag")
	}
}

func TestContextCommand(t *testing.T) {
	// Find context command
	var contextCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "context" {
			contextCmd = cmd
			break
		}
	}

	if contextCmd == nil {
		t.Fatal("Context command not found")
	}

	// Test context command properties
	if contextCmd.Use != "context" {
		t.Errorf("Expected context command use to be 'context', got '%s'", contextCmd.Use)
	}

	if contextCmd.Short == "" {
		t.Error("Expected context command to have a short description")
	}

	// Test that context command has the expected subcommands
	expectedSubcommands := []string{"set", "list", "use", "delete", "show", "current", "load"}
	actualSubcommands := make([]string, 0, len(contextCmd.Commands()))
	for _, cmd := range contextCmd.Commands() {
		actualSubcommands = append(actualSubcommands, cmd.Name())
	}

	for _, expected := range expectedSubcommands {
		found := false
		for _, actual := range actualSubcommands {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected context subcommand '%s' not found", expected)
		}
	}
}

func TestExportCommand(t *testing.T) {
	// Find export command
	var exportCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "export" {
			exportCmd = cmd
			break
		}
	}

	if exportCmd == nil {
		t.Fatal("Export command not found")
	}

	// Test export command properties
	if exportCmd.Use != "export" {
		t.Errorf("Expected export command use to be 'export', got '%s'", exportCmd.Use)
	}

	if exportCmd.Short == "" {
		t.Error("Expected export command to have a short description")
	}

	// Test that export command has the expected flags
	expectedFlags := []string{"output", "git-push", "commit-message", "namespaces", "resources", "exclude", "dry-run", "no-commit"}
	for _, expected := range expectedFlags {
		if exportCmd.Flags().Lookup(expected) == nil {
			t.Errorf("Expected export command to have flag '%s'", expected)
		}
	}

	// Test aliases
	expectedAliases := []string{"dump", "backup"}
	for _, expected := range expectedAliases {
		found := false
		for _, actual := range exportCmd.Aliases {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected export command alias '%s' not found", expected)
		}
	}
}

func TestVersionCommand(t *testing.T) {
	// Find version command
	var versionCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "version" {
			versionCmd = cmd
			break
		}
	}

	if versionCmd == nil {
		t.Fatal("Version command not found")
	}

	// Test version command properties
	if versionCmd.Use != "version" {
		t.Errorf("Expected version command use to be 'version', got '%s'", versionCmd.Use)
	}

	if versionCmd.Short == "" {
		t.Error("Expected version command to have a short description")
	}
}

func TestCommandStructure(t *testing.T) {
	// Test that all commands have proper help text
	testCommandHelp(t, rootCmd)
}

func testCommandHelp(t *testing.T, cmd *cobra.Command) {
	// Test that command has help text
	if cmd.Short == "" {
		t.Errorf("Command '%s' missing short description", cmd.Name())
	}

	// Test that command has long description if it's a leaf command
	if len(cmd.Commands()) == 0 && cmd.Long == "" {
		t.Errorf("Leaf command '%s' missing long description", cmd.Name())
	}

	// Recursively test subcommands
	for _, subcmd := range cmd.Commands() {
		testCommandHelp(t, subcmd)
	}
}

func TestFlagConsistency(t *testing.T) {
	// Test that root command has persistent flags
	if rootCmd.PersistentFlags().Lookup("kubeconfig") == nil {
		t.Error("Root command missing kubeconfig persistent flag")
	}

	if rootCmd.PersistentFlags().Lookup("verbose") == nil {
		t.Error("Root command missing verbose persistent flag")
	}

	if rootCmd.PersistentFlags().Lookup("no-color") == nil {
		t.Error("Root command missing no-color persistent flag")
	}
}
