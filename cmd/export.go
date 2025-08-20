package cmd

import (
	"fmt"
	"time"

	"kalco/pkg/dumper"
	"kalco/pkg/git"
	"kalco/pkg/kube"
	"kalco/pkg/reports"

	"github.com/spf13/cobra"
)

var (
	exportGitPush       bool
	exportCommitMessage string
	exportDryRun        bool
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export cluster resources to organized YAML files",
	Long: formatLongDescription(`
Export all Kubernetes resources from your cluster into organized YAML files.
This command discovers all available API resources (including CRDs) and exports
them with clean metadata suitable for re-application.

The export creates an intuitive directory structure:
  • Namespaced resources: <output>/<namespace>/<kind>/<name>.yaml
  • Cluster resources: <output>/_cluster/<kind>/<name>.yaml

Includes automatic Git integration for version control and change tracking.
`),

	RunE: func(cmd *cobra.Command, args []string) error {
		return runExport()
	},
}

func runExport() error {
	// Require active context
	requireActiveContext()

	// Create Kubernetes clients
	printInfo("Connecting to Kubernetes cluster...")

	// Use context kubeconfig (always required)
	activeContext, err := getActiveContext()
	if err != nil {
		return fmt.Errorf("failed to get active context: %w", err)
	}

	kubeconfigPath := activeContext.KubeConfig
	if kubeconfigPath == "" {
		return fmt.Errorf("context must have a kubeconfig configured")
	}

	clientset, discoveryClient, dynamicClient, err := kube.NewClients(kubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes clients: %w", err)
	}

	// Get cluster information
	serverVersion, err := discoveryClient.ServerVersion()
	if err != nil {
		printWarning("Could not retrieve cluster version information")
	} else {
		printClusterInfo("Connected", "Kubernetes API", serverVersion.String())
	}

	// Create dumper instance
	d := dumper.NewDumper(clientset, discoveryClient)
	d.SetDynamicClient(dynamicClient)

	// Set output callback for resource export
	d.SetOutputCallback(func(level, message string) {
		switch level {
		case "SUCCESS":
			printSuccess(message)
		case "WARNING":
			printWarning(message)
		case "ERROR":
			printError(message)
		default:
			printInfo(message)
		}
	})

	// Use context output directory (always required)
	outputDir := activeContext.OutputDir
	if outputDir == "" {
		return fmt.Errorf("context must have an output directory configured")
	}

	if exportDryRun {
		printWarning("Dry run mode - no files will be written")
		printInfo(fmt.Sprintf("Would export to %s", outputDir))
		return nil
	}

	printSeparator()

	// Execute the main dump function
	if err := d.DumpAllResources(outputDir); err != nil {
		return fmt.Errorf("failed to export resources: %w", err)
	}

	printSuccess("Resource export completed")

	// Handle Git repository operations (always commit)
	printSeparator()
	commitMsg := exportCommitMessage
	if commitMsg == "" {
		commitMsg = fmt.Sprintf("Kalco export: %s", time.Now().Format("2006-01-02 15:04:05"))
	}

	gitRepo := git.NewGitRepo(outputDir)
	if err := gitRepo.SetupAndCommit(commitMsg, exportGitPush); err != nil {
		printWarning(fmt.Sprintf("Git operations failed: %v", err))
	} else {
		printSuccess("Git repository updated")
		if exportGitPush {
			printSuccess("Changes pushed to remote origin")
		}
	}

	// Generate change report
	printSeparator()
	reportGen := reports.NewReportGenerator(outputDir)
	if err := reportGen.GenerateReport(commitMsg); err != nil {
		printWarning(fmt.Sprintf("Report generation failed: %v", err))
	} else {
		printSuccess("Analysis report generated")
	}

	// Success summary
	printSuccess(fmt.Sprintf("Export completed successfully to %s", outputDir))

	return nil
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Add flags
	exportCmd.Flags().BoolVar(&exportGitPush, "git-push", false, "automatically push changes to remote origin")
	exportCmd.Flags().StringVarP(&exportCommitMessage, "commit-message", "m", "", "custom Git commit message")
	exportCmd.Flags().BoolVar(&exportDryRun, "dry-run", false, "show what would be exported without writing files")

	// Add aliases
	exportCmd.Aliases = []string{"dump", "backup"}
}
