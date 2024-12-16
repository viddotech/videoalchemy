# Convert Video to HLS Format

Convert videos to HTTP Live Streaming (HLS) format using FFmpeg, enabling adaptive streaming over the web for various devices and network conditions.

## VideoAlchemy Compose File

```yaml
version: 1  # Schema version of viddo-compose

generate_path: "./generated"  # Directory of log and command files

tasks:
  - name: Convert Video to HLS Format
    command: ffmpeg
    inputs:
      - id: input_11
        source: 'input.mp4'
    outputs:
      - id: output_11
        overwrite: true
        source: 'output.m3u8'
    streams:
      - codec_name:
          video: copy
          audio: copy
        hls:
          time: 10
          list_size: 0
          start_number: 0
```


## Command

```bash
ffmpeg -i input.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls output.m3u8
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-codec: copy`**: Copies both the video and audio codecs from the input file without re-encoding, preserving the original quality.
- **`-start_number 0`**: Sets the starting segment number for the output HLS playlist.
- **`-hls_time 10`**: Sets the maximum duration of each segment (in seconds). `10` seconds is a common choice for balancing file size and playback smoothness.
- **`-hls_list_size 0`**: Specifies the maximum number of playlist entries. Setting it to `0` includes all segments in the playlist, allowing for unlimited playback duration.
- **`-f hls`**: Sets the format to HLS for the output file.
- **`output.m3u8`**: Specifies the name of the output HLS playlist file. Replace `output.m3u8` with your desired output file name.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Incompatible codec for HLS**: May occur if the input file's codec is not compatible with HLS. Consider re-encoding with `-codec:v libx264 -codec:a aac` for wider compatibility.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to write the output files. Check that the destination directory has the correct permissions.

## Additional Information

- **Re-encoding for Compatibility**: To ensure compatibility across all devices, you may need to re-encode your video to H.264/AAC. Use `-codec:v libx264 -codec:a aac` before the output file name.
- **Segment Duration**: Adjusting `-hls_time` affects the trade-off between download efficiency and the ability to adjust to changing network conditions. Shorter segments offer faster adaptation but can increase overhead.
- **Encryption**: For content protection, HLS supports AES-128 encryption. Use `-hls_key_info_file` followed by a path to a key info file to enable encryption.
- **Bandwidth Adaptation**: To create a multi-bitrate HLS stream, you will need to encode your input video at various bitrates and resolutions and then generate an HLS master playlist linking to the individual variant playlists.
- **Audio-Only HLS**: For audio-only streaming, you can omit video-related parameters and codecs, focusing on optimizing audio quality and compatibility.

This markdown text provides a structured and detailed overview of the command for converting video to HLS format using FFmpeg, covering the syntax, parameters, possible errors, and additional information, all formatted in markdown.
