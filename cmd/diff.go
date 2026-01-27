package cmd

import (
	"fmt"

	"kalco/pkg/diff"
	"kalco/pkg/dumper"
	"kalco/pkg/kube"

	"github.com/spf13/cobra"
)

var (
	diffTarget            string
	diffSnapshotDir       string
	diffIgnoreBase        string
	diffIgnoreExtraFields bool // New flag
	diffRefresh           bool
	diffFormat            string
	diffNamespace         string
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Compare local manifest with snapshot or live cluster state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDiff()
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flags().StringVar(&diffTarget, "target", "", "Path to local manifest file")
	diffCmd.Flags().StringVar(&diffSnapshotDir, "snapshot-dir", "", "Root directory of snapshot repository")
	diffCmd.Flags().BoolVar(&diffRefresh, "refresh", false, "Fetch live resource from K8s instead of snapshot")
	diffCmd.Flags().StringVar(&diffFormat, "format", "text", "Output format (text|markdown)")
	diffCmd.Flags().StringVarP(&diffNamespace, "namespace", "n", "", "Override namespace for local resources if missing")
	diffCmd.Flags().StringVar(&diffIgnoreBase, "ignore-base", "", "Path to base manifest file for exclusion of intentional changes")
	diffCmd.Flags().BoolVar(&diffIgnoreExtraFields, "ignore-extra-fields", false, "Ignore fields in snapshot that are not present in local manifest (e.g. K8s defaults)")

	diffCmd.MarkFlagRequired("target")
	diffCmd.MarkFlagRequired("snapshot-dir")
}

func runDiff() error {
	// 1. Refresh/Export if requested
	if diffRefresh {
		printInfo("Refreshing snapshot via export...")
		requireActiveContext()
		activeContext, err := getActiveContext()
		if err != nil {
			return err
		}

		clientset, discoveryClient, dynamicClient, err := kube.NewClients(activeContext.KubeConfig)
		if err != nil {
			return fmt.Errorf("failed to create K8s clients: %w", err)
		}

		d := dumper.NewDumper(clientset, discoveryClient)
		d.SetDynamicClient(dynamicClient)

		// Use diffSnapshotDir as the output for this "refresh" export
		// The user specified this dir, so we update it.
		// Note: This exports ALL resources.
		if err := d.DumpAllResources(diffSnapshotDir); err != nil {
			return fmt.Errorf("failed to refresh snapshot: %w", err)
		}
		printSuccess("Snapshot refreshed")
	}

	// 2. Initialize Differ
	differ := diff.NewDiffer(diffIgnoreExtraFields)

	// 3. Load local resources (slice)
	localResources, err := differ.LoadLocalResource(diffTarget)
	if err != nil {
		return fmt.Errorf("failed to load target files: %w", err)
	}

	// Load base resources if provided
	var baseResources []*diff.Resource
	if diffIgnoreBase != "" {
		baseResources, err = differ.LoadLocalResource(diffIgnoreBase)
		if err != nil {
			// Log warning but don't fail? Or fail? Fail is safer to ensure validation logic holds.
			return fmt.Errorf("failed to load base file: %w", err)
		}
	}

	for _, localResource := range localResources {
		// 4. Get comparison resource (Snapshot)
		// We always compare against snapshot (which might have been just refreshed)
		// Pass diffNamespace override
		remoteResource, err := differ.LoadSnapshotResource(diffSnapshotDir, localResource.Object, diffNamespace)
		if err != nil {
			return fmt.Errorf("failed to load snapshot: %w", err)
		}

		// Find matching base resource if available
		var baseResource *diff.Resource
		if len(baseResources) > 0 {
			for _, b := range baseResources {
				// Simple GVK + Name Match
				if b.Object.GetKind() == localResource.Object.GetKind() &&
					b.Object.GetName() == localResource.Object.GetName() {
					baseResource = b
					break
				}
			}
		}

		// 5. Perform Diff (3-way)
		result, err := differ.Diff(localResource, remoteResource, baseResource)
		if err != nil {
			return fmt.Errorf("diff failed: %w", err)
		}

		// 6. Print Result
		if len(localResources) > 1 {
			// Print header for muli-doc files
			kind := localResource.Object.GetKind()
			name := localResource.Object.GetName()
			fmt.Printf("--- %s/%s ---\n", kind, name)
		}
		diff.PrintDiff(result, diffFormat)
	}

	return nil

	return nil
}
