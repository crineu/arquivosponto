package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateStatus(t *testing.T) {
	// Create temporary stow-like package
	tmpDir := t.TempDir()
	pkgDir := filepath.Join(tmpDir, "testpkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a dotfile
	if err := os.WriteFile(filepath.Join(pkgDir, "dot-testrc"), []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)
	t.Setenv("ARQUIVOSPONTO_FOLDER", tmpDir+"/")

	// No symlink yet → NotInstalled
	if status := CalculateStatus("testpkg"); status != NotInstalled {
		t.Errorf("expected NotInstalled, got %v", status)
	}

	// Create correct symlink
	target := filepath.Join(homeDir, ".testrc")
	if err := os.Symlink(filepath.Join(pkgDir, "dot-testrc"), target); err != nil {
		t.Fatal(err)
	}

	// Now should be Installed
	if status := CalculateStatus("testpkg"); status != Installed {
		t.Errorf("expected Installed, got %v", status)
	}
}

func TestCalculateStatus_NoDotFiles(t *testing.T) {
	tmpDir := t.TempDir()
	pkgDir := filepath.Join(tmpDir, "scripts")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "script.sh"), []byte("#!/bin/sh"), 0755); err != nil {
		t.Fatal(err)
	}

	t.Setenv("ARQUIVOSPONTO_FOLDER", tmpDir+"/")

	status := CalculateStatus("scripts")
	if status != NotInstalled {
		t.Errorf("scripts without symlinks should be NotInstalled, got %v", status)
	}
}

