# Combining Videos

Concatenate or join multiple video files into a single file using FFmpeg.

## Command

For files with the same codecs:

```bash
ffmpeg -f concat -safe 0 -i input.txt -c copy output.mp4
```

For files that need re-encoding:

```bash
ffmpeg -i "concat:input1.mp4|input2.mp4" -c:v libx264 -c:a aac output.mp4
```


## Parameters

- **`-f concat`**: Specifies the use of the concat demuxer, which is required for combining files.
- **`-safe 0`**: Allows the use of absolute file paths in the input file list.
- **`-i input.txt`**: Specifies the input file, which contains a list of files to concatenate. `input.txt` should contain lines in the format: `file 'path/to/file1.mp4'` on each line.
- **`-c copy`**: Copies the video and audio codecs from the input files without re-encoding. Use this option for fast processing when all input files have the same codecs.
- **`-c:v libx264 -c:a aac`**: Specifies the video and audio codecs for re-encoding. Use this option if your input files have different codecs or formats.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate one of the input files. Ensure all paths in the input list are correct.
- **Codec mismatch**: Occurs if the input files have different codecs and `-c copy` is used. Ensure all input files use the same codecs or re-encode them using `-c:v` and `-c:a`.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Combining videos with GPU acceleration involves re-encoding. Here's an example for Nvidia GPUs:

```bash
ffmpeg -f concat -safe 0 -i input.txt -c:v h264_nvenc -c:a aac output.mp4
```


## Additional Information

- **Input File Format**: The input text file for the concat demuxer should list each file to concatenate on a new line, prefixed with `file '`, and followed by the file path and a closing `'`.
- **Compatibility**: Ensure all video files have the same resolution and frame rate before concatenating to avoid playback issues in the output file.
- **Re-encoding**: Re-encoding can affect video quality. To maintain quality, adjust bitrate or quality settings appropriately (e.g., using `-b:v` for video bitrate or `-crf` for constant rate factor with x264 and x265 codecs).
- **Audio Sync**: Pay attention to audio synchronization. Concatenating files with varying audio formats or settings can result in audio sync issues in the output file.
