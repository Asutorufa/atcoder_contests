package main

import (
	"fmt"
	"os"
	"strings"
)

func loadEnv() {
	content, err := os.ReadFile(".env")
	if err == nil {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			}
		}
	}
}

func main() {
	loadEnv()

	// Parse flags and remove them from os.Args
	var filteredArgs []string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--cookie=") {
			cookieStr := strings.TrimPrefix(arg, "--cookie=")
			if strings.HasPrefix(cookieStr, "REVEL_SESSION=") {
				cookieStr = strings.TrimPrefix(cookieStr, "REVEL_SESSION=")
			}
			os.Setenv("REVEL_SESSION", cookieStr)
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	os.Args = filteredArgs

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  autotool init <contest_id> [--cookie=\"REVEL_SESSION=...\"]")
		fmt.Println("  autotool test <problem_label>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Usage: autotool init <contest_id> [--cookie=\"REVEL_SESSION=...\"]")
			os.Exit(1)
		}
		contestID := os.Args[2]
		initContest(contestID)
	case "test":
		if len(os.Args) < 3 {
			fmt.Println("Usage: autotool test <problem_label>")
			os.Exit(1)
		}
		problemLabel := os.Args[2]
		testProblem(problemLabel)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
