# Reverse a Video

Create a reversed playback effect for a video file using FFmpeg, making the video play backward, which can be used for creative effects or analysis.

## VideoAlchemy Compose File

```yaml
version: 1

generate_path: "./generated"

tasks:
  - name: Reverse a Video
    command: ffmpeg
    inputs:
      - id: input_1
        source: 'input.mp4'
    codecs:
      - video_filters:
          - name: reverse
      - audio_filters:
          - name: areverse
    outputs:
      - id: output_1
        overwrite: true
        source: 'output.mp4'
```


## Command

```bash
ffmpeg -i input.mp4 -vf "reverse" -af "areverse" output.mp4
```

## Parameters

- `-i input.mp4`: Specifies the input video file. Replace `input.mp4` with the path to your source video file.
- `-vf "reverse"`: Applies the reverse filter to reverse the video frames, playing the video backward.
- `-af "areverse"`: Applies the reverse filter to the audio track, playing the audio backward in sync with the video.

## Possible Errors

- **File not found**: Occurs if FFmpeg cannot locate the input file. Ensure the path to the file is correct.
- **Memory limitations**: Reversing a video requires loading it into memory, which means extremely long or high-resolution videos might exceed system memory limits, leading to process failure.
- **Permission denied**: Arises if FFmpeg does not have the necessary permissions to read the input file or write to the output file. Check that the files and directories have the correct permissions.

## Additional Information

- **Performance Consideration**: The reverse filter is resource-intensive since it processes the entire video before outputting. Larger files are expected to have longer processing times.
- **Quality Preservation**: Reversing the video maintains the original quality. However, if adjustments to quality or file size are needed, consider using encoding parameters such as `-crf` for x264 and x265 codecs.
- **Partial Reversing**: To reverse only a portion of the video, you might use the trim, setpts, and asetpts filters to specify the segment before applying the reverse filters.
- **Compatibility**: Most media players and devices should be compatible with the reversed video. Nonetheless, testing the output is recommended to ensure it meets your specific needs, especially for use in web pages or other media.
