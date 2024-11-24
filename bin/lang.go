package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Record represents the PocketBase record structure
type Record struct {
	Key         string `json:"key"`
	Translation string `json:"translated"`
	Lang        string `json:"lang"` // Assumes records include a language field
}

// Function to fetch data from the PocketBase API
func fetchPocketBaseData(baseURL, collectionName, authToken string) ([]Record, error) {
	url := fmt.Sprintf("%s/api/collections/%s/records?perPage=1000", baseURL, collectionName)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add authorization if token is provided
	if authToken != "" {
		req.Header.Add("Authorization", "Bearer "+authToken)
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", string(body))
	}

	// Unmarshal JSON into the records slice
	var result struct {
		Items []Record `json:"items"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

// Function to convert records to a nested JSON object
func buildNestedJSON(records []Record) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, record := range records {
		keys := strings.Split(record.Key, ".")
		current := result

		for i, key := range keys {
			if i == len(keys)-1 {
				// Final key: check for conflicts
				if existing, exists := current[key]; exists {
					// Conflict: log a warning and skip
					if _, ok := existing.(map[string]interface{}); ok {
						return nil, fmt.Errorf(
							"conflict: key '%s' is a map but tried to assign a string",
							record.Key,
						)
					}
				}
				current[key] = record.Translation
			} else {
				// Intermediate key: ensure it's a map
				if existing, exists := current[key]; exists {
					if _, ok := existing.(map[string]interface{}); !ok {
						// Conflict: log a warning and overwrite
						return nil, fmt.Errorf(
							"conflict: key '%s' is a string but tried to assign a map",
							strings.Join(keys[:i+1], "."),
						)
					}
				} else {
					// Key doesn't exist, create a new map
					current[key] = make(map[string]interface{})
				}
				current = current[key].(map[string]interface{})
			}
		}
	}

	return result, nil
}

// Function to save translations by language
func saveTranslationsByLanguage(records []Record, langs []string, outputDir string) error {
	for _, lang := range langs {
		// Filter records by the current language
		var filteredRecords []Record
		for _, record := range records {
			if record.Lang == lang {
				filteredRecords = append(filteredRecords, record)
			}
		}

		// Build nested JSON for this language
		nestedJSON, err := buildNestedJSON(filteredRecords)
		if err != nil {
			return fmt.Errorf("error building nested JSON for language %s: %w", lang, err)
		}

		// Create output directory if it doesn't exist
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}

		// Create the file
		filePath := filepath.Join(outputDir, lang+".json")
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", filePath, err)
		}
		defer file.Close()

		// Use a custom encoder to write JSON without escaping HTML characters
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		encoder.SetEscapeHTML(false)

		if err := encoder.Encode(nestedJSON); err != nil {
			return fmt.Errorf("error writing JSON to file %s: %w", filePath, err)
		}

		fmt.Printf("Saved translations for language %s to %s\n", lang, filePath)
	}

	return nil
}

func main() {
	// PocketBase instance details
	baseURL := "https://admin.pocketstore.io"
	collectionName := "translations"
	authToken := "" // Replace with your actual PocketBase API token

	// List of language codes
	langs := []string{"en", "de"} // Add or modify the language codes as needed

	// Output directory
	outputDir := "i18n/locales"

	// Fetch data from PocketBase
	records, err := fetchPocketBaseData(baseURL, collectionName, authToken)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// Save translations by language
	err = saveTranslationsByLanguage(records, langs, outputDir)
	if err != nil {
		fmt.Println("Error saving translations:", err)
		return
	}
}
