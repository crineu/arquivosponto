package main

import (
	"os"
	"strings"
)

func ListArquivosPontoTools() []string {
	var tools []string

	files, err := os.ReadDir(ArquivosPontoFolder())
	if err != nil {
		RaiseErrorAndExit("folders.go", err)
	}

	// Iterate through the contents
	for _, file := range files {
		if file.IsDir() {
			if !strings.HasPrefix(file.Name(), "_") && !strings.HasPrefix(file.Name(), ".") {
				tools = append(tools, file.Name())
			}
		}
	}
	return tools
}
