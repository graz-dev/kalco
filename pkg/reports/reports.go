package reports

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"kalco/pkg/validation"
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

	// Generate validation report
	validationContent, err := r.generateValidationReport()
	if err != nil {
		fmt.Printf("  ⚠️  Warning: Validation report generation failed: %v\n", err)
	} else {
		// Insert validation content before the footer
		content = strings.Replace(content, "\n\n---\n*Report generated automatically by kalco*", "\n\n"+validationContent+"\n\n---\n*Report generated automatically by kalco*", 1)
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
		commitMessage = "Cluster snapshot " + time.Now().Format("2006-01-02-15-04-05")
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
	content.WriteString("**Generated**: " + time.Now().Format("2006-01-02 15:04:05 UTC") + "\n")
	content.WriteString("**Commit Message**: " + commitMessage + "\n\n")

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
		content.WriteString("Failed to retrieve Git information: " + err.Error() + "\n\n")
		return content.String(), nil
	}

	content.WriteString("**Commit Hash**: `" + commitHash + "`\n\n")

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
	content.WriteString("**Previous Commit**: `" + prevCommit + "`\n\n")

	// Get changed files
	changedFiles, err := r.getChangedFiles(prevCommit, commitHash)
	if err != nil {
		content.WriteString("### Error Getting Changes\n\n")
		content.WriteString("Failed to retrieve changes: " + err.Error() + "\n\n")
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
	content.WriteString("- **Total Files Changed**: " + strconv.Itoa(len(changedFiles)) + "\n")
	content.WriteString("- **Namespaces Affected**: " + strconv.Itoa(len(changes.Namespaces)) + "\n")
	content.WriteString("- **Resource Types Changed**: " + strconv.Itoa(len(changes.ResourceTypes)) + "\n")
	content.WriteString("- **New Resources**: " + strconv.Itoa(changes.NewResources) + "\n")
	content.WriteString("- **Modified Resources**: " + strconv.Itoa(changes.ModifiedResources) + "\n")
	content.WriteString("- **Deleted Resources**: " + strconv.Itoa(changes.DeletedResources) + "\n")
	content.WriteString("\n")

	// Write detailed changes with diff information
	content.WriteString("### Detailed Changes\n\n")

	// Group by namespace
	for namespace, resources := range changes.ByNamespace {
		if namespace == "_cluster" {
			content.WriteString("#### 🌐 Cluster-Scoped Resources\n\n")
		} else {
			content.WriteString("#### 📁 Namespace: `" + namespace + "`\n\n")
		}

		for resourceType, files := range resources {
			content.WriteString("**" + resourceType + "**:\n")
			for _, file := range files {
				status := r.getFileStatus(file, prevCommit, commitHash)
				content.WriteString("- " + status + " `" + filepath.Base(file) + "`\n")
			}
			content.WriteString("\n")
		}
	}

	// Write detailed diff information for each changed file
	content.WriteString("## 🔍 Detailed Resource Changes\n\n")
	content.WriteString("This section shows the specific changes made to each resource:\n\n")

	for namespace, resources := range changes.ByNamespace {
		if namespace == "_cluster" {
			content.WriteString("### 🌐 Cluster-Scoped Resources\n\n")
		} else {
			content.WriteString("### 📁 Namespace: `" + namespace + "`\n\n")
		}

		for resourceType, files := range resources {
			content.WriteString("#### " + resourceType + "\n\n")

			for _, file := range files {
				status := r.getFileStatus(file, prevCommit, commitHash)
				filename := filepath.Base(file)
				resourceName := strings.TrimSuffix(filename, ".yaml")

				content.WriteString("**" + status + "** `" + resourceName + "` (" + filename + ")\n\n")

				// Get detailed diff information
				diffInfo, err := r.getDetailedDiff(file, prevCommit, commitHash, status)
				if err != nil {
					content.WriteString("⚠️ Error getting diff: " + err.Error() + "\n\n")
					continue
				}

				content.WriteString(diffInfo)
				content.WriteString("\n---\n\n")
			}
		}
	}

	// Write resource type summary
	content.WriteString("## 📊 Resource Type Summary\n\n")
	for resourceType, count := range changes.ResourceTypes {
		content.WriteString("- **" + resourceType + "**: " + strconv.Itoa(count) + " changes\n")
	}
	content.WriteString("\n")

	// Write Git commands for reference
	content.WriteString("## 💻 Git Commands for Reference\n\n")
	content.WriteString("```bash\n")
	content.WriteString("# View this commit\n")
	content.WriteString("git show " + commitHash + "\n\n")
	content.WriteString("# Compare with previous snapshot\n")
	content.WriteString("git diff " + prevCommit + ".." + commitHash + "\n\n")
	content.WriteString("# View file changes\n")
	content.WriteString("git diff --name-status " + prevCommit + ".." + commitHash + "\n\n")
	content.WriteString("# View specific file diff\n")
	content.WriteString("git diff " + prevCommit + ".." + commitHash + " -- <filename>\n")
	content.WriteString("```\n\n")

	// Footer
	content.WriteString("---\n")
	content.WriteString("*Report generated automatically by kalco*\n")

	return content.String(), nil
}

