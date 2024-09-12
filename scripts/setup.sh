#!/bin/sh

# Base URL for downloads
BASE_URL="https://github.com/viddotech/videoalchemy/releases/download"

# Default version
DEFAULT_VERSION=""

# Check if a version was provided as a parameter
if [ "$#" -eq 1 ]; then
  VERSION="$1"
else
  VERSION="$DEFAULT_VERSION"
fi

# Construct the full download URL
DOWNLOAD_URL="${BASE_URL}/${VERSION}/videoalchemy"


# Download and install VideoAlchemy
echo "Downloading VideoAlchemy from $DOWNLOAD_URL..."
curl -L -o videoalchemy "$DOWNLOAD_URL"
if [ $? -ne 0 ]; then
  echo "Failed to download VideoAlchemy. Exiting."
  exit 1
fi

chmod +x videoalchemy
if [ $? -ne 0 ]; then
  echo "Failed to make VideoAlchemy executable. Exiting."
  exit 1
fi

sudo mv videoalchemy /usr/local/bin/
if [ $? -ne 0 ]; then
  echo "Failed to move VideoAlchemy to /usr/local/bin. Exiting."
  exit 1
fi

# Shell detection and autocompletion setup
if [ -n "$BASH_VERSION" ]; then
  # Setup autocompletion for bash
  videoalchemy completion bash > /etc/bash_completion.d/videoalchemy
  source /etc/bash_completion.d/videoalchemy
  echo "Autocompletion enabled for bash."

elif [ -n "$ZSH_VERSION" ]; then
  # Setup autocompletion for zsh
  videoalchemy completion zsh > "${fpath[1]}/_videoalchemy"
  autoload -U compinit
  compinit
  echo "Autocompletion enabled for zsh."

elif [ -n "$FISH_VERSION" ]; then
  # Setup autocompletion for fish
  videoalchemy completion fish > ~/.config/fish/completions/videoalchemy.fish
  echo "Autocompletion enabled for fish."

fi

echo "VideoAlchemy installation completed!"
