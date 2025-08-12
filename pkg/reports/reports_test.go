package reports

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewReportGenerator(t *testing.T) {
	tempDir := t.TempDir()
	gen := NewReportGenerator(tempDir)

	if gen == nil {
		t.Fatal("NewReportGenerator returned nil")
	}

	if gen.outputDir != tempDir {
		t.Errorf("expected outputDir %s, got %s", tempDir, gen.outputDir)
	}
}

func TestGenerateFilename(t *testing.T) {
	tempDir := t.TempDir()
	gen := NewReportGenerator(tempDir)

	// Test with custom message
	filename := gen.generateFilename("Test commit message")
	expected := "Test-commit-message.md"
	if filename != expected {
		t.Errorf("expected filename %s, got %s", expected, filename)
	}

	// Test with empty message (should use timestamp)
	filename = gen.generateFilename("")
	if !strings.Contains(filename, "Cluster-snapshot") {
		t.Error("empty message should generate timestamp-based filename")
	}
	if !strings.HasSuffix(filename, ".md") {
		t.Error("filename should end with .md extension")
	}
}

func TestGenerateReport(t *testing.T) {
	tempDir := t.TempDir()
	gen := NewReportGenerator(tempDir)

	// Test report generation
	err := gen.GenerateReport("Test report")
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Check if reports directory was created
	reportsDir := filepath.Join(tempDir, "kalco-reports")
	if _, err := os.Stat(reportsDir); os.IsNotExist(err) {
		t.Error("reports directory was not created")
	}

	// Check if report file was created
	reportFile := filepath.Join(reportsDir, "Test-report.md")
	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("report file was not created")
	}
}

func TestIsGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	gen := NewReportGenerator(tempDir)

	// Should not be a git repo initially
	if gen.IsGitRepo() {
		t.Error("directory should not be a git repo initially")
	}

	// Create .git directory
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("failed to create .git directory: %v", err)
	}

	// Should now be a git repo
	if !gen.IsGitRepo() {
		t.Error("directory should be a git repo after creating .git")
	}
}
