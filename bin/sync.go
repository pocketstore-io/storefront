package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Config represents the configuration from pocketbase.json
type Config struct {
	Domain    string            `json:"domain"`
	AuthToken string            `json:"authToken"`
	Settings  map[string]interface{} `json:"settings"`
}

func main() {
	// Load configuration from pocketbase.json
	configFile := "pocketstore.json"
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// API URL for the "settings" collection
	apiURL := fmt.Sprintf("%s/api/collections/settings/records", config.Domain)

	// First, delete all existing settings
	err = deleteAllSettings(apiURL, config.AuthToken)
	if err != nil {
		fmt.Println("Error deleting settings:", err)
		os.Exit(1)
	}

	// Upload each setting to PocketBase
	for key, value := range config.Settings {
		// Marshal value into a JSON object
		setting := map[string]interface{}{
			"key":   key,
			"value": value,
		}
		err := uploadSetting(apiURL, config.AuthToken, setting)
		if err != nil {
			fmt.Printf("Error uploading setting %s: %v\n", key, err)
			continue
		}
		fmt.Printf("Uploaded setting: %s\n", key)
	}
}

// loadConfig reads and parses the configuration file
func loadConfig(configFile string) (*Config, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}

// deleteAllSettings deletes all settings from the PocketBase collection
func deleteAllSettings(apiURL, authToken string) error {
	// Create the HTTP request to fetch all records
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("error creating GET request to fetch records: %w", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", "Bearer "+authToken)

	// Send the GET request to retrieve existing records
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending GET request: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error response from server while fetching records: %s", string(body))
	}

	// Parse the response body to get records
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	// Extract record IDs
	items, ok := result["items"].([]interface{})
	if !ok {
		return fmt.Errorf("unexpected response structure: no items found")
	}

	// Delete each record by its ID
	for _, item := range items {
		record, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		recordID, ok := record["id"].(string)
		if !ok {
			continue
		}

		deleteURL := fmt.Sprintf("%s/%s", apiURL, recordID)
		err := deleteSetting(deleteURL, authToken)
		if err != nil {
			fmt.Printf("Error deleting setting with ID %s: %v\n", recordID, err)
		} else {
			fmt.Printf("Deleted setting with ID %s\n", recordID)
		}
	}

	return nil
}

// deleteSetting deletes a single setting record by ID
func deleteSetting(apiURL, authToken string) error {
	// Create the HTTP DELETE request
	req, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return fmt.Errorf("error creating DELETE request: %w", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", "Bearer "+authToken)

	// Send the DELETE request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending DELETE request: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful deletion
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error response from server while deleting record: %s", string(body))
	}

	return nil
}

// uploadSetting uploads a single setting to PocketBase
func uploadSetting(apiURL, authToken string, setting map[string]interface{}) error {
	// Marshal the setting into JSON
	settingJSON, err := json.Marshal(setting)
	if err != nil {
		return fmt.Errorf("error marshaling setting to JSON: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(settingJSON))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error response from server: %s", string(body))
	}

	return nil
}