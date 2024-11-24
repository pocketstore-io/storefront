package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define the API endpoint URL and output file
	apiURL := "https://admin.pocketstore.io/api/collections/settings/records"
	outputFile := "configuration/settings.json"

	// Fetch data from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data from PocketBase API:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	// Parse the response to extract the "items" array
	var parsedResponse map[string]interface{}
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		fmt.Println("Error parsing JSON response:", err)
		os.Exit(1)
	}

	items, ok := parsedResponse["items"]
	if !ok {
		fmt.Println("Error: 'items' key not found in API response")
		os.Exit(1)
	}

	// Create the output directory if it doesn't exist
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}

	// Write the "items" array to the output file
	itemsJSON, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling 'items' to JSON:", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outputFile, itemsJSON, 0644); err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	// Confirm success
	fmt.Printf("Data saved to %s\n", outputFile)
}
