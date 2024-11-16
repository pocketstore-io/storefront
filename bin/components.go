package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// maxLines defines the line limit for Vue files
const maxLines = 50

// checkFileLines checks if a file has more than the specified number of lines
func checkFileLines(filePath string) (bool, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, 0, err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if lineCount > maxLines {
		return true, lineCount, nil
	}

	if err := scanner.Err(); err != nil {
		return false, 0, err
	}

	return false, lineCount, nil
}

// walkDir recursively walks through the directory and checks Vue files
func walkDir(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-files and non-.vue files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".vue") {
			return nil
		}

		// Check file line count
		isLong, lines, err := checkFileLines(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return nil
		}

		// Print file info if longer than the limit
		if isLong {
			fmt.Printf("File: %s - Lines: %d (exceeds %d lines)\n", path, lines, maxLines)
		}

		return nil
	})
}

func main() {
	// Replace this with the path to your components folder
	componentsDir := "./components"

	fmt.Printf("Checking Vue files in: %s\n", componentsDir)
	if err := walkDir(componentsDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
