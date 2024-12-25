#!/bin/sh

# Base URL for downloads
BASE_URL="https://github.com/viddotech/videoalchemy/releases/download"

# Default version
DEFAULT_VERSION="v0.0.1-alpha" # Replace with your default version if needed

# Check if a version was provided as a parameter
if [ "$#" -eq 1 ]; then
  VERSION="$1"
else
  VERSION="$DEFAULT_VERSION"
fi

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture to expected naming
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64) ARCH="arm64" ;;
  armv7l) ARCH="armv7" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Construct the full download URL for the specific binary
DOWNLOAD_URL="${BASE_URL}/${VERSION}/videoalchemy-${OS}-${ARCH}"

# Download and install VideoAlchemy
echo "Downloading VideoAlchemy for ${OS}-${ARCH} from $DOWNLOAD_URL..."
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

mv videoalchemy /usr/local/bin/
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
