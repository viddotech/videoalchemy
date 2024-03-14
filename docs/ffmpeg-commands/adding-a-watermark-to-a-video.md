# Adding a Watermark to a Video

Embed a watermark image onto a video file using FFmpeg to brand or copyright protect your video content.

## Command

```bash
ffmpeg -i input.mp4 -i watermark.png -filter_complex "overlay=10:10" output.mp4
```

## Parameters

- **`-i input.mp4`**: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- **`-i watermark.png`**: Specifies the watermark image file. Replace `watermark.png` with the path to your watermark image. The image format can be PNG, JPG, or any other image format supported by FFmpeg.
- **`-filter_complex "overlay=10:10"`**: Applies a complex filter to overlay the watermark image onto the video. `10:10` positions the top-left corner of the watermark image 10 pixels from the top and 10 pixels from the left of the video frame. Adjust these values to change the watermark position.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input video file or the watermark image file. Ensure all file paths are correct.
- **Invalid overlay position**: Occurs if the specified position values for the overlay filter are outside the bounds of the video frame. Ensure the position values are within the video dimensions.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

While adding a watermark is primarily a CPU-bound process, you can utilize GPU acceleration for the video encoding part of the process. For Nvidia GPUs:

```bash
ffmpeg -i input.mp4 -i watermark.png -filter_complex "overlay=10:10" -c:v h264_nvenc output.mp4
```


## Additional Information

- **Transparency Support**: If your watermark is a PNG with transparency, FFmpeg will respect this transparency in the output video.
- **Watermark Size**: If necessary, you can resize the watermark image before overlaying it onto the video using the `scale` filter within the `-filter_complex` option (e.g., `overlay=10:10,scale=100:50` to resize the watermark to 100x50 pixels).
- **Positioning the Watermark**: The `overlay` filter allows for dynamic positioning using variables such as `main_w`, `main_h` (video width and height), and `overlay_w`, `overlay_h` (watermark width and height). For example, `overlay=(main_w-overlay_w-10):(main_h-overlay_h-10)` positions the watermark 10 pixels from the bottom right corner.
- **Fade In/Out Effects**: To add fade-in and fade-out effects to your watermark, you can use the `fade` filter in combination with the `overlay` filter within the `-filter_complex` option.
