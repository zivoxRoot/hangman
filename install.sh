#!/bin/bash

set -e  # Exit on error

# Define variables
APP_NAME="hangman"
INSTALL_DIR="/usr/local/bin"
GITHUB_REPO="zivoxRoot/hangman"
LATEST_RELEASE="0.0.0"
DOWNLOAD_URL="https://github.com/$GITHUB_REPO/raw/refs/heads/main/bin/linux/hangman"
DEFAULT_WORDLIST_URL="https://raw.githubusercontent.com/$GITHUB_REPO/refs/heads/main/assets/words.txt"
DEFAULT_WORDLIST_PATH="$HOME/.config/hangman"

# Download and install the executable
echo "Downloading $APP_NAME and copying it to $INSTALL_DIR..."
sudo curl -L -o "$INSTALL_DIR/$APP_NAME" "$DOWNLOAD_URL"

# Set permissions
sudo chmod +x "$INSTALL_DIR/$APP_NAME"

# Create the folder for the default wordlist
echo "Creating $DEFAULT_WORDLIST_PATH..."
mkdir -p $DEFAULT_WORDLIST_PATH

# Copy the default wordlist to its location
echo "Copying the default wordlist to $DEFAULT_WORDLIST_PATH"
sudo curl -L -o "$DEFAULT_WORDLIST_PATH/words.txt" "$DEFAULT_WORDLIST_URL"

# Ensure the directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Adding $INSTALL_DIR to PATH..."
    echo 'export PATH="$INSTALL_DIR:$PATH"' >> "$HOME/.bashrc"
    echo 'export PATH="$INSTALL_DIR:$PATH"' >> "$HOME/.zshrc"
    export PATH="$INSTALL_DIR:$PATH"
fi

echo "$APP_NAME installed successfully in $INSTALL_DIR"
echo "Run '$APP_NAME' to start using it!"
