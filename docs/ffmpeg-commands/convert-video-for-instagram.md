# Convert Video for Instagram

Optimize and format videos for Instagram posting, ensuring compatibility with Instagram's video requirements for feed, stories, and IGTV.

## VideoAlchemy Compose File

```yaml
version: 1  # Schema version of viddo-compose

generate_path: "./generated"  # Directory of log and command files

tasks:
  - name: Convert Video for Instagram
    command: ffmpeg
    inputs:
      - id: input_9
        source: 'input.mp4'
    outputs:
      - id: output_9
        overwrite: true
        source: 'output.mp4'
    streams:
      - video_filters:
          - name: scale
            value: 1080:1920
          - name: setsar
            value: 1:1
        codec_name:
          video: libx264
          audio: aac
        crf: 23
        audio_bitrate: 128k
        preset: veryfast
        shortest: true
```

## Command for Feed and Stories

```bash
ffmpeg -i input.mp4 -vf "scale=1080:1920,setsar=1:1" -c:v libx264 -preset veryfast -crf 23 -c:a aac -b:a 128k -shortest output.mp4
```


## Command for IGTV

```bash
ffmpeg -i input.mp4 -vf "scale=720:1280,setsar=1:1" -c:v libx264 -preset veryfast -crf 23 -c:a aac -b:a 128k -shortest output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "scale=1080:1920,setsar=1:1"`** (for Feed and Stories) / **`-vf "scale=720:1280,setsar=1:1"`** (for IGTV): Sets the video filter for scaling the video to Instagram's preferred dimensions for feed and stories (1080x1920) or IGTV (720x1280) and sets the sample aspect ratio to 1:1 for square pixels.
- **`-c:v libx264`**: Uses the H.264 codec for video encoding, providing broad compatibility.
- **`-preset veryfast`**: Balances encoding speed and quality. Other presets can be used depending on CPU capability and desired output quality.
- **`-crf 23`**: Sets the Constant Rate Factor to 23, offering a balance between quality and file size. Adjust as necessary for higher or lower quality.
- **`-c:a aac`**: Sets the audio codec to AAC, ensuring broad compatibility.
- **`-b:a 128k`**: Sets the audio bit rate to 128 kbps, offering good audio quality for most purposes.
- **`-shortest`**: Ensures the output file's duration matches the shortest stream (usually the video), which is helpful when adding a static image or adjusting audio length.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid scale dimensions**: Happens if the specified scaling dimensions are not supported by Instagram or the input file. Ensure the aspect ratio and resolution fit Instagram's guidelines.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## Additional Information

- **Aspect Ratios and Dimensions**: Instagram supports various aspect ratios and resolutions depending on the content type (feed, stories, IGTV). Check Instagram's current specifications as they may update over time.
- **Quality and Compression**: Adjusting the `-crf` value can help manage file size and quality. Lower values increase quality and file size, while higher values decrease them.
- **Audio Sync**: Pay attention to audio synchronization, especially if altering video length or combining clips. Test the output file to ensure audio remains in sync.
- **Testing**: Before posting, test the output video for playback on different devices if possible, to ensure compatibility and satisfactory visual quality.
