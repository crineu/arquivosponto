package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runStowCommand(arg string) ([]string, error) {
	home, _ := os.UserHomeDir()
	cmd := exec.Command("stow", arg)
	cmd.Dir = home + "/arquivosponto/"

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // Capture standard error output as well

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var result []string
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if shouldIncludeLine(line) {
			result = append(result, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// shouldIncludeLine filters lines based on custom logic.
// Replace this with your actual filtering logic.
func shouldIncludeLine(line string) bool {
	// Example filtering: include lines that contain the word "example"
	return strings.Contains(line, "a")
}

func main() {
	output, err := runStowCommand("vim")
	if err != nil {
		fmt.Printf("Error running stow command: %v", err)
	}

	fmt.Println("Filtered output:")
	for _, line := range output {
		fmt.Println(line)
	}
}
