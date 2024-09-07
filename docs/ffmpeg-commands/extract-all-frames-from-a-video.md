# Extract All Frames from a Video

Extract every frame from a video file as individual image files using FFmpeg, useful for detailed editing, analysis, or creating frame-by-frame animations.

## VideoAlchemy Compose File
    
```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Extract All Frames from a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    outputs:
      - id: output_1
        overwrite: true
        source: 'output%d.png'
```

## Command

```bash
ffmpeg -i input.mp4 output%d.png
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `output%d.png`: Defines the pattern for the output image files. `%d` will be replaced by the frame number, creating a sequence of images.

## Possible Errors

- **File not found**: This error occurs if FFmpeg cannot locate the input file. Make sure the file path is correct.
- **Disk space**: Extracting all frames from a video requires a significant amount of disk space, which can be an issue for long videos or those with high frame rates. Ensure there is enough disk space available.
- **Permission denied**: Occurs if FFmpeg lacks the necessary permissions to write the output files. Verify that the destination directory has appropriate permissions.

## Additional Information

- **Image Format**: The example uses PNG for its lossless quality. Other formats like JPG can be used by changing the file extension in the output pattern, but it's important to consider the trade-offs in file size and quality.
- **Frame Rate and Duration**: The video's frame rate and duration should be considered, as high frame rates or long durations will produce many frames.
- **Naming and Sorting**: Using the `%d` pattern ensures frames are numbered sequentially, which helps with easy sorting and accessing individual frames.
- **Storage Considerations**: Given the potential for generating a large number of files, it's wise to plan the storage location carefully. Using an external drive or a directory with ample free space might be advisable.
