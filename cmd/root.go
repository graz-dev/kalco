package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Global flags
var (
	kubeconfig string
	verbose    bool
	noColor    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kalco",
	Short: "🚀 Kubernetes Analysis & Lifecycle Control",
	Long: formatLongDescription(`
Kalco is a powerful CLI tool for comprehensive Kubernetes cluster analysis, 
resource extraction, validation, and lifecycle management.

Extract, validate, analyze, and version control your entire cluster with 
comprehensive validation and Git integration.
`),
	Example: `  # Export entire cluster to timestamped directory
  kalco export

  # Export to specific directory with custom options
  kalco export --output ./my-backup --git-push

  # Validate cluster resources for issues
  kalco validate

  # Analyze cluster for orphaned resources
  kalco analyze orphaned

  # Generate cluster report
  kalco report --format json

  # List all available resources in cluster
  kalco resources list`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add persistent flags
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
	
	// Set custom help template
	rootCmd.SetHelpTemplate(getHelpTemplate())
}

// formatLongDescription formats the long description with proper styling
func formatLongDescription(desc string) string {
	if noColor {
		return strings.TrimSpace(desc)
	}
	
	lines := strings.Split(strings.TrimSpace(desc), "\n")
	var formatted []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			formatted = append(formatted, "")
			continue
		}
		formatted = append(formatted, "  "+line)
	}
	
	return strings.Join(formatted, "\n")
}

// getHelpTemplate returns a custom help template with better styling
func getHelpTemplate() string {
	return `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
}
