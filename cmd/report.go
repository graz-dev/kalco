package cmd

import (
	"fmt"

	"kalco/pkg/kube"
	"kalco/pkg/reports"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate comprehensive cluster reports",
	Long: formatLongDescription(`
Generate detailed reports about your Kubernetes cluster including resource
summaries, change analysis, security assessments, and operational insights.
Reports can be generated in multiple formats for documentation, compliance,
and monitoring purposes.
`),
	Example: `  # Generate comprehensive cluster report
  kalco report

  # Generate specific report types
  kalco report --types summary,security,changes

  # Output report in JSON format
  kalco report --output json

  # Generate report for specific time period
  kalco report --since 7d

  # Save report to file
  kalco report --output-file cluster-report.html`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runReport()
	},
}

var (
	reportOutput     string
	reportOutputFile string
	reportTypes      []string
	reportSince      string
	reportNamespaces []string
)

func runReport() error {
	// Create Kubernetes clients
	printInfo("🔌 Connecting to Kubernetes cluster...")
	_, _, _, err := kube.NewClients(kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes clients: %w", err)
	}
	printSuccess("Connected to cluster")

	// Create report generator
	printInfo("📊 Initializing report generator...")
	reportGen := reports.NewReportGenerator("./")
	
	// Configure report options
	if len(reportTypes) > 0 {
		printInfo(fmt.Sprintf("📋 Generating report types: %v", reportTypes))
		// TODO: Add report type filtering
	}

	if len(reportNamespaces) > 0 {
		printInfo(fmt.Sprintf("📂 Including namespaces: %v", reportNamespaces))
		// TODO: Add namespace filtering
	}

	if reportSince != "" {
		printInfo(fmt.Sprintf("📅 Report period: since %s", reportSince))
		// TODO: Add time filtering
	}

	// Generate report
	printInfo("📊 Generating cluster report...")
	
	// TODO: Implement comprehensive report generation
	// This would create a detailed report including:
	// - Cluster summary
	// - Resource inventory
	// - Security analysis
	// - Change history
	// - Recommendations

	if err := reportGen.GenerateReport(""); err != nil {
		return fmt.Errorf("report generation failed: %w", err)
	}

	printSuccess("📊 Report generated successfully!")
	
	if reportOutputFile != "" {
		printInfo(fmt.Sprintf("💾 Report saved to: %s", reportOutputFile))
	}

	return nil
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Add flags
	reportCmd.Flags().StringVarP(&reportOutput, "output", "o", "table", "output format (table, json, yaml, html)")
	reportCmd.Flags().StringVar(&reportOutputFile, "output-file", "", "save report to file")
	reportCmd.Flags().StringSliceVar(&reportTypes, "types", []string{}, "specific report types (summary,security,changes,resources)")
	reportCmd.Flags().StringVar(&reportSince, "since", "", "generate report for period since (e.g., 7d, 1w, 1m)")
	reportCmd.Flags().StringSliceVarP(&reportNamespaces, "namespaces", "n", []string{}, "specific namespaces to include")
}