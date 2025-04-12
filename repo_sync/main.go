package main

import (
	"os"
	"path/filepath"
    "fmt"
    "os/exec"
    "sync"
    "io"
)

// GetRepos returns a list of absolute paths to subfolders in the given path
func GetRepos(path string) []string {
	var repos []string

	// Read the directory entries
	entries, err := os.ReadDir(path)
	if err != nil {
		return repos // Return empty slice if there's an error
	}

	// Filter for directories and append their absolute paths to the list
	for _, entry := range entries {
		if entry.IsDir() {
			absPath, err := filepath.Abs(filepath.Join(path, entry.Name()))
			if err == nil {
				repos = append(repos, absPath)
			}
		}
	}

	return repos
}

// executeGitPull performs a `git pull` operation in the specified repository path
func executeGitPull(repo string) error {
	// Change directory to the repository path
	if err := os.Chdir(repo); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", repo, err)
	}

	// Execute `git pull`
	cmd := exec.Command("git", "pull")
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull command failed: %w", err)
	} else {
        fmt.Printf("Synced %s\n", repo)
    }

	return nil
}

// SyncRepos iterates over the list of repository paths and attempts to sync them in parallel
func SyncRepos(repos []string) {
	var wg sync.WaitGroup

	for _, repo := range repos {
		wg.Add(1)

		go func(repo string) {
			defer wg.Done()

			// Check if the .git folder exists
			gitDir := filepath.Join(repo, ".git")
			if _, err := os.Stat(gitDir); os.IsNotExist(err) {
				return
			}

			// Execute git pull
			if err := executeGitPull(repo); err != nil {
				fmt.Printf("Error syncing %s: %v\n", repo, err)
			}
		}(repo)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func main() {
    path := os.Args[1]
    repos := GetRepos(path)
    SyncRepos(repos)
}
