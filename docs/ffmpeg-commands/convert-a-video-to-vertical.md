# Convert a Video to Vertical

Transform a horizontal (landscape) video into a vertical (portrait) format using FFmpeg, making it suitable for platforms that prefer vertical content, such as social media stories or mobile viewing.

## Command

```bash
ffmpeg -i input.mp4 -vf "pad=ih*9/16:ih:(ow-iw)/2:0" -c:a copy output.mp4
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-vf "pad=ih*9/16:ih:(ow-iw)/2:0"`: Applies the pad filter to adjust the video frame, creating a vertical video with a 9:16 aspect ratio. This command adds padding to the sides of the original video to fit the vertical aspect ratio, centering the video horizontally.
- `-c:a copy`: Copies the original audio track without any changes.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the file path is correct.
- **Invalid filter settings**: Happens if the filter syntax is incorrect or results in an unsupported configuration. Ensure the formula for the padding and aspect ratio is applied correctly.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## Additional Information

- **Aspect Ratio Considerations**: The provided command converts the video to a 9:16 aspect ratio, which is typical for vertical videos. Adjust the formula within the pad filter for different aspect ratios as needed.
- **Quality and Resolution**: Upscaling or padding the video can impact visual quality. Take into account the source video's resolution and the desired output resolution to ensure satisfactory quality.
- **Black Bars**: The padding will default to black bars. You can change the color by adding `:color=white` (or any other color) to the pad filter arguments.
- **Audio Sync**: The audio will remain unchanged and in sync with the video. If the video's duration changes as a result of the conversion process, verify that the audio still aligns as expected.
- **Advanced Editing**: For more complex transformations or to add background images instead of simple color padding, you may want to explore more advanced FFmpeg filters or editing software that offers greater control over the video layout.
