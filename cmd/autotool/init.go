package main

import (
	"atcoder/template"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var httpClient = &http.Client{}

var (
	reIn  = regexp.MustCompile(`(?is)<h3>Sample Input (\d+)</h3>.*?<pre>(.*?)</pre>`)
	reOut = regexp.MustCompile(`(?is)<h3>Sample Output (\d+)</h3>.*?<pre>(.*?)</pre>`)
)

func fetchWithAuth(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	sessionCookie := os.Getenv("REVEL_SESSION")
	if sessionCookie != "" {
		req.Header.Set("Cookie", "REVEL_SESSION="+sessionCookie)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code %d for %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type Task struct {
	Label string
	ID    string
}

func parseTasks(contestID string, html string) []Task {
	pattern := fmt.Sprintf(`<td class="text-center no-break"><a href="/contests/%s/tasks/([^"]+)">([^<]+)</a></td>`, contestID)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(html, -1)

	var tasks []Task
	for _, m := range matches {
		tasks = append(tasks, Task{
			ID:    m[1],
			Label: m[2],
		})
	}
	return tasks
}

type TestCase struct {
	Input  string
	Output string
}

func parseTestCases(html string) []TestCase {
	ins := reIn.FindAllStringSubmatch(html, -1)
	outs := reOut.FindAllStringSubmatch(html, -1)

	var tcs []TestCase

	outMap := make(map[string]string)
	for _, out := range outs {
		outMap[out[1]] = strings.TrimLeft(out[2], "\r\n")
	}

	seenInput := make(map[string]bool)
	for _, in := range ins {
		id := in[1]
		inStr := strings.TrimLeft(in[2], "\r\n")
		outStr := outMap[id]

		if !seenInput[inStr] {
			seenInput[inStr] = true
			tcs = append(tcs, TestCase{
				Input:  inStr,
				Output: outStr,
			})
		}
	}

	return tcs
}

func findContestDirByID(contestID string) string {
	entries, err := os.ReadDir("cmd")
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			if strings.HasSuffix(name, "_"+contestID) {
				return filepath.Join("cmd", name)
			}
		}
	}
	return ""
}

func initContest(contestID string) error {
	existingDir := findContestDirByID(contestID)
	if existingDir != "" {
		fmt.Printf("Contest directory %s already exists. Skipping initialization.\n", existingDir)
		return nil
	}

	fmt.Printf("Initializing contest %s...\n", contestID)

	tasksURL := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks", contestID)
	html, err := fetchWithAuth(tasksURL)
	if err != nil {
		return fmt.Errorf("failed to fetch tasks list: %w", err)
	}

	tasks := parseTasks(contestID, html)
	if len(tasks) == 0 {
		return fmt.Errorf("no tasks found for contest %s. Ensure the contest exists and check REVEL_SESSION if it's ongoing", contestID)
	}

	dateStr := time.Now().Format("20060102")
	dirName := fmt.Sprintf("%s_%s", dateStr, contestID)
	baseDir := filepath.Join("cmd", dirName)

	fmt.Printf("Creating directory %s...\n", baseDir)

	for _, task := range tasks {
		label := task.Label
		taskDir := filepath.Join(baseDir, label)

		if err := os.MkdirAll(taskDir, 0755); err != nil {
			fmt.Printf("Failed to create dir %s: %v\n", taskDir, err)
			continue
		}

		goFile := filepath.Join(taskDir, label+".go")
		if _, err := os.Stat(goFile); os.IsNotExist(err) {
			err = os.WriteFile(goFile, template.Template, 0644)
			if err != nil {
				fmt.Printf("Failed to write template to %s: %v\n", goFile, err)
			}
		}

		taskURL := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s", contestID, task.ID)
		taskHTML, err := fetchWithAuth(taskURL)
		if err != nil {
			fmt.Printf("Failed to fetch task %s: %v\n", task.ID, err)
			continue
		}

		tcs := parseTestCases(taskHTML)
		if len(tcs) > 0 {
			testDir := filepath.Join(taskDir, "test")
			os.MkdirAll(testDir, 0755)

			for i, tc := range tcs {
				inFile := filepath.Join(testDir, fmt.Sprintf("in_%d.txt", i+1))
				outFile := filepath.Join(testDir, fmt.Sprintf("out_%d.txt", i+1))

				if err := os.WriteFile(inFile, []byte(tc.Input), 0644); err != nil {
					fmt.Printf("Error writing file %s: %v\n", inFile, err)
				}
				if err := os.WriteFile(outFile, []byte(tc.Output), 0644); err != nil {
					fmt.Printf("Error writing file %s: %v\n", outFile, err)
				}
			}
			fmt.Printf("Downloaded %d test cases for %s (%s).\n", len(tcs), label, task.ID)
		} else {
			fmt.Printf("No test cases found for %s (%s).\n", label, task.ID)
		}
	}
	fmt.Println("Initialization complete.")
	return nil
}
