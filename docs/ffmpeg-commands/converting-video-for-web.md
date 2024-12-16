# Converting Video for Web

Optimize and convert videos for web playback, ensuring compatibility and efficient loading across various browsers and devices using FFmpeg.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Converting Video for Web
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    streams:
      - codec_name:
          video: libx264
          audio: aac
        video_bitrate: 1M
        audio_bitrate: 128k
        preset: medium
        move_flags: faststart
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```


## Command

```bash
ffmpeg -i input.mp4 -c:v libx264 -preset medium -b:v 1M -c:a aac -b:a 128k -movflags +faststart output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-c:v libx264`**: Sets the video codec to H.264, widely supported across web browsers and devices.
- **`-preset medium`**: Balances encoding speed and output quality. Other presets include `fast`, `slow`, and `veryfast`.
- **`-b:v 1M`**: Sets the video bit rate to 1 Mbps. Adjust this rate based on your quality and bandwidth requirements.
- **`-c:a aac`**: Sets the audio codec to AAC, ensuring broad compatibility.
- **`-b:a 128k`**: Sets the audio bit rate to 128 kbps, offering a good balance between quality and file size.
- **`-movflags +faststart`**: Moves some data to the beginning of your file, facilitating quicker playback start times when streamed online.
- **`output.mp4`**: Specifies the name of the output video file. Replace `output.mp4` with your desired output file name.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Unsupported codec**: May occur if the selected codecs are not compatible with the input file's format. Ensure that the input file format supports the chosen codecs.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## GPU Acceleration Command

For faster encoding, especially useful for high-resolution videos, you can use GPU acceleration. For Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -c:v h264_nvenc -preset:v hq -b:v 1M -c:a aac -b:a 128k -movflags +faststart output.mp4
```

## Additional Information

- **Web Compatibility**: The H.264 video codec and AAC audio codec combination offers broad compatibility with web browsers and mobile devices.
- **Quality and File Size**: Adjust the `-b:v` and `-b:a` parameters to find the right balance between video quality and file size. Lower bit rates result in smaller files but might reduce quality.
- **Preset Options**: The `-preset` parameter affects encoding speed and output file size; a slower preset offers better compression (smaller file size) at the cost of increased encoding time.
- **Fast Start**: The `-movflags +faststart` flag is crucial for web videos, allowing playback to begin before the file is fully downloaded, enhancing the user experience on streaming platforms.

