# Converting Audio Files

Convert audio files from one format to another using FFmpeg.

## VideoAlchemy Compose File
```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Converting Audio Files
    command: ffmpeg
    inputs:
      - id: input_audio
        source: 'input.wav'
    outputs:
      - id: output_audio
        overwrite: true
        source: 'output.mp3'
```

## Command

```bash
ffmpeg -i input.wav output.mp3
```

## Parameters

- **`-i input.wav`**: Specifies the input audio file. Replace `input.wav` with the path to your source audio file.
- **`output.mp3`**: Specifies the output audio file. The extension of the output file determines the output format. Replace `output.mp3` with your desired output file name and format.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Unsupported codec**: Occurs if the specified audio codec is not supported or not installed. Ensure that the desired audio format is supported.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Audio conversion typically does not benefit from GPU acceleration as it is primarily an audio processing task. However, if you are dealing with an audio file that requires decoding with GPU acceleration, you can use:

For Nvidia GPUs:

```bash
ffmpeg -hwaccel cuda -i input.wav -c:a libmp3lame output.mp3
```


## Additional Information

- **Audio Formats**: Common audio formats include MP3, AAC, WAV, and FLAC. The format is determined by the file extension of the output file.
- **Re-encoding**: If you need to re-encode the audio into a different format or adjust the bitrate, specify the desired audio codec with `-c:a` and set the bitrate with `-b:a` (e.g., `-c:a libmp3lame -b:a 192k` for MP3 at 192 kbps).
- **Quality**: When re-encoding, consider the trade-off between file size and audio quality. Higher bitrates generally result in better quality but larger file sizes.
- **Metadata**: To preserve metadata (e.g., artist, album, title) during conversion, use the `-map_metadata` option (e.g., `-map_metadata 0` to copy metadata from the first input file to the output file).


