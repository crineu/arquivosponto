package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListArquivosPontoTools(t *testing.T) {
	// Create a temporary directory structure to simulate arquivosponto
	tmpDir := t.TempDir()

	// Create directories that should be included (with dot-* to be valid stow packages)
	includedDirs := []string{"zsh", "gitconfig", "neovim", "mise"}
	for _, dir := range includedDirs {
		dirPath := filepath.Join(tmpDir, dir)
		if err := os.Mkdir(dirPath, 0755); err != nil {
			t.Fatalf("failed to create dir %s: %v", dir, err)
		}
		// Create a dot-* file to mark as valid stow package
		if err := os.WriteFile(filepath.Join(dirPath, "dot-zshrc"), []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create dotfile in %s: %v", dir, err)
		}
	}

	// Create directories that should be excluded (prefix _ or .)
	excludedDirs := []string{"_tui", ".git", ".hidden"}
	for _, dir := range excludedDirs {
		if err := os.Mkdir(filepath.Join(tmpDir, dir), 0755); err != nil {
			t.Fatalf("failed to create dir %s: %v", dir, err)
		}
	}

	// Create a file (should be ignored)
	if err := os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("test"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	// Override the folder path via environment variable
	t.Setenv("ARQUIVOSPONTO_FOLDER", tmpDir+"/")

	tools := ListArquivosPontoTools()

	// Verify included directories are present
	if len(tools) != len(includedDirs) {
		t.Errorf("expected %d tools, got %d: %v", len(includedDirs), len(tools), tools)
	}

	for _, expected := range includedDirs {
		found := false
		for _, tool := range tools {
			if tool == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected tool %s not found in result: %v", expected, tools)
		}
	}

	// Verify excluded directories are not present
	for _, excluded := range excludedDirs {
		for _, tool := range tools {
			if tool == excluded {
				t.Errorf("excluded directory %s should not be in result: %v", excluded, tools)
			}
		}
	}
}

func TestListArquivosPontoTools_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("ARQUIVOSPONTO_FOLDER", tmpDir+"/")

	tools := ListArquivosPontoTools()

	if len(tools) != 0 {
		t.Errorf("expected empty result for empty directory, got: %v", tools)
	}
}
