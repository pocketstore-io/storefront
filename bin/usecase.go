package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	maxLines  = 100
	targetDir = "usecase"
)

func main() {
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .js or .ts extension
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".js") || strings.HasSuffix(info.Name(), ".ts")) {
			lineCount, err := countLines(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return nil
			}

			// If the file exceeds maxLines, print the filename and line count
			if lineCount > maxLines {
				fmt.Printf("File: %s, Lines: %d\n", path, lineCount)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
	}
}

// countLines counts the number of lines in a file
func countLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}