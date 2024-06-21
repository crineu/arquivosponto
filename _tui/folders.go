package main

import (
	"os"
	"strings"
)

func ListArquivosPontoTools() ([]string, error) {
	var tools []string

	home, err := os.UserHomeDir()
	files, err := os.ReadDir(home + "/arquivosponto/.")
	if err != nil {
		return nil, err
	}

	// Iterate through the contents
	for _, file := range files {
		if file.IsDir() {
			if !strings.HasPrefix(file.Name(), "_") && !strings.HasPrefix(file.Name(), ".") {
				tools = append(tools, file.Name())
			}
		}
	}
	return tools, nil
}
