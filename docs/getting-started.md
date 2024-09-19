
# Getting Started

## Introduction
Welcome to VideoAlchemy! This guide will walk you through the installation process and help you set up your first video processing tasks. Follow these steps to get started quickly.

## Installation Requirements
Before installing VideoAlchemy, ensure your system meets the following requirements:

1. **FFmpeg**: VideoAlchemy uses FFmpeg as the core processing engine. You need to have FFmpeg installed on your system.
    - Visit the [FFmpeg download page](https://ffmpeg.org/download.html) for installation instructions specific to your operating system.
    - **macOS**: You can use Homebrew to install FFmpeg:

    ```bash
    brew install ffmpeg
    ```
      
    - **Linux**: For Debian-based distributions, install FFmpeg with:
   
      ```bash
      sudo apt-get install ffmpeg
      ```

---

## Installing VideoAlchemy

To install VideoAlchemy on your system, follow the steps below based on your platform:

### macOS and Linux (via script)
1. Open your terminal.
2. Run the following command to install VideoAlchemy:

```bash
version=$(curl -s https://api.github.com/repos/viddotech/videoalchemy/releases/latest | grep -oP '"tag_name": "\K(.*)(?=")') && curl -o setup.sh "https://raw.githubusercontent.com/viddotech/videoalchemy/main/scripts/setup.sh" && chmod +x setup.sh && sudo ./setup.sh "$version" && rm -rf setup.sh
```

This script will download and install the latest version of VideoAlchemy.

### Alpine Linux
If you’re using Alpine Linux, you can follow the same steps as above. The installation script is compatible with Alpine without the need for `zsh` or `fish` shells.

### From Source
Alternatively, you can clone the VideoAlchemy repository and build the tool from the source:

1. Clone the repository:
   ```bash
   git clone https://github.com/viddotech/videoalchemy.git
   ```
2. Navigate to the project directory:
   ```bash
   cd videoalchemy
   ```
3. Build the tool:
   ```bash
   make build
   ```
4. After building, you can move the binary to your system’s `$PATH` to use it globally.

---

## Verifying the Installation
Once VideoAlchemy is installed, you can verify the installation by running:
```bash
videoalchemy --version
```

This command will display the installed version of VideoAlchemy.

---

## Next Steps
After installation, you're ready to start processing videos with VideoAlchemy! To begin:

1. Create a YAML file with your video tasks.
2. Run the following command to execute your tasks:
   ```bash
   videoalchemy compose -f viddo-compose.yaml
   ```
Check out the [Guides](guides.md) section for detailed examples and use cases.
