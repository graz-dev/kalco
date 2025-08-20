package cmd

import (
	"fmt"
	"strings"
	"time"

	"kalco/pkg/dumper"
	"kalco/pkg/git"
	"kalco/pkg/kube"
	"kalco/pkg/reports"

	"github.com/spf13/cobra"
)

var (
	exportOutputDir     string
	exportGitPush       bool
	exportCommitMessage string
	exportNamespaces    []string
	exportResources     []string
	exportExclude       []string
	exportDryRun        bool
	exportNoCommit      bool
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
	printCommandHeader("CLUSTER EXPORT", "Exporting Kubernetes resources to organized YAML files")

	// Require active context
	requireActiveContext()

	// Create Kubernetes clients
	printInfo("Connecting to Kubernetes cluster...")

	// Use context kubeconfig if available, otherwise use flag
	activeContext, err := getActiveContext()
	if err != nil {
		return fmt.Errorf("failed to get active context: %w", err)
	}

	kubeconfigPath := kubeconfig
	if activeContext.KubeConfig != "" {
		kubeconfigPath = activeContext.KubeConfig
	}

	clientset, discoveryClient, dynamicClient, err := kube.NewClients(kubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes clients: %w", err)
	}
	printSuccess("Connected to cluster")

	// Create dumper instance
	d := dumper.NewDumper(clientset, discoveryClient)
	d.SetDynamicClient(dynamicClient)

	// Configure dumper options
	if len(exportNamespaces) > 0 {
		printInfo(fmt.Sprintf("Filtering namespaces: %s",
			colorize(ColorYellow, strings.Join(exportNamespaces, ", "))))
	}

	if len(exportResources) > 0 {
		printInfo(fmt.Sprintf("Filtering resources: %s",
			colorize(ColorYellow, strings.Join(exportResources, ", "))))
	}

	if len(exportExclude) > 0 {
		printInfo(fmt.Sprintf("Excluding resources: %s",
			colorize(ColorRed, strings.Join(exportExclude, ", "))))
	}

	// Use context output directory if available, otherwise use flag
	outputDir := exportOutputDir
	if activeContext.OutputDir != "" {
		outputDir = activeContext.OutputDir
	}

	if exportDryRun {
		printWarning("Dry run mode - no files will be written")
		printInfo(fmt.Sprintf("Would export to: %s", colorize(ColorCyan, outputDir)))
		return nil
	}

	printSeparator()

	// Execute the main dump function
	printSubHeader("Resource Discovery & Export")
	printInfo("Discovering available API resources...")
	printInfo("Building directory structure...")
	printInfo("Exporting resources...")

	if err := d.DumpAllResources(outputDir); err != nil {
		return fmt.Errorf("failed to export resources: %w", err)
	}

	printSuccess("Resource export completed")

	// Handle Git repository operations
	if !exportNoCommit {
		printSeparator()
		printSubHeader("Git Integration")
		printInfo("Setting up Git repository...")

		gitRepo := git.NewGitRepo(outputDir)
		if err := gitRepo.SetupAndCommit(exportCommitMessage, exportGitPush); err != nil {
			printWarning(fmt.Sprintf("Git operations failed: %v", err))
		} else {
			printSuccess("Git repository updated")
			if exportGitPush {
				printSuccess("Changes pushed to remote origin")
			}
		}
	}

	// Generate change report
	printSeparator()
	printSubHeader("Report Generation")
	printInfo("Generating cluster analysis report...")

	reportGen := reports.NewReportGenerator(outputDir)
	if err := reportGen.GenerateReport(exportCommitMessage); err != nil {
		printWarning(fmt.Sprintf("Report generation failed: %v", err))
	} else {
		printSuccess("Analysis report generated")
	}

	// Success summary
	printSeparator()
	printHeader("EXPORT COMPLETE")

	fmt.Printf("Resources exported to: %s\n",
		colorize(ColorCyan+ColorBold, outputDir))
	fmt.Println()

	printInfo("Your cluster snapshot is ready for:")
	fmt.Printf("   %s Backup and disaster recovery\n", colorize(ColorGreen, "•"))
	fmt.Printf("   %s Resource auditing and compliance\n", colorize(ColorGreen, "•"))
	fmt.Printf("   %s Development environment replication\n", colorize(ColorGreen, "•"))

	return nil
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Generate default output directory with timestamp
	timestamp := time.Now().Format("20060102-150405")
	defaultOutputDir := "./kalco-export-" + timestamp

	// Add flags
	exportCmd.Flags().StringVarP(&exportOutputDir, "output", "o", defaultOutputDir, "output directory path")
	exportCmd.Flags().BoolVar(&exportGitPush, "git-push", false, "automatically push changes to remote origin")
	exportCmd.Flags().StringVarP(&exportCommitMessage, "commit-message", "m", "", "custom Git commit message")
	exportCmd.Flags().StringSliceVarP(&exportNamespaces, "namespaces", "n", []string{}, "specific namespaces to export (comma-separated)")
	exportCmd.Flags().StringSliceVarP(&exportResources, "resources", "r", []string{}, "specific resource types to export (comma-separated)")
	exportCmd.Flags().StringSliceVar(&exportExclude, "exclude", []string{}, "resource types to exclude (comma-separated)")
	exportCmd.Flags().BoolVar(&exportDryRun, "dry-run", false, "show what would be exported without writing files")
	exportCmd.Flags().BoolVar(&exportNoCommit, "no-commit", false, "skip Git commit operations")

	// Add aliases
	exportCmd.Aliases = []string{"dump", "backup"}
}
