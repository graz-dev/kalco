package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"kalco/pkg/diff"

	"github.com/spf13/cobra"
)

var (
	validateBaseRef    string
	validateSnapshot   string
	validateDirs       []string
	validateNamespace  string
	validateReportFile string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate local manifests against snapshot excluding PR changes (GitOps mode)",
	Long: `Automatically detects changed files between current branch and base branch,
fetches the base version for smart drift detection, and compares against the snapshot.
This encapsulates the logic previously done in CI scripts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate()
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVar(&validateBaseRef, "base", "origin/master", "Base reference or branch to compare against (e.g. origin/main)")
	validateCmd.Flags().StringVar(&validateSnapshot, "snapshot-dir", "", "Directory containing snapshots")
	validateCmd.Flags().StringSliceVar(&validateDirs, "dirs", []string{"."}, "Directories to scan for changes")
	validateCmd.Flags().StringVarP(&validateNamespace, "namespace", "n", "", "Override namespace for resources")
	validateCmd.Flags().StringVar(&validateReportFile, "report-file", "", "File to write the drift report to (markdown)")

	validateCmd.MarkFlagRequired("snapshot-dir")
}

func runValidate() error {
	// 1. Fetch base ref to ensure we can diff
	// We assume 'git' is in path
	printInfo(fmt.Sprintf("Fetching %s...", validateBaseRef))

	// Split remote and branch if possible, but git fetch origin master works if refs match.
	// Safer to just try fetching the ref spec directly if it looks like a remote/branch.
	// Or just 'git fetch' all?
	// Let's rely on user having fetched or fetch explicitly if it's "origin/master".
	parts := strings.Split(validateBaseRef, "/")
	if len(parts) > 1 {
		remote := parts[0]
		branch := strings.Join(parts[1:], "/")
		exec.Command("git", "fetch", remote, branch).Run()
	}

	// 2. Detect changed files
	printInfo("Detecting changed files...")
	cmd := exec.Command("git", "diff", "--name-only", validateBaseRef, "--")
	cmd.Args = append(cmd.Args, validateDirs...)

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("git diff failed: %w", err)
	}

	files := strings.Split(strings.TrimSpace(string(out)), "\n")
	var changedFiles []string
	for _, f := range files {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}
		if strings.HasSuffix(f, ".yaml") || strings.HasSuffix(f, ".yml") {
			// Check if file exists (it implies it wasn't deleted)
			if _, err := os.Stat(f); err == nil {
				changedFiles = append(changedFiles, f)
			}
		}
	}

	if len(changedFiles) == 0 {
		printSuccess("No changed YAML manifests detected.")
		return nil
	}

	printInfo(fmt.Sprintf("Found %d changed files.", len(changedFiles)))

	// 3. Init Differ
	// Enable IgnoreExtraFields by default for validation mode as it's safer for GitOps
	differ := diff.NewDiffer(true)

	totalDiff := ""
	hasDrift := false

	// 4. Process each file
	for _, file := range changedFiles {
		printInfo(fmt.Sprintf("Validating %s...", file))

		// Load Local
		localResources, err := differ.LoadLocalResource(file)
		if err != nil {
			printError(fmt.Sprintf("Failed to load local file %s: %v", file, err))
			continue
		}

		// Load Base Content (git show base:file)
		var baseResources []*diff.Resource
		baseContentCmd := exec.Command("git", "show", fmt.Sprintf("%s:%s", validateBaseRef, file))
		baseOut, err := baseContentCmd.Output()
		if err == nil {
			// Base exists
			baseResources, _ = differ.ParseResources(bytes.NewReader(baseOut), "base")
		} else {
			// New file in PR, base is empty
		}

		fileDrift := false
		fileOutput := ""

		for _, local := range localResources {
			// Load Snapshot
			remote, err := differ.LoadSnapshotResource(validateSnapshot, local.Object, validateNamespace)
			if err != nil {
				printError(fmt.Sprintf("Failed to load snapshot for %s: %v", local.Object.GetName(), err))
				continue
			}

			// Find matching base
			var base *diff.Resource
			if len(baseResources) > 0 {
				for _, b := range baseResources {
					if b.Object.GetKind() == local.Object.GetKind() &&
						b.Object.GetName() == local.Object.GetName() {
						base = b
						break
					}
				}
			}

			// Diff
			result, err := differ.Diff(local, remote, base)
			if err != nil {
				printError(fmt.Sprintf("Diff failed for %s: %v", local.Object.GetName(), err))
				continue
			}

			if result.HasDrift {
				fileDrift = true
				hasDrift = true
			}

			// Format Output
			header := ""
			if len(localResources) > 1 {
				header = fmt.Sprintf("\n--- %s/%s ---\n", local.Object.GetKind(), local.Object.GetName())
			}

			// Simple Markdown formatting
			if result.Diff != "" {
				fileOutput += header + "```diff\n" + result.Diff + "\n```\n"
			} else if result.IsNew {
				fileOutput += header + "> [!NOTE]\n> Start tracking new resource\n"
			}
		}

		if fileOutput != "" {
			totalDiff += fmt.Sprintf("\n#### File: `%s`\n%s\n", file, fileOutput)
			if fileDrift {
				totalDiff += "\n> [!CAUTION]\n> **Drift Detected**\n"
				printError(fmt.Sprintf("Drift detected in %s", file))
			} else {
				printSuccess(fmt.Sprintf("No drift in %s", file))
			}
		} else {
			printSuccess(fmt.Sprintf("Clean: %s", file))
		}
	}

	// Output Report
	if validateReportFile != "" {
		if totalDiff == "" {
			totalDiff = "No drift detected."
		}
		if err := os.WriteFile(validateReportFile, []byte(totalDiff), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to write report: %v\n", err)
		}
	} else {
		// If no file, print report to stdout (if not empty)
		if totalDiff != "" {
			fmt.Println(totalDiff)
		}
	}

	if hasDrift {
		return fmt.Errorf("Validation failed: Drift detected.")
	}

	return nil
}
