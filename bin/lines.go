package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxLineLength = 80 // Set your desired maximum line length here

func main() {
	componentsFolder := "./components"

	// Walk through the components folder
	err := filepath.Walk(componentsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check for .vue files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".vue") {
			fmt.Printf("Checking file: %s\n", path)
			checkLineLengths(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through components folder: %v\n", err)
	}
}

// checkLineLengths checks line lengths in the specified file and prints lines that exceed the max length
func checkLineLengths(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxLineLength {
			fmt.Printf("Line %d in %s exceeds %d characters:\n%s\n", lineNumber, filePath, maxLineLength, line)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
	}
}
