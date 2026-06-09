package main

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var stowAvailable bool

func init() {
	if _, err := exec.LookPath("stow"); err != nil {
		stowAvailable = false
	} else {
		stowAvailable = true
	}
}

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
	return runStowCommand("-D", folder)
}
func StowAddDry(folder string) []string {
	return runStowCommand("-Sn", folder)
}
func StowAdd(folder string) []string {
	return runStowCommand("-S", folder)
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
		// Exit code 2 from stow dry-run means "nothing to do" - not a fatal error
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 2 {
			// Return whatever output we captured (may be empty or contain info)
		} else {
			RaiseErrorAndExit("stow.go:cmd.Run", err)
		}
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

// CalculateStatus inspects symlinks in $HOME to determine stow package status.
// It does not spawn stow processes — it directly checks if dot-* entries
// have corresponding symlinks pointing to the correct repo paths.
func CalculateStatus(pkg string) Status {
	repoRoot := ArquivosPontoFolder()
	pkgDir := filepath.Join(repoRoot, pkg)

	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return NotInstalled
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return NotInstalled
	}

	totalDotEntries := 0
	correctLinks := 0

	for _, entry := range entries {
		name := entry.Name()
		// Support both dot-* convention and direct dotfiles/dirs (e.g. .config/)
		isDotEntry := strings.HasPrefix(name, "dot-") || strings.HasPrefix(name, ".")
		if !isDotEntry || name == "." || name == ".." {
			continue
		}
		totalDotEntries++

		// stow convention: dot-foo -> ~/.foo, or .config -> ~/.config
		var targetName string
		if strings.HasPrefix(name, "dot-") {
			targetName = "." + strings.TrimPrefix(name, "dot-")
		} else {
			targetName = name // already has dot prefix like .config
		}
		targetPath := filepath.Join(homeDir, targetName)
		sourcePath := filepath.Join(pkgDir, name)

		// Check if targetPath is a symlink pointing to sourcePath
		info, err := os.Lstat(targetPath)
		if err != nil {
			continue // does not exist
		}
		if info.Mode()&os.ModeSymlink == 0 {
			// Might be a directory that stow populates (not a direct symlink)
			// For nested stow (e.g. .config/zellij), check inside
			if info.IsDir() {
				// Recursively check if any symlink inside points to source
				if hasAnySymlinkToSource(targetPath, sourcePath) {
					correctLinks++
				}
			}
			continue
		}

		linkDest, err := os.Readlink(targetPath)
		if err != nil {
			continue
		}

		if !filepath.IsAbs(linkDest) {
			linkDest = filepath.Join(homeDir, linkDest)
		}

		linkDestAbs, _ := filepath.Abs(linkDest)
		sourceAbs, _ := filepath.Abs(sourcePath)

		if linkDestAbs == sourceAbs {
			correctLinks++
		}
	}

	if totalDotEntries > 0 {
		if correctLinks == 0 {
			return NotInstalled
		}
		if correctLinks == totalDotEntries {
			return Installed
		}
		return Outdated
	}

	// No dot-* entries: for non-dotfile packages, check if stow was used
	// by looking for any symlink in common locations that points to this package.
	// This is a heuristic - stow links files from the package root to target.
	commonTargets := []string{
		filepath.Join(homeDir, "bin"),
		filepath.Join(homeDir, ".local", "bin"),
		filepath.Join(homeDir, "scripts"), // if user stows to ~/scripts
	}

	for _, targetDir := range commonTargets {
		// Read targetDir entries and check if any symlink points into pkgDir
		tEntries, err := os.ReadDir(targetDir)
		if err != nil {
			continue
		}
		for _, tEntry := range tEntries {
			tPath := filepath.Join(targetDir, tEntry.Name())
			info, err := os.Lstat(tPath)
			if err != nil || info.Mode()&os.ModeSymlink == 0 {
				continue
			}
			linkDest, err := os.Readlink(tPath)
			if err != nil {
				continue
			}
			if !filepath.IsAbs(linkDest) {
				linkDest = filepath.Join(targetDir, linkDest)
			}
			linkDestAbs, _ := filepath.Abs(linkDest)
			if strings.HasPrefix(linkDestAbs, pkgDir) {
				return Installed
			}
		}
	}

	return NotInstalled
}

// hasAnySymlinkToSource recursively checks if any symlink inside targetDir
// points to a location inside sourceDir (for nested stow packages like .config/zellij)
func hasAnySymlinkToSource(targetDir, sourceDir string) bool {
	found := false
	_ = filepath.WalkDir(targetDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || found {
			return nil
		}
		if d.Type()&os.ModeSymlink != 0 {
			linkDest, err := os.Readlink(path)
			if err != nil {
				return nil
			}
			if !filepath.IsAbs(linkDest) {
				linkDest = filepath.Join(filepath.Dir(path), linkDest)
			}
			linkDestAbs, _ := filepath.Abs(linkDest)
			sourceAbs, _ := filepath.Abs(sourceDir)
			if strings.HasPrefix(linkDestAbs, sourceAbs) {
				found = true
			}
		}
		return nil
	})
	return found
}
