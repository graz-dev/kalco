package git

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	repo := NewGitRepo(tempDir)

	if repo == nil {
		t.Fatal("NewGitRepo returned nil")
	}

	if repo.path != tempDir {
		t.Errorf("expected path %s, got %s", tempDir, repo.path)
	}
}

func TestIsGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	repo := NewGitRepo(tempDir)

	// Should not be a git repo initially
	if repo.IsGitRepo() {
		t.Error("directory should not be a git repo initially")
	}

	// Create .git directory
	gitDir := filepath.Join(tempDir, ".git")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("failed to create .git directory: %v", err)
	}

	// Should now be a git repo
	if !repo.IsGitRepo() {
		t.Error("directory should be a git repo after creating .git")
	}
}

func TestCreateGitignore(t *testing.T) {
	tempDir := t.TempDir()
	repo := NewGitRepo(tempDir)

	if err := repo.createGitignore(); err != nil {
		t.Fatalf("failed to create .gitignore: %v", err)
	}

	gitignorePath := filepath.Join(tempDir, ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		t.Error(".gitignore file was not created")
	}
}
