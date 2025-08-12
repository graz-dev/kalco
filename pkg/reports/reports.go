package reports

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ReportGenerator handles the creation of cluster change reports
type ReportGenerator struct {
	outputDir string
	repoPath  string
}

// NewReportGenerator creates a new ReportGenerator instance
func NewReportGenerator(outputDir string) *ReportGenerator {
	return &ReportGenerator{
		outputDir: outputDir,
		repoPath:  outputDir,
	}
}

// GenerateReport creates a comprehensive markdown report of cluster changes
func (r *ReportGenerator) GenerateReport(commitMessage string) error {
	// Create reports directory
	reportsDir := filepath.Join(r.outputDir, "kalco-reports")
	if err := os.MkdirAll(reportsDir, 0755); err != nil {
		return fmt.Errorf("failed to create reports directory: %w", err)
	}

	// Generate filename from commit message
	filename := r.generateFilename(commitMessage)
	reportPath := filepath.Join(reportsDir, filename)

	// Generate report content
	content, err := r.generateReportContent(commitMessage)
	if err != nil {
		return fmt.Errorf("failed to generate report content: %w", err)
	}

	// Write report to file
	if err := os.WriteFile(reportPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}

	fmt.Printf("  📊 Generated change report: %s\n", filename)
	return nil
}

// generateFilename creates a filename from the commit message
func (r *ReportGenerator) generateFilename(commitMessage string) string {
	// Use timestamp if no custom message
	if commitMessage == "" {
		commitMessage = fmt.Sprintf("Cluster snapshot %s", time.Now().Format("2006-01-02-15-04-05"))
	}

	// Clean the message for filename
	filename := strings.ReplaceAll(commitMessage, " ", "-")
	filename = strings.ReplaceAll(filename, ":", "-")
	filename = strings.ReplaceAll(filename, "/", "-")
	filename = strings.ReplaceAll(filename, "\\", "-")
	filename = strings.ReplaceAll(filename, "*", "-")
	filename = strings.ReplaceAll(filename, "?", "-")
	filename = strings.ReplaceAll(filename, "\"", "-")
	filename = strings.ReplaceAll(filename, "<", "-")
	filename = strings.ReplaceAll(filename, ">", "-")
	filename = strings.ReplaceAll(filename, "|", "-")

	// Limit length and add extension
	if len(filename) > 100 {
		filename = filename[:100]
	}

	return filename + ".md"
}

// generateReportContent creates the markdown content for the report
func (r *ReportGenerator) generateReportContent(commitMessage string) (string, error) {
	var content strings.Builder

	// Header
	content.WriteString("# Cluster Change Report\n\n")
	content.WriteString(fmt.Sprintf("**Generated**: %s\n", time.Now().Format("2006-01-02 15:04:05 UTC")))
	content.WriteString(fmt.Sprintf("**Commit Message**: %s\n\n", commitMessage))

	// Check if this is a Git repository
	if !r.IsGitRepo() {
		content.WriteString("## Initial Snapshot\n\n")
		content.WriteString("This is the first export of the cluster. All resources have been captured.\n\n")
		content.WriteString("### Resource Summary\n")
		content.WriteString("- Complete cluster snapshot\n")
		content.WriteString("- All namespaces exported\n")
		content.WriteString("- All resource types captured\n")
		content.WriteString("- Git repository initialized\n\n")
		return content.String(), nil
	}

	// Get Git information
	commitHash, err := r.getCurrentCommitHash()
	if err != nil {
		content.WriteString("## Error Getting Git Information\n\n")
		content.WriteString(fmt.Sprintf("Failed to retrieve Git information: %v\n\n", err))
		return content.String(), nil
	}

	content.WriteString(fmt.Sprintf("**Commit Hash**: `%s`\n\n", commitHash))

	// Check if there are previous commits
	prevCommit, err := r.getPreviousCommitHash()
	if err != nil {
		content.WriteString("## Initial Snapshot\n\n")
		content.WriteString("This is the first export of the cluster. All resources have been captured.\n\n")
		content.WriteString("### Resource Summary\n")
		content.WriteString("- Complete cluster snapshot\n")
		content.WriteString("- All namespaces exported\n")
		content.WriteString("- All resource types captured\n")
		content.WriteString("- Git repository initialized\n\n")
		return content.String(), nil
	}

	// Generate change report
	content.WriteString("## Changes Since Previous Snapshot\n\n")
	content.WriteString(fmt.Sprintf("**Previous Commit**: `%s`\n\n", prevCommit))

	// Get changed files
	changedFiles, err := r.getChangedFiles(prevCommit, commitHash)
	if err != nil {
		content.WriteString("### Error Getting Changes\n\n")
		content.WriteString(fmt.Sprintf("Failed to retrieve changes: %v\n\n", err))
		return content.String(), nil
	}

	if len(changedFiles) == 0 {
		content.WriteString("### No Changes Detected\n\n")
		content.WriteString("No changes were detected between snapshots.\n\n")
		return content.String(), nil
	}

	// Categorize changes
	changes := r.categorizeChanges(changedFiles)

	// Write change summary
	content.WriteString("### Change Summary\n\n")
	content.WriteString(fmt.Sprintf("- **Total Files Changed**: %d\n", len(changedFiles)))
	content.WriteString(fmt.Sprintf("- **Namespaces Affected**: %d\n", len(changes.Namespaces)))
	content.WriteString(fmt.Sprintf("- **Resource Types Changed**: %d\n", len(changes.ResourceTypes)))
	content.WriteString(fmt.Sprintf("- **New Resources**: %d\n", changes.NewResources))
	content.WriteString(fmt.Sprintf("- **Modified Resources**: %d\n", changes.ModifiedResources))
	content.WriteString(fmt.Sprintf("- **Deleted Resources**: %d\n", changes.DeletedResources))
	content.WriteString("\n")

	// Write detailed changes
	content.WriteString("### Detailed Changes\n\n")

	// Group by namespace
	for namespace, resources := range changes.ByNamespace {
		if namespace == "_cluster" {
			content.WriteString("#### 🌐 Cluster-Scoped Resources\n\n")
		} else {
			content.WriteString(fmt.Sprintf("#### 📁 Namespace: `%s`\n\n", namespace))
		}

		for resourceType, files := range resources {
			content.WriteString(fmt.Sprintf("**%s**:\n", resourceType))
			for _, file := range files {
				status := r.getFileStatus(file, prevCommit, commitHash)
				content.WriteString(fmt.Sprintf("- %s `%s`\n", status, filepath.Base(file)))
			}
			content.WriteString("\n")
		}
	}

	// Write resource type summary
	content.WriteString("### Resource Type Summary\n\n")
	for resourceType, count := range changes.ResourceTypes {
		content.WriteString(fmt.Sprintf("- **%s**: %d changes\n", resourceType, count))
	}
	content.WriteString("\n")

	// Write Git commands for reference
	content.WriteString("### Git Commands for Reference\n\n")
	content.WriteString("```bash\n")
	content.WriteString(fmt.Sprintf("# View this commit\n"))
	content.WriteString(fmt.Sprintf("git show %s\n\n", commitHash))
	content.WriteString(fmt.Sprintf("# Compare with previous snapshot\n"))
	content.WriteString(fmt.Sprintf("git diff %s..%s\n\n", prevCommit, commitHash))
	content.WriteString(fmt.Sprintf("# View file changes\n"))
	content.WriteString(fmt.Sprintf("git diff --name-status %s..%s\n", prevCommit, commitHash))
	content.WriteString("```\n\n")

	// Footer
	content.WriteString("---\n")
	content.WriteString("*Report generated automatically by kalco*\n")

	return content.String(), nil
}

