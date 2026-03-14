package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type TestCase struct {
	Input  string `toml:"input,multiline"`
	Output string `toml:"output,multiline"`
}

type Samples struct {
	TestCases []TestCase `toml:"test"`
}

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

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  autotool init [flags] <contest_id>")
		fmt.Println("  autotool test [flags] <problem_label>")
		os.Exit(1)
	}

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	cookieFlag := initCmd.String("cookie", "", "Session cookie (e.g., REVEL_SESSION=...)")

	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
	contestFlag := testCmd.String("contest", "", "Specify a specific contest ID to test (e.g., abc300)")

	command := os.Args[1]

	switch command {
	case "init":
		initCmd.Parse(os.Args[2:])
		if initCmd.NArg() < 1 {
			fmt.Println("Usage: autotool init [flags] <contest_id>")
			initCmd.PrintDefaults()
			os.Exit(1)
		}

		if *cookieFlag != "" {
			cookieStr := *cookieFlag
			if after, ok := strings.CutPrefix(cookieStr, "REVEL_SESSION="); ok {
				cookieStr = after
			}
			os.Setenv("REVEL_SESSION", cookieStr)
		}

		contestID := initCmd.Arg(0)
		if err := initContest(contestID); err != nil {
			fmt.Printf("Init failed: %v\n", err)
			os.Exit(1)
		}
	case "test":
		testCmd.Parse(os.Args[2:])
		if testCmd.NArg() < 1 {
			fmt.Println("Usage: autotool test [flags] <problem_label>")
			testCmd.PrintDefaults()
			os.Exit(1)
		}
		problemLabel := testCmd.Arg(0)
		if err := testProblem(problemLabel, *contestFlag); err != nil {
			fmt.Printf("Test failed: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
