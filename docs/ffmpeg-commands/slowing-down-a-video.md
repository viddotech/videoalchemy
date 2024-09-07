# Slowing Down a Video

Decrease the playback speed of a video file using FFmpeg, ideal for detailed analyses or creating slow-motion effects.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Slowing Down a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    codecs:
      - video_filters:
          - name: setpts
            value: '2.0*PTS'
        audio_none: true
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```

## Command

```bash
ffmpeg -i input.mp4 -filter:v "setpts=2.0*PTS" -an output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-filter:v "setpts=2.0*PTS"`**: Applies a video filter to adjust the presentation timestamps (PTS) of the video frames, effectively halving the speed of the video. The factor `2.0` increases the interval between frames, slowing down the video. Adjust this value to control the speed reduction (e.g., `4.0` for quarter speed).
- **`-an`**: Removes the audio track from the output video. Slowing the video without adjusting the audio can result in desynchronization.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter expression**: Happens if the expression passed to the `setpts` filter is incorrect. Ensure the syntax for the filter expression is correct.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## GPU Acceleration Command

While adjusting playback speed is a process handled by altering frame presentation times and does not directly benefit from GPU acceleration, encoding the video during the process can utilize GPU acceleration. For Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -filter:v "setpts=2.0*PTS" -c:v h264_nvenc -an output.mp4
```


## Additional Information

- **Audio Speed Adjustment**: If you wish to maintain the audio and adjust its speed to match the video slowdown, you can use the `atempo` audio filter. Note that `atempo` has a limited range (0.5 to 2.0), and multiple filters may be chained for greater speed reductions (e.g., `-filter:a "atempo=0.5,atempo=0.5"` for quarter speed).
- **Quality Considerations**: Slowing down a video may require re-encoding, which can affect quality. To maintain high quality, specify encoding options such as `-crf` for x264 and x265 codecs to balance between quality and file size.
- **Frame Rate Adjustment**: For significant slowdowns, the video may appear choppy due to the reduced effective frame rate. Consider interpolating frames using filters like `minterpolate` to achieve smoother slow-motion effects.
- **Complex Filter Graphs**: For advanced slowdown effects, especially when adjusting both video and audio in sync, use the `-filter_complex` option to handle both streams within a single command effectively.
