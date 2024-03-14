# Creating a Video from Images

Create a video file from a sequence of images using FFmpeg, ideal for timelapse videos or animations.

## Command

```bash
ffmpeg -framerate 24 -i input%d.jpg -c:v libx264 -pix_fmt yuv420p output.mp4
```


## Parameters

- **`-framerate 24`**: Specifies the frame rate of the output video. Replace `24` with your desired frame rate.
- **`-i input%d.jpg`**: Specifies the input file pattern. `%d` acts as a placeholder for the sequence numbers of your images (e.g., `input1.jpg`, `input2.jpg`, ...). Replace `input` with the base name of your image files.
- **`-c:v libx264`**: Sets the video codec for the output file to H.264, which is widely supported across devices and platforms.
- **`-pix_fmt yuv420p`**: Sets the pixel format to `yuv420p`, ensuring compatibility with most devices and media players.
- **`output.mp4`**: Specifies the name of the output video file. Replace `output.mp4` with your desired output file name and format.

## Possible Errors

- **File sequence mismatch**: Occurs if FFmpeg cannot find a continuous sequence of image files matching the specified pattern. Ensure your image files are named sequentially without gaps.
- **Unsupported codec or format**: Occurs if the specified codec or pixel format is not supported or incorrectly specified. Verify that you have the correct codec and pixel format for your output file.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Ensure that the files and directories have the correct permissions.

## GPU Acceleration Command

While encoding to video, you can utilize GPU acceleration to speed up the process. Here’s how you can do it for Nvidia GPUs:

```bash
ffmpeg -framerate 24 -i input%d.jpg -c:v h264_nvenc -pix_fmt yuv420p output.mp4
```


Replace `h264_nvenc` with your GPU's specific hardware encoder (e.g., `hevc_nvenc` for HEVC).

## Additional Information

- **Image Formats**: You can use various image formats as input (e.g., PNG, JPEG). The choice of format can affect the quality and size of your output video.
- **Frame Rate Adjustment**: Adjusting the frame rate (`-framerate`) changes the playback speed of the resulting video. A higher frame rate results in smoother video, while a lower frame rate can create a fast-motion effect.
- **Quality Control**: To control the quality of the output video, use the `-crf` option with `-c:v libx264`. Lower values produce better quality at the expense of larger file sizes (e.g., `-crf 18` for high quality).
- **Resolution Adjustment**: If your images are of different sizes or if you wish to change the video resolution, you can use the `scale` video filter (e.g., `-vf "scale=1920:1080"` to set the video resolution to 1920x1080).
