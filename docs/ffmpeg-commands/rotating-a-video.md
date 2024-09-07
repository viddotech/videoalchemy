# Rotating a Video

Rotate a video by 90, 180, or 270 degrees using FFmpeg, useful for correcting orientation after recording or for creative effects.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Rotating a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    codecs:
      - video_filters:
          - name: transpose
            value: 1
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```


## Command

```bash
ffmpeg -i input.mp4 -vf "transpose=1" output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "transpose=1"`**: Applies the transpose filter to rotate the video. The `transpose` parameter can take the following values:
  - `0`: Rotate by 90 degrees counterclockwise and vertically flip.
  - `1`: Rotate by 90 degrees clockwise.
  - `2`: Rotate by 90 degrees counterclockwise.
  - `3`: Rotate by 90 degrees clockwise and vertically flip.
  Replace `1` with your desired rotation option.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter argument**: Happens if an invalid argument is passed to the `transpose` filter. Ensure the value is within the valid range (0-3).
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## GPU Acceleration Command

Rotating a video is primarily a CPU-bound process due to the nature of the operation. While GPU acceleration is not directly applicable to the rotation itself, encoding the video during the process can benefit from GPU acceleration. For Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -vf "transpose=1" -c:v h264_nvenc output.mp4
```


## Additional Information

- **Aspect Ratio**: Rotating the video may change the aspect ratio. Consider adjusting the aspect ratio or scaling the video to maintain the original dimensions using additional filters like `scale`.
- **Multiple Rotations**: To rotate the video by 180 degrees or apply multiple rotations, you might chain the transpose filter (e.g., `-vf "transpose=2,transpose=2"` for 180 degrees).
- **Quality**: If re-encoding is required, consider specifying bitrate or quality settings to maintain high video quality. For example, use `-crf 20` with `-c:v libx264` for H.264 encoding.
- **Audio**: Rotating a video does not affect the audio track. If your video rotation results in an orientation change that impacts the viewing experience, consider if any adjustments to the audio channels are necessary.