// ChangeSummary represents a summary of changes
type ChangeSummary struct {
	Namespaces        map[string]bool
	ResourceTypes     map[string]int
	ByNamespace       map[string]map[string][]string
	NewResources      int
	ModifiedResources int
	DeletedResources  int
}

// categorizeChanges organizes changed files into meaningful categories
func (r *ReportGenerator) categorizeChanges(changedFiles []string) *ChangeSummary {
	summary := &ChangeSummary{
		Namespaces:    make(map[string]bool),
		ResourceTypes: make(map[string]int),
		ByNamespace:   make(map[string]map[string][]string),
	}

	for _, file := range changedFiles {
		// Skip .git files
		if strings.Contains(file, ".git/") {
			continue
		}

		parts := strings.Split(file, string(os.PathSeparator))
		if len(parts) < 3 {
			continue
		}

		namespace := parts[0]
		resourceType := parts[1]
		filename := parts[2]

		// Track namespaces
		summary.Namespaces[namespace] = true

		// Track resource types
		summary.ResourceTypes[resourceType]++

		// Group by namespace and resource type
		if summary.ByNamespace[namespace] == nil {
			summary.ByNamespace[namespace] = make(map[string][]string)
		}
		summary.ByNamespace[namespace][resourceType] = append(summary.ByNamespace[namespace][resourceType], file)

		// Count new/modified/deleted
		if strings.HasSuffix(filename, ".yaml") {
			summary.ModifiedResources++
		}
	}

	return summary
}

// getFileStatus determines if a file was added, modified, or deleted
func (r *ReportGenerator) getFileStatus(file, prevCommit, currentCommit string) string {
	// Check if file exists in previous commit
	cmd := exec.Command("git", "show", fmt.Sprintf("%s:%s", prevCommit, file))
	cmd.Dir = r.repoPath
	if err := cmd.Run(); err != nil {
		return "🆕" // New file
	}

	// Check if file exists in current commit
	cmd = exec.Command("git", "show", fmt.Sprintf("%s:%s", currentCommit, file))
	cmd.Dir = r.repoPath
	if err := cmd.Run(); err != nil {
		return "🗑️" // Deleted file
	}

	return "✏️" // Modified file
}

// IsGitRepo checks if the directory is a Git repository
func (r *ReportGenerator) IsGitRepo() bool {
	gitDir := filepath.Join(r.repoPath, ".git")
	info, err := os.Stat(gitDir)
	return err == nil && info.IsDir()
}

// getCurrentCommitHash gets the current commit hash
func (r *ReportGenerator) getCurrentCommitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = r.repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// getPreviousCommitHash gets the previous commit hash
func (r *ReportGenerator) getPreviousCommitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD~1")
	cmd.Dir = r.repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// getChangedFiles gets the list of changed files between two commits
func (r *ReportGenerator) getChangedFiles(prevCommit, currentCommit string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", prevCommit, currentCommit)
	cmd.Dir = r.repoPath
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if len(output) == 0 {
		return []string{}, nil
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	var result []string
	for _, file := range files {
		if strings.TrimSpace(file) != "" {
			result = append(result, strings.TrimSpace(file))
		}
	}

	return result, nil
}
