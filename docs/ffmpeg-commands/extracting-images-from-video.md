# Extracting Images from Video

Extract frames from a video file as individual image files using FFmpeg, useful for thumbnails or analysis.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Extracting Images from Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    codecs:
      - video_filters:
          - name: fps
            value: 1
    outputs:
      - id: output_1
        overwrite: true
        source: 'output%d.png'
```

## Command

```bash
ffmpeg -i input.mp4 -vf "fps=1" output%d.png
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "fps=1"`**: Applies a video filter to extract one frame per second. Adjust the number after `fps=` to change the extraction rate (e.g., `fps=0.5` for one frame every 2 seconds).
- **`output%d.png`**: Specifies the output file pattern. `%d` will be replaced by the frame number. Replace `output` with your desired base file name, and `.png` with your preferred image format if necessary.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter settings**: Occurs if the specified filter graph (e.g., fps value) is not valid. Ensure the filter syntax and values are correct.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output directory. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Extracting images from a video is primarily a CPU-bound process, and GPU acceleration may not significantly improve performance. However, if decoding the input video is resource-intensive, GPU acceleration can be beneficial for this step:

For Nvidia GPUs:

```bash
ffmpeg -hwaccel cuda -i input.mp4 -vf "fps=1" output%d.png
```


## Additional Information

- **Image Format**: You can choose different image formats (e.g., jpg, png) based on your needs. PNG offers lossless compression, while JPEG might be preferable for smaller file sizes.
- **Extraction Rate**: Adjusting the `fps` value in the video filter allows for flexible control over how many frames are extracted. A lower value extracts fewer frames, suitable for longer videos or when only a few frames are needed.
- **Frame Accurate Extraction**: To extract specific frames, consider using the `select` filter (e.g., `-vf "select=eq(n\,100)"` to extract only the 101st frame) for more precise control.
- **Disk Space**: Be mindful of disk space when extracting a large number of frames, especially at high resolutions or in lossless formats like PNG.
