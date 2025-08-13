package cmd

import (
	"fmt"
	"os"
	"time"

	"kalco/pkg/dumper"
	"kalco/pkg/git"
	"kalco/pkg/kube"
	"kalco/pkg/reports"

	"github.com/spf13/cobra"
)

var (
	kubeconfig    string
	outputDir     string
	gitPush       bool
	commitMessage string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kalco",
	Short: "🚀 Kubernetes Cluster Resource Dumper",
	Long: `
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🚀 KALCO - Kubernetes Cluster Resource Dumper 🚀

🎯 Comprehensive cluster resource extraction tool
🔍 Discovers ALL resources (native + CRDs) automatically
📁 Creates organized YAML exports with clean directory structure
🧹 Cleans metadata for easy re-application
🌐 Works both in-cluster and out-of-cluster
⚡ Fast, efficient, and production-ready

Perfect for:
  • Cluster backups and disaster recovery
  • Resource auditing and compliance
  • Development environment replication
  • Documentation and resource cataloging

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create Kubernetes clients
		clientset, discoveryClient, dynamicClient, err := kube.NewClients(kubeconfig)
		if err != nil {
			return fmt.Errorf("failed to create Kubernetes clients: %w", err)
		}

		// Create dumper instance
		d := dumper.NewDumper(clientset, discoveryClient)
		d.SetDynamicClient(dynamicClient)

		// Execute the main dump function
		fmt.Println("🚀 Starting Kubernetes cluster resource dump...")
		fmt.Println("🔍 Discovering resources and building directory structure...")

		if err := d.DumpAllResources(outputDir); err != nil {
			return fmt.Errorf("❌ failed to dump resources: %w", err)
		}

		// Handle Git repository operations
		fmt.Println("📦 Setting up Git repository for version control...")
		gitRepo := git.NewGitRepo(outputDir)
		if err := gitRepo.SetupAndCommit(commitMessage, gitPush); err != nil {
			fmt.Printf("⚠️  Warning: Git operations failed: %v\n", err)
			fmt.Println("  Continuing without Git version control...")
		}

		// Generate change report
		fmt.Println("📊 Generating cluster change report...")
		reportGen := reports.NewReportGenerator(outputDir)
		if err := reportGen.GenerateReport(commitMessage); err != nil {
			fmt.Printf("⚠️  Warning: Report generation failed: %v\n", err)
			fmt.Println("  Continuing without change report...")
		}

		fmt.Println()
		fmt.Println("🎉 SUCCESS! 🎉")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		fmt.Printf("📁 All resources have been successfully dumped to:\n")
		fmt.Printf("   %s\n", outputDir)
		fmt.Println()
		fmt.Println("🎯 Your cluster snapshot is ready for:")
		fmt.Println("   • Backup and disaster recovery")
		fmt.Println("   • Resource auditing and compliance")
		fmt.Println("   • Development environment replication")
		fmt.Println("   • Documentation and resource cataloging")
		fmt.Println()
		if gitRepo.IsGitRepo() {
			fmt.Println("📦 Git repository initialized/updated for version control")
			if gitRepo.HasRemoteOrigin() {
				if gitPush {
					fmt.Println("🚀 Changes pushed to remote origin")
				} else {
					fmt.Println("💡 Use --git-push flag to automatically push changes")
				}
			}
		}
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		return nil
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
	// Generate default output directory with timestamp
	timestamp := time.Now().Format("20060102-150405")
	defaultOutputDir := "./kalco-dump-" + timestamp

	// Add persistent flags
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "path to the kubeconfig file (optional)")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", defaultOutputDir, "path to the output directory")
	rootCmd.PersistentFlags().BoolVar(&gitPush, "git-push", false, "automatically push changes to remote origin if available")
	rootCmd.PersistentFlags().StringVar(&commitMessage, "commit-message", "", "custom commit message (default: timestamp-based message)")
}
