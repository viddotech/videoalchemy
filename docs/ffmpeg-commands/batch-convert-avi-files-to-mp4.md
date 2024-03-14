# Batch Convert .AVI Files to .MP4

Automate the conversion of multiple .AVI video files to the .MP4 format using FFmpeg, streamlining the process for efficiency and consistency.

## Command (Linux and macOS)

```bash
for i in *.avi; do ffmpeg -i "$i" -c:v libx264 -preset fast -c:a aac "${i%.avi}.mp4"; done
```
## Command (Windows Command Prompt)

```bash
FOR %G IN (*.avi) DO ffmpeg -i "%G" -c:v libx264 -preset fast -c:a aac "%~nG.mp4"
```

## Parameters

- `-i "$i" / "%G"`: Specifies the input .AVI file. The variables `$i` (for bash) and `%G` (for CMD) represent each .AVI file found by the loop.
- `-c:v libx264`: Utilizes the H.264 codec for video encoding to ensure good compatibility and a balance between quality and file size.
- `-preset fast`: Selects a preset for the libx264 encoder that optimizes the balance between encoding speed and output quality. Alternatives include `veryfast`, `faster`, `medium`, `slow`, and `veryslow`.
- `-c:a aac`: Sets the audio codec to AAC for broad compatibility.
- `${i%.avi}.mp4" / "%~nG.mp4`: Defines the name of the output .MP4 file. The expressions `${i%.avi}.mp4` and `%~nG.mp4` remove the `.avi` extension from the input file name and append `.mp4`.

## Possible Errors

- **File not found**: This error may occur if no .AVI files are present in the directory. Ensure you are in the correct directory that contains .AVI files.
- **Codec compatibility issues**: If the .AVI files use codecs that are incompatible for conversion to H.264/AAC, it may be necessary to adjust codec parameters or investigate the codecs used in the source files.
- **Permission denied**: This issue can arise if you lack the necessary permissions to read the .AVI files or write the .MP4 files. Verify the permissions for both files and directories.

## Additional Information

- **Directory Navigation**: Before executing the batch conversion command, navigate to the directory containing your .AVI files. Use `cd /path/to/directory` on Linux/macOS or `CD \path\to\directory` on Windows.
- **Quality Adjustment**: The quality of the output video can be modified by altering the `-crf` value for the libx264 encoder (e.g., `-crf 23`). Lower values enhance quality and increase file size, whereas higher values reduce both.
- **Audio Quality**: To improve audio quality, adjust the bit rate using the `-b:a` option (e.g., `-b:a 192k`).
- **Scripting**: For convenience, consider saving these commands in a script file (e.g., `convert.sh



