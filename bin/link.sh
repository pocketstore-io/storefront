#!/bin/bash

# Define the source and destination paths
SOURCE="$(pwd)/.nginx/nuxt.conf"
DESTINATION="/etc/nginx/sites-enabled/nuxt.conf"

# Check if the source file exists
if [[ ! -f "$SOURCE" ]]; then
  echo "Error: Source file $SOURCE does not exist."
  exit 1
fi

# Link the configuration file
echo "Linking $SOURCE to $DESTINATION"
sudo ln -sf "$SOURCE" "$DESTINATION"

# Reload nginx to apply the new configuration
echo "Reloading nginx..."
sudo systemctl reload nginx

echo "Done."