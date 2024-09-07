# Convert Video to 4K

Upscale a video to 4K resolution using FFmpeg, enhancing its visual quality for displays supporting 4K resolution.

## VideoAlchemy Compose File

```yaml
version: 1  # Schema version of viddo-compose

generate_path: "./generated"  # Directory of log and command files

tasks:
  - name: Convert Video to 4K
    command: ffmpeg
    inputs:
      - id: input_10
        source: 'input.mp4'
    outputs:
      - id: output_10
        overwrite: true
        source: 'output.mp4'
    codecs:
      - video_filters:
          - name: scale
            value: 3840:2160
```

## Command

```bash
ffmpeg -i input.mp4 -vf "scale=3840:2160" output.mp4
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-vf "scale=3840:2160"`: Applies the scale video filter to resize the video to 4K resolution, which is 3840 by 2160 pixels.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the file path is correct.
- **Invalid scale settings**: Happens if the specified resolution does not match the input video's aspect ratio. Adjust the scale parameters or consider adding padding to maintain the aspect ratio.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## Additional Information

- **Aspect Ratio Preservation**: To maintain the original aspect ratio while scaling, use `-1` for one of the dimensions (e.g., `scale=3840:-1` or `scale=-1:2160`). FFmpeg will automatically calculate the other dimension.
- **Quality Consideration**: Upscaling to 4K may introduce visual artifacts or blur due to the increase in resolution. Consider using additional filters like `unsharp` for sharpening or `zscale` for high-quality scaling.
- **Encoding Settings**: To ensure high quality for the upscaled 4K video, use high-quality encoding settings, such as `-c:v libx264 -preset slow -crf 18` for H.264 encoding.
- **Performance and File Size**: Upscaling and encoding to 4K are resource-intensive processes and can result in large file sizes. Make sure your system has adequate processing power and disk space. Adjust encoding settings to find a balance between quality and file size.
- **Audio Handling**: This command focuses on video upscaling, and the audio stream will be copied without changes. If necessary, audio parameters can be adjusted with `-c:a` for codec and `-b:a` for bitrate.
