# Adding a Simple Text Overlay to a Video

Embed a text overlay onto a video file using FFmpeg's `drawtext` filter, ideal for watermarking, titles, or annotations.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Adding a Simple Text Overlay to a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    streams:
      - video_filters:
          - name: drawtext
            value: "text='Your Text Here':fontcolor=white:fontsize=24:x=10:y=10"
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```

## Command

```bash
ffmpeg -i input.mp4 -vf "drawtext=text='Your Text Here':fontcolor=white:fontsize=24:x=10:y=10" output.mp4
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-vf "drawtext=text='Your Text Here':fontcolor=white:fontsize=24:x=10:y=10"`: Applies the drawtext video filter with options to overlay text on the video:
  - `text='Your Text Here'`: The text to overlay on the video. Replace 'Your Text Here' with your desired text.
  - `fontcolor=white`: Sets the color of the text. Replace `white` with any valid color name or hex code.
  - `fontsize=24`: Sets the size of the text. Adjust `24` to your desired font size.
  - `x=10:y=10`: Positions the text 10 pixels from the left (`x`) and 10 pixels from the top (`y`) of the video frame. Adjust these values to change the text position.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Invalid filter options**: Happens if options passed to the drawtext filter are incorrect or unsupported. Verify the syntax and values for all filter options.
- **Font not found**: Arises if the specified font is not available on your system. Specify a font file directly with the `fontfile` option or ensure the font name is correct.
- **Permission denied**: Occurs if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## Additional Information

- **Customizing Text Appearance**: Beyond font size and color, the drawtext filter supports various options for customizing text appearance, including `fontfile` (to specify a custom font), `shadowcolor`, `shadowx`, `shadowy` (for text shadows), and `box` (to add a background box behind the text).
- **Dynamic Text Placement**: For more dynamic text placement, FFmpeg supports expressions for the `x` and `y` parameters, such as `x=(w-text_w)/2` to center the text horizontally.
- **Timecode Overlay**: The drawtext filter can display dynamic content like the current timecode using special variables and expressions, e.g., `text='%{pts\:hms}'` for a running timestamp.
- **Performance**: Adding a text overlay is relatively lightweight, but re-encoding the video can affect processing time, especially for videos that are high-resolution or long.
