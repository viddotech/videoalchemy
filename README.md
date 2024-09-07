
# VideoAlchemy: Simplified, Structured Video Processing

Welcome to **VideoAlchemy**, an advanced toolkit that offers a more readable and structured way to handle video processing tasks compared to traditional FFmpeg commands. VideoAlchemy enables users to define video processing workflows using a simple YAML-based configuration, enriched with built-in validation to minimize errors and streamline command execution.

Whether you're processing a single video or executing a sequence of commands, VideoAlchemy ensures a smoother and more intuitive experience.

![Screenshot](docs/assets/videoalchemy-demo.gif)

## Table of Contents

- [Introduction](#introduction)
- [Why VideoAlchemy?](#why-videoalchemy)
- [VideoAlchemy Command](#videoalchemy-command)
- [Viddo Compose Tutorial](docs/viddo-compose-tutorial.md)
- [Contributing](#contributing)
- [Upcoming Enhancements and Support](#upcoming-enhancements-and-support)
- [License](#license)


## Introduction

**VideoAlchemy** provides a user-friendly alternative to directly writing complex FFmpeg commands. With a focus on clarity, ease of use, and sequence management, VideoAlchemy allows you to write video processing tasks using readable YAML files. These files come with built-in validation to help avoid common mistakes, making it ideal for both beginners and seasoned professionals.

## Why VideoAlchemy?

VideoAlchemy enhances video processing in several ways:

1. **Readable Attributes**: VideoAlchemy transforms the complexity of FFmpeg command-line syntax into a clear, structured format using YAML attributes. This reduces the chance of errors and improves the legibility of the command.

2. **Rich YAML Validation**: The `viddo-compose.yaml` file provides real-time validation, helping you craft correct FFmpeg commands by offering clear syntax and error-checking mechanisms. Common FFmpeg pitfalls are automatically mitigated, resulting in smoother workflows.

3. **Command Sequencing**: Need to run multiple FFmpeg commands in sequence? VideoAlchemy makes this simple by allowing you to chain commands and set dependencies between tasks within the YAML file. This is especially useful for projects requiring multiple steps, like video conversion followed by audio extraction.

## VideoAlchemy Command

The `videoalchemy` command is a binary written in Go that processes a series of FFmpeg tasks defined in a YAML configuration file. The configuration file is easy to read and write, and it allows you to execute FFmpeg commands sequentially without manually chaining them.

### Run VideoAlchemy

```bash
videoalchemy compose -f viddo-compose.yaml
```

### Example `viddo-compose.yaml` File

```yaml
version: 1  # Schema version of viddo-compose

generate_path: "./generated"  # Directory of log and command files

tasks:  
  - name: Convert to AVI  
    command: ffmpeg  
    inputs:  
      - id: input_video  
        source: 'input.mp4'  
    outputs:  
      - id: output_avi  
        source: 'output.avi'  
        overwrite: true  

  - name: Extract Audio  
    command: ffmpeg  
    inputs:  
      - id: output_avi  
        output_id: output_avi  # Reference from the previous task  
    codecs:  
      - codec_name:
          audio: copy  
        video_none: true  
    outputs:  
      - id: audio_only  
        source: 'output_audio.mp3'  
        overwrite: true  
```

### Key Features:

- **Readable and Organized**: Rather than memorizing FFmpeg's complex flags and options, you can clearly define inputs, outputs, codecs, and filters in an intuitive format.
- **Command Sequencing**: Define dependencies between tasks using `run_after` to ensure commands run in the desired order.
- **Error Prevention**: YAML validation ensures that FFmpeg commands are properly formed, helping to avoid errors early in the workflow.

## Contributing

We welcome contributions from the community to make VideoAlchemy better! Here's how you can get involved:

- **Improve Readability**: Suggest improvements to the way we structure attributes for video processing.
- **Expand Validation**: Help us enhance our YAML validation to cover more complex FFmpeg use cases.
- **Create New Commands**: Propose and implement new commands and workflows.

## Upcoming Enhancements and Support

We are continuously working to enhance VideoAlchemy. Here are some exciting features we plan to introduce in the future:

1. **Support MPEG-DASH**: Extend the toolkit to support MPEG-DASH, enabling adaptive streaming for high-quality video delivery over the internet.

2. **Complete FFmpeg Parameters**: Expand the YAML schema to support all FFmpeg parameters, providing users with the full range of FFmpeg's capabilities directly within the `viddo-compose.yaml` file.

3. **Cloud Storage Integration**: Add support for various cloud storage solutions such as Azure Storage, AWS S3, Minio, and more, allowing users to specify these as source files for their video processing tasks.

4. **Improved Validation**: Enhance the built-in validation mechanisms to cover more complex scenarios and edge cases, ensuring that all FFmpeg commands are correctly formed and executed without errors.

---

Your support keeps this project growing. Consider donating to help us continue developing VideoAlchemy.

- [Bitcoin (BTC)](donate/donate.md)
- [Ethereum (ETH)](donate/donate.md)

## License

VideoAlchemy is licensed under the [MIT License](LICENSE). Contributions are licensed under the same terms.

