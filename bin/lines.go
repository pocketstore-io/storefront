package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxLines = 80 // Set your desired maximum line count here

func main() {
	componentsFolder := "./components"

	// Walk through the components folder
	err := filepath.Walk(componentsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check for .vue files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".vue") {
			if countLines(path) > maxLines {
				fmt.Println(path) // Output file name if line count exceeds maxLines
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through components folder: %v\n", err)
	}
}

// countLines counts the number of lines in the specified file
func countLines(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
	}
	return lineCount
}
