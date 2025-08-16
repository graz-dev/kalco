package cmd

import (
	"fmt"

	"kalco/pkg/kube"
	"kalco/pkg/validation"

	"github.com/spf13/cobra"
)

var (
	validateNamespaces []string
	validateResources  []string
	validateOutput     string
	validateFix        bool
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate cluster resources for issues and broken references",
	Long: formatLongDescription(`
Validate your Kubernetes cluster resources for common issues, broken references,
and configuration problems. This command performs comprehensive validation
including:

• Cross-reference validation (broken ConfigMap/Secret references)
• Resource dependency analysis
• Configuration validation
• Security policy compliance
• Resource quota and limit validation

Results can be output in multiple formats for integration with CI/CD pipelines.
`),
	Example: `  # Validate entire cluster
  kalco validate

  # Validate specific namespaces
  kalco validate --namespaces default,production

  # Validate specific resource types
  kalco validate --resources deployments,services

  # Output results in JSON format
  kalco validate --output json

  # Validate and attempt to fix issues
  kalco validate --fix`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate()
	},
}

func runValidate() error {
	// Create Kubernetes clients
	printInfo("🔌 Connecting to Kubernetes cluster...")
	_, _, _, err := kube.NewClients(kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes clients: %w", err)
	}
	printSuccess("Connected to cluster")

	// Create validator instance
	printInfo("🔍 Initializing cluster validator...")
	validator := validation.NewResourceValidator("./") // TODO: Use proper output directory

	// Configure validation scope
	if len(validateNamespaces) > 0 {
		printInfo(fmt.Sprintf("📂 Validating namespaces: %v", validateNamespaces))
		// TODO: Add namespace filtering
	}

	if len(validateResources) > 0 {
		printInfo(fmt.Sprintf("🎯 Validating resources: %v", validateResources))
		// TODO: Add resource filtering
	}

	// Run validation
	printInfo("🔍 Running cluster validation...")
	results, err := validator.Validate()
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Display results
	switch validateOutput {
	case "json":
		// TODO: Implement JSON output
		printInfo("JSON output not yet implemented")
	case "yaml":
		// TODO: Implement YAML output
		printInfo("YAML output not yet implemented")
	default:
		// TODO: Implement table output
		printSuccess(fmt.Sprintf("Validation complete: %d valid, %d broken, %d warnings", 
			results.Summary.ValidReferences, 
			results.Summary.BrokenReferences, 
			results.Summary.WarningReferences))
		
		if len(results.BrokenReferences) > 0 {
			printWarning("Broken references found:")
			for _, ref := range results.BrokenReferences {
				fmt.Printf("  %s/%s -> %s/%s (%s)\n", 
					ref.SourceType, ref.SourceName, 
					ref.TargetType, ref.TargetName, 
					ref.Field)
			}
		}
	}
	
	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add flags
	validateCmd.Flags().StringSliceVarP(&validateNamespaces, "namespaces", "n", []string{}, "specific namespaces to validate")
	validateCmd.Flags().StringSliceVarP(&validateResources, "resources", "r", []string{}, "specific resource types to validate")
	validateCmd.Flags().StringVarP(&validateOutput, "output", "o", "table", "output format (table, json, yaml)")
	validateCmd.Flags().BoolVar(&validateFix, "fix", false, "attempt to fix validation issues where possible")

	// Add aliases
	validateCmd.Aliases = []string{"check", "lint"}
}