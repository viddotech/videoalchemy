# Resizing Videos

Resize video files to a different resolution using FFmpeg.

## VideoAlchemy Compose File
    
```yaml
version: 1

generate_path: "./generated"

tasks:
- name: Resizing Video
  command: ffmpeg
  inputs:
    - id: input_video
      source: 'input.mp4'
      streams:
    - video_filters:
        - name: scale
          value: "1280:720"
  outputs:
    - id: resized_output
      overwrite: true
      source: 'output.mp4'
```


## Command

```bash
ffmpeg -i input.mp4 -vf "scale=1280:720" output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "scale=1280:720"`**: Applies a video filter to scale the video to the specified width and height. Replace `1280:720` with your desired resolution.
- **`output.mp4`**: Specifies the output video file. Replace `output.mp4` with your desired output file name.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid scale dimensions**: Occurs if the specified dimensions are not valid. Ensure that the width and height are positive integers.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

For Nvidia GPUs, use:

```bash
ffmpeg -hwaccel cuda -i input.mp4 -vf "scale_cuda=1280:720" output.mp4
```


Note: GPU acceleration for video scaling may require specific hardware support and FFmpeg configurations.

## Additional Information

- **Aspect Ratio**: To maintain the original aspect ratio while scaling, use `-1` for one of the dimensions (e.g., `scale=1280:-1` to scale the width to 1280 pixels and adjust the height proportionally).
- **Quality**: The quality of the resized video can be affected by the scaling algorithm. You can specify the scaling algorithm using the `flags` option (e.g., `scale=1280:720:flags=lanczos` for Lanczos resampling).
- **Performance**: Resizing videos can be computationally intensive. Using GPU acceleration (if available) can significantly speed up the process.
