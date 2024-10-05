package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestReadCurrentDirectory(t *testing.T) {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Traverse all files in the current directory and subdirectories
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Print the path of each file or directory
		fmt.Printf("Found: %s\n", path)
		return nil
	})

	if err != nil {
		t.Errorf("Error walking the path %q: %v", dir, err)
	}
}
