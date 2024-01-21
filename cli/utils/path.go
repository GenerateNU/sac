package utils

import (
	"fmt"
	"os"
	"path/filepath"
)


func GetRootDir() (string, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Find the closest directory containing "install.sh" (the root directory)
	rootDir, err := FindRootDir(currentDir)
	if err != nil {
		return "", err
	}

	return rootDir, nil
}

func FindRootDir(dir string) (string, error) {
	// Check if "main.go" exists in the current directory
	mainGoPath := filepath.Join(dir, "install.sh")
	_, err := os.Stat(mainGoPath)
	if err == nil {
		// "main.go" found, this is the root directory
		return dir, nil
	}

	// If not found, go up one level
	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		// Reached the top without finding "main.go"
		return "", fmt.Errorf("could not find root directory containing main.go")
	}

	// Recursively search in the parent directory
	return FindRootDir(parentDir)
}