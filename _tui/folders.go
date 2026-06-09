package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListArquivosPontoTools() []string {
	var tools []string

	repoPath := ArquivosPontoFolder()
	files, err := os.ReadDir(repoPath)
	if err != nil {
		// Graceful fallback: return empty list instead of crashing
		fmt.Fprintf(os.Stderr, "Warning: could not read directory %s: %v\n", repoPath, err)
		return tools
	}

	for _, file := range files {
		if file.IsDir() {
			name := file.Name()
			if strings.HasPrefix(name, "_") || strings.HasPrefix(name, ".") {
				continue
			}
			// Validate it's a stow package: contains dot-* files/dirs or is non-empty
			if isValidStowPackage(name) {
				tools = append(tools, name)
			}
		}
	}
	return tools
}

// isValidStowPackage checks if directory contains stow-managed files (dot-* prefix)
func isValidStowPackage(dirName string) bool {
	fullPath := filepath.Join(ArquivosPontoFolder(), dirName)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return false
	}

	if len(entries) == 0 {
		return false // empty dir is not a valid stow package
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "dot-") {
			return true
		}
	}
	// Also accept if directory has any content (common case for stow packages)
	return true
}
