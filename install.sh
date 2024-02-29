#!/bin/bash

# Define the URL and the download location within the home directory
URL="https://github.com/dmikey/go-vite-app-p2p/releases/download/v0.0.9/myapp-darwin.arm64.tar.gz"
TARGET_DIR="$HOME/.local/bin"
APP_NAME="myapp"

# Create target directory if it doesn't exist
mkdir -p "$TARGET_DIR"

# Download the file to /tmp
wget -O /tmp/myapp.tar.gz "$URL"

# Navigate to /tmp
cd /tmp

# Extract the tar.gz file
tar -xzf myapp.tar.gz

# Move the binary to the target directory
mv $APP_NAME "$TARGET_DIR"

# Change permissions to make it executable
chmod +x "$TARGET_DIR/$APP_NAME"

echo "Installation complete. The $APP_NAME is now available in $TARGET_DIR"
echo "To make $APP_NAME available from any terminal session, add the following line to your .bashrc or .zshrc:"
echo "export PATH=\"\$PATH:$TARGET_DIR\""
