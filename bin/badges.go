package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Badge URL template
const badgeURLTemplate = "https://img.shields.io/badge/%s-%s-brightgreen"

// Folder to save badges
const badgeFolder = ".github/badges"

// Read package.json and parse dependencies
func readPackageJSON(filePath string) (map[string]string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data struct {
		Dependencies map[string]string `json:"dependencies"`
	}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}
	return data.Dependencies, nil
}

// Clean up version strings to keep only numbers and dots
func cleanVersion(version string) string {
	re := regexp.MustCompile(`[^\d.]`)
	return re.ReplaceAllString(version, "")
}

// Replace dashes in package names with underscores and "@" with "#"
func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, "-", "+")
	return name
}

// URL encode a string for safe inclusion in URLs
func urlEncode(input string) string {
	return url.QueryEscape(input)
}

// Generate badge using wget
func generateBadge(fullName, version string) error {
	// Ensure the badge folder exists
	if err := os.MkdirAll(badgeFolder, os.ModePerm); err != nil {
		return err
	}

	// Clean and sanitize inputs
	badgeName := sanitizeName(fullName)
	badgeVersion := cleanVersion(version)

	// URL-encode the badge parameters
	encodedName := urlEncode(badgeName)
	encodedVersion := urlEncode(badgeVersion)

	// Format the badge URL
	url := fmt.Sprintf(badgeURLTemplate, encodedName, encodedVersion)

	// Save the badge to the folder (without version in the name)
	outputFile := filepath.Join(badgeFolder, fmt.Sprintf("%s.svg", badgeName))
	cmd := exec.Command("wget", "-q", "-O", outputFile, url)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download badge for %s: %v", fullName, err)
	}

	fmt.Printf("Generated badge for %s: %s\n", fullName, outputFile)
	return nil
}

func main() {
	// Path to package.json
	packageJSONPath := "package.json"

	// Read dependencies
	dependencies, err := readPackageJSON(packageJSONPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading package.json: %v\n", err)
		os.Exit(1)
	}

	// Generate badges for each dependency without "/"
	for fullName, version := range dependencies {
		if !strings.Contains(fullName, "/") { // Skip names with "/"
			if err := generateBadge(fullName, version); err != nil {
				fmt.Fprintf(os.Stderr, "Error generating badge for %s: %v\n", fullName, err)
			}
		} else {
			fmt.Printf("Skipping dependency with '/' in name: %s\n", fullName)
		}
	}
}
