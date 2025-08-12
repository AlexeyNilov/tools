package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// FindFile searches for a file named inputFileName starting from rootFolder.
// It returns the full path if found, or an error if not found.
func FindFile(rootFolder, inputFileName string) (string, error) {
	var foundPath string
	err := filepath.WalkDir(rootFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Stop walking if we can’t access something
			return err
		}
		if !d.IsDir() && d.Name() == inputFileName {
			foundPath = path
			// Stop walking immediately after finding the file
			return fs.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	if foundPath == "" {
		return "", errors.New("file not found")
	}
	return foundPath, nil
}

func printFileToStdout(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	_, _ = io.Copy(os.Stdout, file)
}

func main() {
	folder := os.Getenv("REPOS_PATH")
	needle := os.Args[1]
	fileName := fmt.Sprintf("values.%s.yaml", needle)
	filePath, err := FindFile(folder, fileName)
	if err == nil && filePath != "" {
		printFileToStdout(filePath)
	}
}