// getDetailedDiff gets detailed diff information for a specific file
func (r *ReportGenerator) getDetailedDiff(file, prevCommit, currentCommit, status string) (string, error) {
	var content strings.Builder

	switch status {
	case "🆕":
		// New file - show the complete content
		content.WriteString("**New Resource Created**\n\n")
		content.WriteString("```yaml\n")
		currentContent, err := r.getFileContent(file, currentCommit)
		if err != nil {
			return "⚠️ Error reading file content: " + err.Error(), nil
		}
		content.WriteString(currentContent)
		content.WriteString("\n```\n\n")

		// Add metadata summary
		content.WriteString("**Resource Details**:\n")
		content.WriteString("- Type: New resource\n")
		content.WriteString("- Status: Created in this snapshot\n")
		content.WriteString("- File: `" + file + "`\n\n")

	case "🗑️":
		// Deleted file - show what was removed
		content.WriteString("**Resource Deleted**\n\n")
		content.WriteString("```yaml\n")
		previousContent, err := r.getFileContent(file, prevCommit)
		if err != nil {
			return "⚠️ Error reading previous file content: " + err.Error(), nil
		}
		content.WriteString(previousContent)
		content.WriteString("\n```\n\n")

		// Add metadata summary
		content.WriteString("**Resource Details**:\n")
		content.WriteString("- Type: Deleted resource\n")
		content.WriteString("- Status: Removed in this snapshot\n")
		content.WriteString("- File: `" + file + "`\n\n")

	case "✏️":
		// Modified file - show the diff
		content.WriteString("**Resource Modified**\n\n")

		// Get the actual diff output
		diffOutput, err := r.getGitDiff(file, prevCommit, currentCommit)
		if err != nil {
			content.WriteString("⚠️ Error getting diff: " + err.Error() + "\n\n")
			// Fallback to showing before/after
			content.WriteString("**Before (Previous Snapshot):**\n")
			content.WriteString("```yaml\n")
			previousContent, err := r.getFileContent(file, prevCommit)
			if err != nil {
				content.WriteString("Error reading previous content: " + err.Error())
			} else {
				content.WriteString(previousContent)
			}
			content.WriteString("\n```\n\n")

			content.WriteString("**After (Current Snapshot):**\n")
			content.WriteString("```yaml\n")
			currentContent, err := r.getFileContent(file, currentCommit)
			if err != nil {
				content.WriteString("Error reading current content: " + err.Error())
			} else {
				content.WriteString(currentContent)
			}
			content.WriteString("\n```\n\n")
		} else {
			content.WriteString("**Changes Detected:**\n")
			content.WriteString("```diff\n")
			content.WriteString(diffOutput)
			content.WriteString("\n```\n\n")
		}

		// Add metadata summary
		content.WriteString("**Resource Details**:\n")
		content.WriteString("- Type: Modified resource\n")
		content.WriteString("- Status: Updated in this snapshot\n")
		content.WriteString("- File: `" + file + "`\n\n")

		// Add change summary
		changeSummary, err := r.getChangeSummary(file, prevCommit, currentCommit)
		if err == nil && changeSummary != "" {
			content.WriteString("**Change Summary:**\n")
			content.WriteString(changeSummary)
			content.WriteString("\n")
		}
	}

	return content.String(), nil
}

