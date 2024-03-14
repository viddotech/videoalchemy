# Cropping Videos

Crop video files to remove unwanted portions from the frame using FFmpeg.

## Command

```bash
ffmpeg -i input.mp4 -vf "crop=640:480:100:50" output.mp4
```

## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-vf "crop=640:480:100:50"`**: Applies a video filter to crop the video. The `crop` filter syntax is `crop=width:height:x:y`, where `width` and `height` are the dimensions of the output video, and `x` and `y` are the top left coordinates of the crop area in the input video. Adjust these values to your desired cropping area.
- **`output.mp4`**: Specifies the output video file. Replace `output.mp4` with your desired output file name.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid crop dimensions**: Occurs if the specified dimensions are not valid or if the crop area exceeds the boundaries of the input video. Ensure that the dimensions and coordinates are correct and within the range of the input video's size.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

Cropping with GPU acceleration can be performed using filters supported by your GPU. However, direct GPU-accelerated cropping commands vary based on the FFmpeg version and GPU. As a general approach for Nvidia GPUs, you can use:


```bash
ffmpeg -hwaccel cuda -i input.mp4 -vf "crop=640:480:100:50" -c:a copy output.mp4
```


Note: Ensure your FFmpeg build has been compiled with support for Nvidia CUDA or the relevant GPU acceleration hardware you're using. Direct GPU cropping might not be supported for all hardware, and using `-vf` with a hardware-accelerated encoder might be the most effective solution.

## Additional Information

- **Aspect Ratio**: Be mindful of the aspect ratio when cropping videos. An improper aspect ratio can lead to stretched or squished playback.
- **Quality**: Cropping a video re-encodes it, which can affect quality. To maintain quality, use a high-quality codec and consider adjusting bitrate or quality settings (e.g., using `-crf` for x264 and x265 encoders).
- **Audio**: The command above copies the audio stream without changes. If needed, audio can be re-encoded or manipulated separately.
