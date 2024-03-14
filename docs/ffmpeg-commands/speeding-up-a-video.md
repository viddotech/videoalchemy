# Speeding Up a Video

Increase the playback speed of a video file using FFmpeg, useful for creating time-lapse effects or summarizing content.

## Command

```bash
ffmpeg -i input.mp4 -filter:v "setpts=0.5*PTS" -an output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-filter:v "setpts=0.5*PTS"`**: Applies a video filter to adjust the presentation timestamps (PTS) of the video frames, effectively doubling the speed of the video. The factor `0.5` decreases the interval between frames to half, speeding up the video. Adjust this value to control the speed (e.g., `0.25` for 4x speed).
- **`-an`**: Removes the audio track from the output video. This is often necessary because changing video speed without adjusting audio can lead to desynchronization.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter expression**: Happens if the expression passed to the `setpts` filter is incorrect. Ensure the syntax for the filter expression is correct.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## GPU Acceleration Command

Speed adjustments are typically handled by altering the frame presentation times, a process that is not directly accelerated by GPUs. However, if re-encoding is required, GPU acceleration can be utilized for the encoding step. For Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -filter:v "setpts=0.5*PTS" -c:v h264_nvenc -an output.mp4
```


## Additional Information

- **Maintaining Audio**: If you wish to keep the audio and adjust its speed to match the video, you can use the `atempo` audio filter. Note that `atempo` has a limited range (0.5 to 2.0), and multiple filters may be chained for greater speed changes (e.g., `-filter:a "atempo=2.0,atempo=2.0"` for 4x speed).
- **Quality Preservation**: Speeding up video may require re-encoding, which can affect quality. Consider specifying quality-related encoding options (e.g., `-crf` for x264 and x265) to balance speed and quality.
- **Frame Rate Consideration**: Increasing the speed reduces the effective frame rate. For significant speed increases, consider adjusting the frame rate with the `-r` option to maintain smooth playback.
- **Complex Filter Graphs**: For advanced speed adjustments, especially when maintaining audio, consider using a complex filter graph with the `-filter_complex` option to manipulate both video and audio in a single command.
