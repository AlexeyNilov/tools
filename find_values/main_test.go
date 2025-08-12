package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindFile(t *testing.T) {
	// Create a temporary root folder for the test
	rootDir := t.TempDir()

	// Create subdirectories
	subDir := filepath.Join(rootDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	// Create files
	targetFile := "target.txt"
	targetFilePath := filepath.Join(subDir, targetFile)
	if err := os.WriteFile(targetFilePath, []byte("test data"), 0644); err != nil {
		t.Fatalf("Failed to create target file: %v", err)
	}

	// Create another unrelated file
	if err := os.WriteFile(filepath.Join(rootDir, "other.txt"), []byte("other"), 0644); err != nil {
		t.Fatalf("Failed to create other file: %v", err)
	}

	// Call your function
	foundPath, err := FindFile(rootDir, targetFile)
	if err != nil {
		t.Fatalf("FindFile returned error: %v", err)
	}

	// Check the returned path
	if foundPath != targetFilePath {
		t.Errorf("Expected %q, got %q", targetFilePath, foundPath)
	}

	// Negative test: file does not exist
	_, err = FindFile(rootDir, "nonexistent.txt")
	if err == nil {
		t.Errorf("Expected error for nonexistent file, got nil")
	}
}
