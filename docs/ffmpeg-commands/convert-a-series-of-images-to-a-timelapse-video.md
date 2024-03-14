# Convert a Series of Images to a Timelapse Video

Create a timelapse video from a sequence of images using FFmpeg, perfect for showcasing changes over time or creating stunning visual sequences from still photography.

## Command

```bash
ffmpeg -framerate 24 -pattern_type glob -i 'images/*.jpg' -c:v libx264 -pix_fmt yuv420p timelapse.mp4
```

## Parameters

- `-framerate 24`: Sets the frame rate for the output video. Adjust this value based on the desired speed of the timelapse and the number of images. A higher frame rate results in a faster timelapse.
- `-pattern_type glob`: Enables the use of wildcard patterns for input file selection on systems that support it, facilitating the selection of all images in a directory.
- `-i 'images/*.jpg'`: Specifies the input images. Replace `'images/*.jpg'` with the path and pattern matching your image files. Adjust the file extension as needed for different image formats.
- `-c:v libx264`: Utilizes the H.264 codec for video encoding, offering a good compromise between quality and file size.
- `-pix_fmt yuv420p`: Ensures the pixel format is compatible with most devices and video platforms.
- `timelapse.mp4`: Names the output timelapse video file. Replace `timelapse.mp4` with your desired output file name.

## Possible Errors

- **No such file or directory**: Occurs if FFmpeg cannot find the images based on the specified pattern. Ensure the path and pattern accurately match your image files.
- **Unsupported pixel format**: May occur if the chosen pixel format is not supported by the encoder or the output device. Confirm that `yuv420p` is compatible, or try a different pixel format.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input files or write to the output file. Verify that the files and directories have the correct permissions.

## Additional Information

- **Image Naming**: For a successful timelapse, ensure your images are named sequentially (e.g., `image001.jpg`, `image002.jpg`, ...). This guarantees they are processed in the correct order.
- **Resolution and Aspect Ratio**: The resolution of the output video will be the same as that of the input images. If your images vary in size or aspect ratio, consider using additional FFmpeg filters to scale or crop them for consistency.
- **Quality Adjustments**: To control the quality of the output video, you might use the `-crf` option with the libx264 codec (e.g., `-crf 23`). Lower values yield better quality at the expense of larger file sizes.
- **Batch Processing**: This command assumes all relevant images are in a single directory and of the same file type. If your images are spread across different directories or in various formats, organize them or adjust the command as necessary.
