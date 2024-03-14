# Basic Video Conversion

Convert video files from one format to another using FFmpeg.

## Command

```bash
ffmpeg -i input.mp4 output.avi
```

## Parameters

- **`-i input.mp4`**: Specifies the input file. Replace `input.mp4` with the path to your source video file.
- **`output.avi`**: Specifies the output file. The extension of the output file determines the output format. Replace `output.avi` with your desired output file name and format.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Unsupported codec**: Occurs if the output format requires a codec not supported or not installed. Ensure that the desired output format is supported.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

For Nvidia GPUs, use:

```bash
ffmpeg -hwaccel cuda -i input.mp4 -c:v h264_nvenc output.avi
```


## Additional Information

- **Codecs and Formats**: The output format is determined by the file extension of the output file (e.g., `.mp4`, `.avi`, `.mkv`). The codec used for encoding can be specified with the `-c:v` parameter (e.g., `-c:v libx264` for H.264).
- **Quality and Compression**: You can control the quality and compression of the output video by adjusting the bitrate (using `-b:v`) or the constant rate factor (using `-crf` for codecs like H.264).
- **Audio**: By default, FFmpeg will also transcode the audio stream. You can specify the audio codec with `-c:a` (e.g., `-c:a aac` for AAC audio) and adjust the audio bitrate with `-b:a`.
- **Compatibility**: Ensure that the chosen codecs and formats are compatible with your intended playback devices or platforms.
