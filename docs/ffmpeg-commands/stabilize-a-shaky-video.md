# Stabilize a Shaky Video

Stabilize shaky video footage using FFmpeg's `deshake` filter, improving the viewing experience by reducing unwanted camera movements.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Stabilize a Shaky Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    codecs:
      - video_filters:
          - name: deshake
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```

## Command

```bash
ffmpeg -i input.mp4 -vf "deshake" output.mp4
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-vf "deshake"`: Applies the deshake video filter to stabilize the video.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the file path is correct.
- **Filter processing error**: Happens if the deshake filter encounters an issue during the video processing, potentially due to unusual video formats or extreme shakiness. Examine the video properties or consider adjusting filter options.
- **Permission denied**: Arises if FFmpeg lacks the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## Additional Information

- **Filter Options**: The deshake filter offers several options to customize the stabilization process, such as `rx`, `ry` (maximum allowed movement), and `edge` (handling of video edges). For example: `-vf "deshake=rx=10:ry=10"`.
- **Performance Consideration**: Video stabilization is computationally demanding, particularly for videos that are high-resolution or long. Take into account the processing capability of your system and the size of your video.
- **Quality Preservation**: Although stabilization can significantly enhance the usability of shaky footage, it may also lead to slight cropping or quality loss. Assess the output to ensure it aligns with your standards.
- **Alternative Stabilization Tools**: For videos requiring more advanced stabilization, consider using dedicated video editing software that offers more sophisticated stabilization features, as the deshake filter may not effectively handle all types of camera motion.
