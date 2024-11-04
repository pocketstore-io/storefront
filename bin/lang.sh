#!/bin/bash

# Variables
API_URL="https://admin.pocketstore.io/api/collections/settings/records"
OUTPUT_DIR="configuration/lang"

# Fetch all records where the code starts with 'lang_'
response=$(curl -s "$API_URL?filter=key~'lang_%'")

# Check if the response contains any items
if [[ $(echo "$response" | jq -e ".items | length") -eq 0 ]]; then
    echo "No matching records found with code starting with 'lang_'"
    exit 1
fi

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Iterate over each item and save it to a file
echo "$response" | jq -c ".items[]" | while read -r item; do
    # Try to use 'code' for filename; fallback to 'key' if 'code' is null
    code=$(echo "$item" | jq -r ".code // empty")
    key=$(echo "$item" | jq -r ".key")
    filename=${code:-$key} # Use 'code' if available; otherwise, use 'key'

    # Define the output file path
    output_file="$OUTPUT_DIR/$filename.json"
    
     if [[ "$filename" == lang_* ]]; then
        additional=$(echo "$item" | jq '.additional') # Extract the object directly
        echo "$additional" > "$output_file"
        echo "Saved data for $filename to $output_file"
    fi
done
