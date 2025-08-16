package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage kalco configuration",
	Long: formatLongDescription(`
Manage kalco configuration settings. Configuration can be stored globally
or per-project to customize default behavior, output formats, and preferences.
`),
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize kalco configuration",
	Long: formatLongDescription(`
Initialize a new kalco configuration file with default settings.
This creates a .kalco.yaml file in the current directory or updates
the global configuration.
`),
	Example: `  # Initialize config in current directory
  kalco config init

  # Initialize global config
  kalco config init --global

  # Initialize with custom settings
  kalco config init --template advanced`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigInit()
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long: formatLongDescription(`
Display the current kalco configuration settings, including both
global and project-specific configurations.
`),
	Example: `  # Show current config
  kalco config show

  # Show config in JSON format
  kalco config show --output json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigShow()
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long: formatLongDescription(`
Set a specific configuration value. Values can be set globally
or for the current project.
`),
	Example: `  # Set default output directory
  kalco config set output.directory ./backups

  # Set global default
  kalco config set --global output.format yaml

  # Set namespace filter
  kalco config set filters.namespaces "default,kube-system"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigSet(args[0], args[1])
	},
}

var (
	configGlobal   bool
	configTemplate string
	configOutput   string
)

func runConfigInit() error {
	configPath := ".kalco.yaml"
	if configGlobal {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		configPath = filepath.Join(homeDir, ".kalco", "config.yaml")
		
		// Create directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		printWarning(fmt.Sprintf("Configuration file already exists: %s", configPath))
		return nil
	}

	// Create default configuration
	defaultConfig := getDefaultConfig(configTemplate)
	
	if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	printSuccess(fmt.Sprintf("Configuration initialized: %s", configPath))
	return nil
}

func runConfigShow() error {
	printInfo("Current kalco configuration:")
	
	// TODO: Implement config loading and display
	printWarning("Configuration display not yet implemented")
	
	return nil
}

func runConfigSet(key, value string) error {
	printInfo(fmt.Sprintf("Setting %s = %s", key, value))
	
	// TODO: Implement config setting
	printWarning("Configuration setting not yet implemented")
	
	return nil
}

func getDefaultConfig(template string) string {
	switch template {
	case "advanced":
		return `# Kalco Advanced Configuration
output:
  directory: "./kalco-export-{{.Date}}"
  format: "yaml"
  git:
    enabled: true
    auto_push: false
    commit_message: "Kalco export {{.Date}}"

filters:
  namespaces: []
  resources: []
  exclude: ["events", "replicasets"]

validation:
  enabled: true
  strict: false
  
analysis:
  orphaned_resources: true
  security_scan: false
  
reports:
  enabled: true
  formats: ["html", "json"]
  
ui:
  colors: true
  verbose: false
  progress: true
`
	default:
		return `# Kalco Configuration
output:
  directory: "./kalco-export-{{.Date}}"
  format: "yaml"

filters:
  exclude: ["events"]

ui:
  colors: true
  verbose: false
`
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
	
	// Add subcommands
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)

	// Flags for init command
	configInitCmd.Flags().BoolVar(&configGlobal, "global", false, "initialize global configuration")
	configInitCmd.Flags().StringVar(&configTemplate, "template", "default", "configuration template (default, advanced)")

	// Flags for show command
	configShowCmd.Flags().StringVarP(&configOutput, "output", "o", "yaml", "output format (yaml, json)")

	// Flags for set command
	configSetCmd.Flags().BoolVar(&configGlobal, "global", false, "set global configuration")
}