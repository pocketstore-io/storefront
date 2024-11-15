#!/bin/bash

# Define the API endpoint URL and output file
API_URL="https://admin.pocketstore.io/api/collections/settings/records"
OUTPUT_FILE="configuration/settings.json"

# Fetch data from PocketBase API
response=$(curl -s "$API_URL")

mkdir -p configuration

# Check if the API call was successful
if [ $? -ne 0 ]; then
    echo "Error fetching data from PocketBase API."
    exit 1
fi

# Extract only the "items" array and save it to settings.json
echo "$response" | jq '.items' > "$OUTPUT_FILE"

# Confirm that the file has been saved
if [ $? -eq 0 ]; then
    echo "Data saved to $OUTPUT_FILE"
else
    echo "Failed to save data to $OUTPUT_FILE"
    exit 1
fi