// getFileContent gets the content of a file at a specific commit
func (r *ReportGenerator) getFileContent(file, commit string) (string, error) {
	cmd := exec.Command("git", "show", commit+":"+file)
	cmd.Dir = r.repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get file content: %w", err)
	}
	return string(output), nil
}

// getGitDiff gets the git diff output for a file between two commits
func (r *ReportGenerator) getGitDiff(file, prevCommit, currentCommit string) (string, error) {
	cmd := exec.Command("git", "diff", prevCommit, currentCommit, "--", file)
	cmd.Dir = r.repoPath
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git diff: %w", err)
	}
	return string(output), nil
}

// getChangeSummary provides a human-readable summary of what changed
func (r *ReportGenerator) getChangeSummary(file, prevCommit, currentCommit string) (string, error) {
	// Get the diff and analyze it for common patterns
	diffOutput, err := r.getGitDiff(file, prevCommit, currentCommit)
	if err != nil {
		return "", err
	}

	var summary strings.Builder
	lines := strings.Split(diffOutput, "\n")

	addedLines := 0
	removedLines := 0
	changedSections := make(map[string]bool)

	for _, line := range lines {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			addedLines++
			// Try to identify what section changed
			if strings.Contains(line, ":") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) > 0 {
					section := strings.TrimSpace(parts[0])
					if section != "" {
						changedSections[section] = true
					}
				}
			}
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			removedLines++
		}
	}

	if addedLines > 0 || removedLines > 0 {
		summary.WriteString("- **Lines Added**: " + strconv.Itoa(addedLines) + "\n")
		summary.WriteString("- **Lines Removed**: " + strconv.Itoa(removedLines) + "\n")
	}

	if len(changedSections) > 0 {
		summary.WriteString("- **Sections Modified**:\n")
		for section := range changedSections {
			summary.WriteString("  - `" + section + "`\n")
		}
	}

	return summary.String(), nil
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
	cmd := exec.Command("git", "show", prevCommit+":"+file)
	cmd.Dir = r.repoPath
	if err := cmd.Run(); err != nil {
		return "🆕" // New file
	}

	// Check if file exists in current commit
	cmd = exec.Command("git", "show", currentCommit+":"+file)
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

