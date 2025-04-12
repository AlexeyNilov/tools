package main

import (
	"strings"
	"testing"
)

func TestGetRepos(t *testing.T) {
	// Arrange: Define expected output
	expected := "repo_sync"

	// Act: Call the function we want to test
	repos := GetRepos("..")

	// Check that expected is part of one of the repo paths
	found := false
	for _, repo := range repos {
		// Check if expected is part of the absolute path of repo
		if strings.Contains(repo, expected) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected %s to be part of one of the repos, but it was not found", expected)
	}
}
