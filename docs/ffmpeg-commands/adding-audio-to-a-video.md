# Adding Audio to a Video

Combine an audio track with a video file, replacing the original audio of the video using FFmpeg.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Adding Audio to a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'video.mp4'
      - id: input_2
        source: 'audio.mp3'
    codecs:
      - codec_name:
          video: copy
          audio: aac
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```

## Command

```bash
ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -strict experimental output.mp4
```


## Parameters

- **`-i video.mp4`**: Specifies the input video file. Replace `video.mp4` with the path to your source video file.
- **`-i audio.mp3`**: Specifies the input audio file. Replace `audio.mp3` with the path to your source audio file.
- **`-c:v copy`**: Copies the video codec from the input video file without re-encoding.
- **`-c:a aac`**: Specifies the audio codec for the output file. This example uses AAC, but you can replace it with your desired audio codec.
- **`-strict experimental`**: Allows the use of experimental codecs. This flag is sometimes required when using certain audio codecs like AAC.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate one of the input files. Ensure all paths to the video and audio files are correct.
- **Incompatible formats**: Occurs if the video and audio formats are not compatible with the output container. Ensure that the chosen formats are supported by the output file's container format.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Adding audio to a video primarily involves audio processing and does not significantly benefit from GPU acceleration. However, if you need to re-encode the video for any reason, you can utilize GPU acceleration for the video encoding part:

For Nvidia GPUs:

```bash
ffmpeg -i video.mp4 -i audio.mp3 -c:v h264_nvenc -c:a aac output.mp4
```


## Additional Information

- **Sync Issues**: Ensure the audio and video durations match to prevent sync issues. Use additional FFmpeg options like `-shortest` or `-t` to control the output duration if necessary.
- **Audio Quality**: When encoding audio, consider the bitrate for quality and file size. Use the `-b:a` option to specify the audio bitrate (e.g., `-b:a 192k` for 192 kbps).
- **Multiple Audio Tracks**: If the video file already contains an audio track and you want to add another one without removing the original, use the `-map` option to select multiple streams for the output file.
- **Volume Adjustment**: To adjust the volume of the added audio, use the `volume` audio filter (e.g., `-af "volume=1.5"` to increase the volume by 50%).
or adding audio to a video in FFmpeg, covering the syntax, parameters, possible errors, considerations for GPU acceleration, and additional information, all formatted in markdown.
