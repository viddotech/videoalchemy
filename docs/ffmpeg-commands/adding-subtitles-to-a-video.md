# Adding Subtitles to a Video

Embed subtitles into a video file using FFmpeg, enhancing accessibility and viewer experience.

## Command

```bash
ffmpeg -i input.mp4 -vf subtitles=subtitles.srt output.mp4
```


## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf subtitles=subtitles.srt`**: Applies the subtitles video filter. Replace `subtitles.srt` with the path to your subtitle file.
- **`output.mp4`**: Specifies the name of the output video file. Replace `output.mp4` with your desired output file name.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input video file or subtitle file. Ensure all file paths are correct.
- **Unsupported subtitle format**: Occurs if the subtitle format is not supported by FFmpeg. Ensure the subtitle file is in a format FFmpeg can process, such as SRT or ASS.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Adding subtitles is a CPU-bound process, as it involves rendering text onto the video frames. Therefore, there is no direct GPU acceleration command for adding subtitles. However, if you're performing other video processing tasks alongside subtitle addition, you can utilize GPU acceleration for those tasks. For example, for video encoding with Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -vf subtitles=subtitles.srt -c:v h264_nvenc output.mp4
```


## Additional Information

- **Subtitle Encoding**: Ensure your subtitle file's encoding matches the expected encoding (e.g., UTF-8). Use the `sub_charenc` option if you need to specify a character encoding for the subtitles (e.g., `-sub_charenc ISO-8859-1`).
- **Styling Subtitles**: For more control over subtitle appearance, consider using the ASS (Advanced SubStation Alpha) format, which supports advanced styling options. You can convert SRT to ASS using tools like `ffmpeg` or dedicated subtitle editors.
- **Hard Subs vs. Soft Subs**: The command provided burns the subtitles into the video (hard subs), making them always visible. If you prefer to keep subtitles as a separate selectable track (soft subs), use the `-c:s copy` parameter instead of `-vf subtitles=subtitles.srt`.
- **Font Configuration**: When burning subtitles, FFmpeg uses default fonts. You can specify fonts with the `fontsdir` option if your subtitles require a specific font style or family not available by default on your system.
