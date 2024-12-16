# Creating a Thumbnail from a Video

Generate a thumbnail image from a video file at a specific time frame using FFmpeg, useful for previews or video content indexing.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Creating a Thumbnail from a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    streams:
      - video_filters:
          - name: select
            value: 'eq(n\,100)'
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.png'
```

## Command

```bash
ffmpeg -i input.mp4 -ss 00:01:00 -vframes 1 output.png
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-ss 00:01:00`: Sets the timestamp for extracting the thumbnail. Replace `00:01:00` with your desired timestamp in hours:minutes:seconds format.
- `-vframes 1`: Commands FFmpeg to output one video frame, thereby creating a single image.
- `output.png`: Determines the name of the output thumbnail file. Replace `output.png` with your desired output file name and format.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Make sure the file path is correct.
- **Invalid timestamp**: Happens if the specified timestamp is beyond the video's duration. Ensure the timestamp falls within the video's length.
- **Permission denied**: Arises if FFmpeg lacks the necessary permissions to read the input file or write to the output file. Verify that the files and directories have the appropriate permissions.

## Additional Information

- **Image Format**: Different output formats (e.g., JPG, PNG) can be chosen based on your needs. PNG offers lossless compression, whereas JPEG may be preferred for its smaller file sizes or suitability for web use.
- **Choosing the Right Frame**: To select the most representative or visually appealing frame for the thumbnail, experiment with various timestamps.
- **Batch Processing**: For creating thumbnails for multiple videos, consider using a shell script (on Linux/macOS) or a batch file (on Windows) to iterate through video files and apply this FFmpeg command.
- **Quality Adjustments**: For formats like JPEG, the quality can be adjusted using the `-q:v` option (e.g., `-q:v 2` for high quality). Note that PNG uses lossless compression and does not have a quality setting.
- **Resolution Adjustment**: If a specific thumbnail size is required, incorporate the scale filter into your command (e.g., `-vf "scale=320:-1"` to set the width to 320 pixels while adjusting the height to maintain the aspect ratio).
