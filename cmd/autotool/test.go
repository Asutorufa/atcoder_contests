package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func findLatestContestDir() string {
	entries, err := os.ReadDir("cmd")
	if err != nil {
		return ""
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if strings.HasPrefix(name, "20") && strings.Contains(name, "_") {
				dirs = append(dirs, name)
			}
		}
	}

	if len(dirs) == 0 {
		return ""
	}

	sort.Strings(dirs)
	return filepath.Join("cmd", dirs[len(dirs)-1])
}

func testProblem(problemLabel string, specificContestID string) error {
	fmt.Printf("Testing problem %s...\n", problemLabel)

	var targetDir string
	if specificContestID != "" {
		targetDir = findContestDirByID(specificContestID)
		if targetDir == "" {
			return fmt.Errorf("no contest directory found for %s", specificContestID)
		}
	} else {
		targetDir = findLatestContestDir()
		if targetDir == "" {
			return fmt.Errorf("no contest directory found (expected cmd/YYYYMMDD_<contest_id>)")
		}
	}

	fmt.Printf("Using contest directory: %s\n", targetDir)

	taskDir := filepath.Join(targetDir, problemLabel)
	goFile := filepath.Join(taskDir, problemLabel+".go")
	testDir := filepath.Join(taskDir, "test")

	if _, err := os.Stat(goFile); os.IsNotExist(err) {
		return fmt.Errorf("Go file not found: %s", goFile)
	}
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		return fmt.Errorf("Test directory not found: %s", testDir)
	}

	binFile := filepath.Join(taskDir, "exec_bin")
	if os.PathSeparator == '\\' {
		binFile += ".exe"
	}

	fmt.Printf("Compiling %s...\n", goFile)
	buildCmd := exec.Command("go", "build", "-o", binFile, goFile)
	buildCmd.Stderr = os.Stderr
	buildCmd.Stdout = os.Stdout
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}
	defer os.Remove(binFile)

	entries, err := os.ReadDir(testDir)
	if err != nil {
		return fmt.Errorf("failed to read test directory: %w", err)
	}

	passed := 0
	total := 0

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), "in_") {
			inFileName := entry.Name()
			id := strings.TrimPrefix(inFileName, "in_")
			outFileName := "out_" + id

			inFile := filepath.Join(testDir, inFileName)
			outFile := filepath.Join(testDir, outFileName)

			if _, err := os.Stat(outFile); os.IsNotExist(err) {
				continue
			}

			total++

			inContent, err := os.ReadFile(inFile)
			if err != nil {
				fmt.Printf("Test %s: Error reading input: %v\n", id, err)
				continue
			}

			expectedContent, err := os.ReadFile(outFile)
			if err != nil {
				fmt.Printf("Test %s: Error reading output: %v\n", id, err)
				continue
			}

			runCmd := exec.Command(binFile)
			runCmd.Stdin = bytes.NewReader(inContent)
			var stdout, stderr bytes.Buffer
			runCmd.Stdout = &stdout
			runCmd.Stderr = &stderr

			err = runCmd.Run()

			if err != nil {
				fmt.Printf("Test %s: Runtime error: %v\n", id, err)
				fmt.Printf("Stderr:\n%s\n", stderr.String())
				continue
			}

			actual := strings.TrimSpace(stdout.String())
			expected := strings.TrimSpace(string(expectedContent))

			if actual == expected {
				fmt.Printf("Test %s: PASSED\n", id)
				passed++
			} else {
				fmt.Printf("Test %s: FAILED\n", id)
				fmt.Printf("Expected:\n%s\n", expected)
				fmt.Printf("Actual:\n%s\n", actual)
			}
		}
	}

	fmt.Printf("\nSummary: %d / %d tests passed.\n", passed, total)
	if passed != total {
		return fmt.Errorf("not all tests passed")
	}
	return nil
}
