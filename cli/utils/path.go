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

	// Find the closest directory containing "sac" (the root directory)
	rootDir, err := FindRootDir(currentDir)
	if err != nil {
		return "", err
	}

	return rootDir, nil
}

func FindRootDir(dir string) (string, error) {
	// Check if "sac" exists in the current directory
	mainGoPath := filepath.Join(dir, "sac")
	_, err := os.Stat(mainGoPath)
	if err == nil {
		// "sac" found, this is the root directory
		return dir, nil
	}

	// If not found, go up one level
	parentDir := filepath.Dir(dir)
	if parentDir == dir {
		// Reached the top without finding "sac"
		return "", fmt.Errorf("could not find root directory containing sac")
	}

	// Recursively search in the parent directory
	return FindRootDir(parentDir)
}
