#!/bin/bash

# Script to update the Daytona API client from the official repository
# This fetches the latest version from GitHub without needing a full clone

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
SDK_ROOT="$(dirname "$SCRIPT_DIR")"
APICLIENT_DIR="$SDK_ROOT/apiclient"
TEMP_DIR="/tmp/daytona-apiclient-update-$$"

echo "Updating Daytona API client from GitHub..."

# Create temp directory
mkdir -p "$TEMP_DIR"

# Download the API client as a tar archive
echo "Downloading latest API client..."
curl -L -s https://github.com/daytonaio/daytona/tarball/main | \
    tar xz -C "$TEMP_DIR" --strip=3 "*/libs/api-client-go/"

# Backup current apiclient
if [ -d "$APICLIENT_DIR" ]; then
    echo "Backing up current API client..."
    mv "$APICLIENT_DIR" "$APICLIENT_DIR.backup"
fi

# Move new API client
echo "Installing new API client..."
mv "$TEMP_DIR" "$APICLIENT_DIR"

# Remove the go.mod file since apiclient is part of the main module
echo "Removing apiclient go.mod..."
rm -f "$APICLIENT_DIR/go.mod"

# Clean up backup if successful
if [ -d "$APICLIENT_DIR.backup" ]; then
    echo "Removing backup..."
    rm -rf "$APICLIENT_DIR.backup"
fi

echo "✓ API client updated successfully!"
echo ""
echo "Note: You may need to run 'go mod tidy' to update dependencies."