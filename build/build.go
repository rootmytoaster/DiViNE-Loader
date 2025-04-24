package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ANSI color codes for banner
const (
	purple = "\033[38;2;138;43;226m" // bright purple
	reset  = "\033[0m"               // reset code
)

// Apply purple color to the banner
func applyPurpleColor(line string) string {
	if strings.TrimSpace(line) == "" {
		return line
	}

	// Apply purple color to the entire line
	return purple + line + reset
}

// Raw banner without colors
var rawBanner = `
▓▓▒░░   ░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▒▒▒░░░░░░░   ░░░░▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▒▒░    ░░░░░░░▒▒▒▒▒  ▄▄▄▄  ▒▒▒▒   ▌ ░▒▒▒▒░░░░░░░░░   ▌ ░░░░░▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▒░░  █  ░░░░░░░░░▒  ▀   ▀█▌ ▒▒  ██  ░░░░░░░░░   ░  ██   ░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒
▒░░  ▓▓▓█  ░░░░░░░░░░     ▄█▌  ■▀▄▄▄■ ■▀█▓█   ▐▓▀ ■▀▄▄▄■ ■▀▄▄▄■▀▀▄     ▄▄▓▄▄▄  ░
▒░   ▒▒▓▓█   ░░░░░   ▄▄█▀██▒▌   ▐▓▒▌    ▐▓█▌  ██   ▐▒█▌   ▐▒▓▌    █   ░█▀    █
     ░░▒▒▓▓    ░░  █▒█▀DIV▄▒█   █▒█      █▒█ ▐█▌   █▓█    █▒█     █   ▓▒  ■▀▀
      ░░░▒▒▓      █▓▓▒████▓██▌ ▐░▓█▌      █▒▌██   ▐▓▒█▌  ▐▓▒█▌   ▐▓▌  ▐▓█▓▄  ▄■
       ▓░░░░▓      ▀▀▓▀▀▀▀▀▀ ▀■▀▀▀▀▀       ▀▀▀    ▀▓▀▀▀  ▀▀▓▀▀   ▀▀▀    ▀▀▓▀▀
       ▒░░░░░░▒▒					DiViNE Loader - Debug
	     ▒░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓■▄
`

// Apply color to the banner
func getPurpleBanner() string {
	lines := strings.Split(rawBanner, "\n")
	var coloredLines []string

	for _, line := range lines {
		coloredLines = append(coloredLines, applyPurpleColor(line))
	}

	return strings.Join(coloredLines, "\n")
}

func main() {
	// Display banner
	fmt.Println(getPurpleBanner())

	// Check if we're building for analysis mode
	analysisMode := false
	for _, arg := range os.Args {
		if arg == "--analysis" || arg == "-a" {
			analysisMode = true
			break
		}
	}

	if analysisMode {
		fmt.Println("[*Builder]: Building DiViNE Loader as console application for analysis...")

		// Build without the windowsgui flag for analysis mode
		cmd := exec.Command("go", "build", "main.go")
		cmd.Env = os.Environ()
		if runtime.GOOS != "windows" {
			cmd.Env = append(cmd.Env, "GOOS=windows")
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("[!] Build failed: %v\n", err)
			fmt.Println(string(output))
			os.Exit(1)
		}

		// Success messages
		fmt.Println("\n[>] BUILD COMPLETE: DiViNE Has Awoken in Analysis Mode")
		fmt.Println("[>] Run: main.exe --analysis to execute with analysis features")

		// Print a reminder about the two-step process
		fmt.Println("\nReminder: Analysis mode requires both:")
		fmt.Println("  1. Building without the windowsgui flag (which you've done)")
		fmt.Println("  2. Running with the --analysis flag")
	} else {
		fmt.Println("[*Builder] Building DiViNE Loader as Windows GUI application...")

		// Build with the windowsgui flag for normal mode
		cmd := exec.Command("go", "build", "-ldflags=-H windowsgui", "main.go")
		cmd.Env = os.Environ()
		if runtime.GOOS != "windows" {
			cmd.Env = append(cmd.Env, "GOOS=windows")
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("[!] Build failed: %v\n", err)
			fmt.Println(string(output))
			os.Exit(1)
		}

		// Success messages
		fmt.Println("\n[>] BUILD COMPLETE: DiViNE Has Awoken")
		fmt.Println("[>] Run: main.exe to execute normally")
		fmt.Println("[>] Run: main.exe --no-error to run without showing the error")

		// Print a note about analysis mode
		fmt.Println("\nNote: To build for analysis mode, use:")
		fmt.Println("  go run build/build.go --analysis")
	}
}
