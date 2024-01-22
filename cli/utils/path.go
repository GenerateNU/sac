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

	// Find the closest directory containing "sac-cli" (the root directory)
	rootDir, err := FindRootDir(currentDir)
	if err != nil {
		return "", err
	}

	return rootDir, nil
}

func FindRootDir(dir string) (string, error) {
	// Check if "sac-cli" exists in the current directory
	mainGoPath := filepath.Join(dir, "sac-cli")
	_, err := os.Stat(mainGoPath)
	if err == nil {
		// "sac-cli" found, this is the root directory
		return dir, nil
	}

	// If not found, go up one level
	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		// Reached the top without finding "sac-cli"
		return "", fmt.Errorf("could not find root directory containing sac-cli")
	}

	// Recursively search in the parent directory
	return FindRootDir(parentDir)
}