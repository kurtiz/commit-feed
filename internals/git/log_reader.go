package git

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Commit struct {
	Hash    string
	Author  string
	Date    time.Time
	Message string
}

// Check if Git is installed
func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// Check if the current directory is a Git repository
func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "true"
}

// Fetch commits based on range or limit
func GetCommits(rangeArg string, limit int) ([]Commit, error) {
	if !IsGitInstalled() {
		return nil, errors.New("git is not installed on this system")
	}
	if !IsGitRepo() {
		return nil, errors.New("current directory is not a git repository")
	}

	args := []string{"log", "--pretty=format:%h|%an|%ad|%s", "--date=iso"}
	if rangeArg != "" {
		args = append(args, rangeArg)
	} else if limit > 0 {
		args = append(args, fmt.Sprintf("-n%d", limit))
	}

	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to read git logs: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	var commits []Commit

	for _, line := range lines {
		parts := strings.SplitN(line, "|", 4)
		if len(parts) < 4 {
			continue
		}
		date, _ := time.Parse("2006-01-02 15:04:05 -0700", strings.TrimSpace(parts[2]))
		commits = append(commits, Commit{
			Hash:    strings.TrimSpace(parts[0]),
			Author:  strings.TrimSpace(parts[1]),
			Date:    date,
			Message: strings.TrimSpace(parts[3]),
		})
	}

	return commits, nil
}

// GetProjectDescription reads README files to provide context about the project
func GetProjectDescription() (string, error) {
	possibleReadmes := []string{"README.md", "README.txt", "README", "readme.md", "readme.txt", "readme"}

	for _, readme := range possibleReadmes {
		if _, err := os.Stat(readme); err == nil {
			content, err := os.ReadFile(readme)
			if err != nil {
				continue
			}

			// Extract first few lines or first paragraph for context
			lines := strings.Split(string(content), "\n")
			var description strings.Builder

			for i, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				// Skip markdown headers and formatting
				if strings.HasPrefix(line, "#") {
					line = strings.TrimSpace(strings.TrimPrefix(line, "#"))
				}
				if strings.HasPrefix(line, ">") {
					line = strings.TrimSpace(strings.TrimPrefix(line, ">"))
				}

				description.WriteString(line)
				description.WriteString(" ")

				// Limit to first 3-4 meaningful lines to avoid too much context
				if i >= 3 && len(description.String()) > 200 {
					break
				}
			}

			result := strings.TrimSpace(description.String())
			if len(result) > 0 {
				return result, nil
			}
		}
	}

	return "", nil // No README found, return empty string
}
