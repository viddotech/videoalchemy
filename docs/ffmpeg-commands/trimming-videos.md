# Trimming Videos

Trim or cut a portion of a video file to a specific start and end time using FFmpeg.

## VideoAlchemy Compose File

```yaml
tasks:
  - name: Trimming Videos
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
    codecs:
      - time_part:
          start_time: "00:00:10.000"
          end_time: "00:00:20.000"
    outputs:
      - id: trim_video
        overwrite: true
        source: 'sample/outputs/trim.mp4'
```

## Command

```bash
ffmpeg -i input.mp4 -ss 00:01:00 -to 00:02:00 -c copy output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-ss 00:01:00`**: Sets the start time for trimming. Replace `00:01:00` with your desired start time in hours:minutes:seconds format.
- **`-to 00:02:00`**: Sets the end time for trimming. Replace `00:02:00` with your desired end time in hours:minutes:seconds format. The duration of the output video will be from `-ss` to `-to`.
- **`-c copy`**: Copies the video and audio codecs from the input file without re-encoding. This option provides fast trimming but may not be frame-accurate for certain codecs.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid time duration**: Occurs if the specified start or end time is beyond the duration of the input video. Ensure the time values are within the range of the video's length.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Trimming videos typically does not benefit directly from GPU acceleration as the operation involves cutting sections of the video stream without processing the video data. However, if you need to re-encode the video, GPU acceleration can be beneficial:

For Nvidia GPUs, use:

```bash
ffmpeg -hwaccel cuda -i input.mp4 -ss 00:01:00 -to 00:02:00 -c:v h264_nvenc -c:a copy output.mp4

```


## Additional Information

- **Accuracy vs. Speed**: Using `-c copy` with `-ss` and `-to` provides fast trimming but may not be accurate for all video formats due to keyframe placement. For more accuracy, place `-ss` and `-to` before `-i` and remove `-c copy`, but be aware this will re-encode the video and take longer.
- **Re-encoding**: If you need to adjust video quality or format, remove `-c copy` and specify encoding options (e.g., `-c:v libx264 -crf 20` for x264 video codec with a constant rate factor of 20).
- **Audio Sync**: Ensure audio remains in sync with the video after trimming. In most cases, `-c copy` maintains sync, but re-encoding may sometimes be necessary if issues arise.
