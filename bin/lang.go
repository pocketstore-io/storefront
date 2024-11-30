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

// Config represents the structure of the JSON config file
type Config struct {
	Domain string `json:"domain"`
}

// Record represents the PocketBase record structure
type Record struct {
	Key         string `json:"key"`
	Translation string `json:"translated"`
	Lang        string `json:"lang"`
}

// Function to fetch the domain from pocketstore.json
func getDomainFromConfig(configFile string) (string, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return "", fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return "", fmt.Errorf("error decoding config file: %w", err)
	}

	if config.Domain == "" {
		return "", fmt.Errorf("domain not found in config file")
	}

	return config.Domain, nil
}

// Function to fetch data from the PocketBase API
func fetchPocketBaseData(baseURL, collectionName, authToken string) ([]Record, error) {
	url := fmt.Sprintf("%s/api/collections/%s/records?perPage=1000", baseURL, collectionName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if authToken != "" {
		req.Header.Add("Authorization", "Bearer "+authToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", string(body))
	}

	var result struct {
		Items []Record `json:"items"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

// Function to build a nested JSON object
func buildNestedJSON(records []Record) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, record := range records {
		keys := strings.Split(record.Key, ".")
		current := result

		for i, key := range keys {
			if i == len(keys)-1 {
				current[key] = record.Translation
			} else {
				if _, exists := current[key]; !exists {
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
		var filteredRecords []Record
		for _, record := range records {
			if record.Lang == lang {
				filteredRecords = append(filteredRecords, record)
			}
		}

		nestedJSON, err := buildNestedJSON(filteredRecords)
		if err != nil {
			return fmt.Errorf("error building nested JSON for language %s: %w", lang, err)
		}

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}

		filePath := filepath.Join(outputDir, lang+".json")
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", filePath, err)
		}
		defer file.Close()

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
	// Load domain from the pocketstore.json configuration
	configFilePath := "pocketstore.json"
	domain, err := getDomainFromConfig(configFilePath)
	if err != nil {
		fmt.Printf("Error reading domain from config file: %v\n", err)
		return
	}

	fmt.Printf("Using domain: %s\n", domain)

	// PocketBase collection details
	collectionName := "translations"
	authToken := "" // Replace with your API token if necessary

	// Language codes
	langs := []string{"en", "de"}

	// Output directory for translations
	outputDir := "i18n/locales"

	// Fetch translations from PocketBase
	records, err := fetchPocketBaseData(domain, collectionName, authToken)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Save translations
	err = saveTranslationsByLanguage(records, langs, outputDir)
	if err != nil {
		fmt.Printf("Error saving translations: %v\n", err)
		return
	}
}
