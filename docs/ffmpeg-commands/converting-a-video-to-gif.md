# Converting a Video to GIF

Create a GIF from a video file using FFmpeg, ideal for short clips or animations.

## VideoAlchemy Compose File

```yaml
version: 1  # Schema version of viddo-compose

generate_path: "./generated"  # Directory of log and command files

tasks:
  - name: Converting a Video to GIF
    command: ffmpeg
    inputs:
      - id: input_12
        source: 'input.mp4'
    outputs:
      - id: output_12
        overwrite: true
        source: 'output.gif'
    codecs:
      - video_filters:
          - name: fps
            value: 10
          - name: scale
            value: 320:-1
            flags: lanczos
        codec_name:
          video: gif
```

## Command

```bash
ffmpeg -i input.mp4 -vf "fps=10,scale=320:-1:flags=lanczos" -c:v gif output.gif
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "fps=10,scale=320:-1:flags=lanczos"`**: Applies video filters to set the frame rate to 10 frames per second, scale the width to 320 pixels while maintaining the aspect ratio, and use the Lanczos filter for scaling.
- **`-c:v gif`**: Sets the video codec to GIF for the output file.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter settings**: Occurs if the specified filter graph (e.g., scaling dimensions or fps) is not valid. Ensure the filter syntax and values are correct.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Converting a video to GIF primarily involves frame extraction and palette generation, processes which do not significantly benefit from GPU acceleration. Therefore, there is no direct GPU acceleration command for converting videos to GIFs with FFmpeg. However, for initial video processing steps (like decoding or scaling) before converting to a GIF, GPU acceleration can be utilized if necessary:

For Nvidia GPUs (example of decoding with GPU, though the GIF conversion process itself is CPU-bound):

```bash
ffmpeg -hwaccel cuda -i input.mp4 -vf "fps=10,scale=320:-1:flags=lanczos,format=rgb24" -c:v gif output.gif
```


## Additional Information

- **Frame Rate (fps)**: Lowering the frame rate can reduce the GIF size. Adjust the `fps` value according to your needs.
- **Scaling**: Adjust the scale to control the GIF size and quality. Smaller dimensions result in smaller file sizes.
- **Optimization**: After creating a GIF, consider using tools like `gifsicle` for further optimization, which can reduce file size without significantly affecting quality.
- **Palette Generation**: For high-quality GIFs, consider using a two-pass approach where you first generate a palette based on the most common colors in the video and then use that palette to create the GIF. This is achieved with the `palettegen` and `paletteuse` filters in FFmpeg.