// generateValidationReport generates a cross-reference validation report
func (r *ReportGenerator) generateValidationReport() (string, error) {
	var content strings.Builder

	// Run cross-reference validation
	validator := validation.NewResourceValidator(r.outputDir)
	result, err := validator.Validate()
	if err != nil {
		return "", fmt.Errorf("failed to run validation: %w", err)
	}

	content.WriteString("## 🔍 Cross-Reference Validation Report\n\n")
	content.WriteString("This section analyzes exported resources for broken references that could cause issues when reapplying.\n\n")
	content.WriteString("> **⚠️  Important**: Broken references will cause errors when you try to reapply resources to a cluster.\n\n")

	// Validation Summary
	content.WriteString("### 📊 Validation Summary\n\n")
	content.WriteString(fmt.Sprintf("- **Total References**: %d\n", result.Summary.TotalReferences))
	content.WriteString(fmt.Sprintf("- **✅ Valid References**: %d\n", result.Summary.ValidReferences))
	content.WriteString(fmt.Sprintf("- **❌ Broken References**: %d\n", result.Summary.BrokenReferences))
	content.WriteString(fmt.Sprintf("- **⚠️  Warning References**: %d\n", result.Summary.WarningReferences))
	content.WriteString("\n")

	// Broken References (most important)
	if len(result.BrokenReferences) > 0 {
		content.WriteString("### ❌ Broken References - ACTION REQUIRED\n\n")
		content.WriteString("**🚨 CRITICAL**: These references will cause errors when reapplying resources to a cluster!\n\n")

		// Group by source type
		grouped := make(map[string][]validation.ResourceReference)
		for _, ref := range result.BrokenReferences {
			grouped[ref.SourceType] = append(grouped[ref.SourceType], ref)
		}

		for sourceType, refs := range grouped {
			content.WriteString(fmt.Sprintf("#### %s\n\n", sourceType))
			for _, ref := range refs {
				content.WriteString(fmt.Sprintf("**❌ BROKEN**: %s `%s` references %s `%s`\n\n",
					ref.SourceType, ref.SourceName, ref.TargetType, ref.TargetName))
				content.WriteString(fmt.Sprintf("- **Problem**: The %s `%s` does not exist\n", ref.TargetType, ref.TargetName))
				content.WriteString(fmt.Sprintf("- **Location**: `%s` in %s `%s`\n", ref.Field, ref.SourceType, ref.SourceName))
				content.WriteString(fmt.Sprintf("- **Namespace**: `%s`\n", ref.SourceNamespace))
				if ref.TargetNamespace != ref.SourceNamespace {
					content.WriteString(fmt.Sprintf("- **Target Namespace**: `%s`\n", ref.TargetNamespace))
				}
				content.WriteString(fmt.Sprintf("- **Impact**: This resource will fail to apply\n\n"))
			}
		}
	} else {
		content.WriteString("### ✅ No Broken References Found\n\n")
		content.WriteString("🎉 **Excellent!** All cross-references in your cluster are valid and safe to reapply.\n\n")
	}

	// Warning References
	if len(result.WarningReferences) > 0 {
		content.WriteString("### ⚠️  Warning References - Manual Verification Needed\n\n")
		content.WriteString("**These references point to external resources that kalco cannot validate:**\n\n")

		// Group by source type
		grouped := make(map[string][]validation.ResourceReference)
		for _, ref := range result.WarningReferences {
			grouped[ref.SourceType] = append(grouped[ref.SourceType], ref)
		}

		for sourceType, refs := range grouped {
			content.WriteString(fmt.Sprintf("#### %s\n\n", sourceType))
			for _, ref := range refs {
				content.WriteString(fmt.Sprintf("**⚠️  EXTERNAL**: %s `%s` references %s `%s`\n\n",
					ref.SourceType, ref.SourceName, ref.TargetType, ref.TargetName))
				content.WriteString(fmt.Sprintf("- **Type**: External %s reference\n", ref.TargetType))
				content.WriteString(fmt.Sprintf("- **Location**: `%s` in %s `%s`\n", ref.Field, ref.SourceType, ref.SourceName))
				content.WriteString(fmt.Sprintf("- **Namespace**: `%s`\n", ref.SourceNamespace))
				content.WriteString(fmt.Sprintf("- **Action**: Verify this %s exists in your authentication system\n", ref.TargetType))
				content.WriteString(fmt.Sprintf("- **Note**: This is normal for system resources and external users/groups\n\n"))
			}
		}
	}

	// Valid References (summary only)
	if len(result.ValidReferences) > 0 {
		content.WriteString("### ✅ Valid References - All Good!\n\n")
		content.WriteString(fmt.Sprintf("**🎉 %d references are properly configured and safe to reapply:**\n\n", len(result.ValidReferences)))

		// Group by source type
		grouped := make(map[string]int)
		for _, ref := range result.ValidReferences {
			grouped[ref.SourceType]++
		}

		for sourceType, count := range grouped {
			content.WriteString(fmt.Sprintf("- **%s**: %d valid references ✅\n", sourceType, count))
		}
		content.WriteString("\n")
	}

	// Recommendations
	content.WriteString("### 💡 Action Plan\n\n")
	if len(result.BrokenReferences) > 0 {
		content.WriteString("**🚨 IMMEDIATE ACTIONS REQUIRED:**\n\n")
		content.WriteString("1. **🔧 Fix Broken References**: Resolve all broken references before reapplying\n")
		content.WriteString("2. **✅ Verify Target Resources**: Ensure all referenced resources exist in the target cluster\n")
		content.WriteString("3. **🌐 Check Namespaces**: Verify cross-namespace references are correct\n")
		content.WriteString("4. **🧪 Test in Staging**: Apply resources to a test environment first\n")
		content.WriteString("5. **📋 Review Warnings**: Check warning references for external resources\n\n")
	} else {
		content.WriteString("**🎉 Your cluster configuration looks excellent!**\n\n")
		content.WriteString("1. **✅ Safe to Reapply**: All references are valid and will work correctly\n")
		content.WriteString("2. **⚠️  Monitor Warnings**: Check warning references if any exist (these are usually normal)\n")
		content.WriteString("3. **🔄 Regular Validation**: Run this validation after cluster changes\n\n")
	}

	content.WriteString("**📝 Note**: This validation only checks for missing resources. It does not validate resource configurations, permissions, or runtime behavior.\n\n")

	return content.String(), nil
}
