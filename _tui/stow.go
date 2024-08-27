package main

import (
	"bufio"
	"bytes"
	"os/exec"
)

// man stow:
//
// -n
// --no
//		Do not perform any operations; merely show what would happen.
// -S
// --stow
//		Stow the packages. This is the default action and so can be omitted
// -D
// --delete
//		Unstow the packages rather than installing them.
// -R
// --restow
//		Restow packages (first unstow, then stow again).
// 		This is useful for pruning obsolete symlinks from the target tree after updating the software in a package.

func StowRemDry(folder string) []string {
	return runStowCommand("-Dn", folder)
}
func StowRem(folder string) []string {
	return runStowCommand("-D ", folder)
}
func StowAddDry(folder string) []string {
	return runStowCommand("-Sn", folder)
}
func StowAdd(folder string) []string {
	return runStowCommand("S", folder)
}
func StowRestowDry(folder string) []string {
	return runStowCommand("-Rn", folder)
}
func StowRestowAdd(folder string) []string {
	return runStowCommand("-R", folder)
}

func runStowCommand(arg string, folder string) []string {
	cmd := exec.Command("stow", arg, folder)
	cmd.Dir = ArquivosPontoFolder()

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // Capture standard error output as well

	err := cmd.Run()
	if err != nil {
		RaiseErrorAndExit("stow.go:cmd.Run", err)
	}

	var result []string
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		// if shouldIncludeLine(line) {
		result = append(result, line)
		// }
	}
	if err := scanner.Err(); err != nil {
		RaiseErrorAndExit("stow.go:scanner.Err", err)
	}

	return result
}

// shouldIncludeLine filters lines based on custom logic.
// func shouldIncludeLine(line string) bool {
// 	return strings.Contains(line, "a")
// }
