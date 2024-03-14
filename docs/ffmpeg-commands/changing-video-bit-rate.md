# Changing Video Bit Rate

Adjust the bit rate of a video file using FFmpeg to change its size and quality, useful for optimizing videos for various bandwidths or storage constraints.

## Command

```bash
ffmpeg -i input.mp4 -b:v 1M output.mp4
```

## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-b:v 1M`**: Sets the video bit rate to 1 Mbps. Replace `1M` with your desired bit rate (e.g., `500k` for 500 kbps).

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid bit rate value**: Occurs if the specified bit rate value is not recognized or is inappropriate for the video resolution or format. Ensure the bit rate value is specified in a valid format (e.g., `k` for kbps, `M` for Mbps).
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Changing the video bit rate usually involves re-encoding, which can benefit from GPU acceleration. Here’s how you can do it for Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -c:v h264_nvenc -b:v 1M output.mp4
```

Replace `h264_nvenc` with your GPU's specific hardware encoder (e.g., `hevc_nvenc` for HEVC) and adjust the bit rate as desired.

## Additional Information

- **Quality vs. Size**: A higher bit rate generally results in better video quality but larger file sizes. Conversely, reducing the bit rate can significantly decrease file size at the cost of quality.
- **Audio Bit Rate**: To change the audio bit rate, use the `-b:a` option followed by the desired audio bit rate (e.g., `-b:a 128k` for 128 kbps). This can further reduce file size or adjust audio quality.
- **Two-Pass Encoding**: For more efficient bit rate utilization and consistent quality, consider using two-pass encoding. This process analyzes the video in the first pass and then adjusts the encoding to achieve the target bit rate more accurately in the second pass.
- **Compatibility**: Ensure the chosen video and audio bit rates are compatible with your intended playback devices or platforms, especially if you are optimizing for streaming or mobile devices.
