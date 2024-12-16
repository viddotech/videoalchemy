## VideoAlchemy Tutorial

### Run VideoAlchemy

```bash
videoalchemy compose -f viddo-compose.yaml
```

The `viddo-compose.yaml` file is a configuration file used to define a series of video processing tasks using FFmpeg. Each task specifies the command to run, the inputs, outputs, and any codecs or filters to apply. Below is a detailed explanation of the attributes used in the `viddo-compose.yaml` file:



### Example `viddo-compose.yaml` File

```yaml
version: 0.1  # The version of the viddo-compose schema

generate_path: "./sample/generated"  # The path where generated files will be stored

tasks:  # List of tasks to be executed
  - name: Basic Video Conversion  # Name of the task
    command: ffmpeg  # Command to be executed
    inputs:  # List of input files
      - id: input_1  # Identifier for the input
        source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'  # Path to the input file
    outputs:  # List of output files
      - id: output_avi  # Identifier for the output
        overwrite: true  # Whether to overwrite the output file if it exists
        source: 'sample/outputs/1280x720_30mb_output.avi'  # Path to the output file

  - name: Extracting Audio from Video
    command: ffmpeg
    inputs:
      - id: input_2
        output_id: output_avi  # Reference to the output of a previous task
    streams:  # List of streams to apply
      - codec_name:
          audio: copy  # Copy the audio codec
        video_none: true  # No video codec
    run_after:  # List of tasks to run before this task
      - Basic Video Conversion
    outputs:
      - id: output_only_audio
        overwrite: true
        source: "sample/outputs/1280x720_30mb_output_only_audio.mp3"

  - name: Resizing Video
    command: ffmpeg
    inputs:
      - id: input_3
        source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
    streams:
      - video_filters:  # List of video filters to apply
          - name: scale
            value: "720:480:flags=lanczos"
    outputs:
      - id: resized_output.mp4
        overwrite: true
        source: 'sample/outputs/resized.mp4'

  - name: Trimming Videos
    command: ffmpeg
    inputs:
      - id: input_4
        source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
    streams:
      - time_part:  # Time range to trim
          start_time: "00:00:10.000"
          end_time: "00:00:20.000"
    outputs:
      - id: trim_video
        overwrite: true
        source: 'sample/outputs/trim.mp4'

  - name: Combining Videos
    command: ffmpeg
    streams:
      - concat_files:  # List of files to concatenate
          - source: "./sample/inputs/sample-10s.mp4"
          - source: "./sample/inputs/sample-20s.mp4"
    outputs:
      - id: concat_video
        overwrite: true
        source: 'sample/outputs/concat.mp4'

  - name: Extract images from video
    command: ffmpeg
    inputs:
      - id: input_5
        source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
    streams:
      - video_filters:
        - name: fps
          value: 1
    outputs:
      - id: extract_image_%04d
        start_number: 0
        length: 10
        overwrite: true
        source: 'sample/outputs/extracted/%04d.png'

  - name: Creating a video from images
    command: ffmpeg
    inputs:
      - id: input_6
        output_id: extract_image_%04d
    streams:
      - input_framerate: 24
        codec_name:
          video: libx264
        pixel_format: yuv420p
    run_after:
      - Extract images from video
    outputs:
      - id: video_from_images
        overwrite: true
        source: 'sample/outputs/video_from_images.mp4'

  - name: Convert video to multi-bitrate HLS format
    command: ffmpeg
    inputs:
      - id: input_7
        source: 'sample/inputs/sample-10s.mp4'
    streams:
      - codec_name:
          audio: aac
        constant_bitrate:
          audio: 128k
        audio_sampling_rate: 48000
      - video_filters:
          - name: scale
            value: "-2:720"
        constant_bitrate:
          video: 3000k
        max_rate: 3200k
        buffer_size: 6000k
        map_input: true
      - video_filters:
          - name: scale
            value: "-2:480"
        constant_bitrate:
          video: 1500k
        max_rate: 1600k
        buffer_size: 3000k
        output_id: hls_video_480p
        map_input: true
      - video_filters:
          - name: scale
            value: "-2:360"
        constant_bitrate:
          video: 800k
        max_rate: 900k
        buffer_size: 1800k
        output_id: hls_video_360p
        map_input: true
      - hls:
          time: 10
          segment_filename: "sample/outputs/hls/%v/segment_%03d.ts"
          master_playlist_name: "master.m3u8"
          playlist_type: "vod"
        variant_stream_map: "v:0,a:0 v:1,a:1 v:2,a:2"
    outputs:
      - id: playlist
        overwrite: true
        format: hls
        source: 'sample/outputs/hls/%v/playlist.m3u8'
```

### Full `viddo-compose.yaml` File with Comments

```yaml
version: 0.1  # The version of the viddo-compose schema

generate_path: "./sample/generated"  # The path where generated files will be stored

tasks:  # List of tasks to be executed
  - name: Task Name  # Name of the task
    command: ffmpeg  # Command to be executed
    inputs:  # List of input files
      - id: input_id  # Identifier for the input
        source: 'path/to/input/file'  # Path to the input file
        output_id: previous_task_output_id  # Reference to the output of a previous task (optional)
    streams:  # List of streams to apply (optional)
      - codec_name:
          video: libx264  # Video codec
          audio: aac  # Audio codec
        video_filters:  # List of video filters to apply (optional)
          - name: scale
            value: "1280:720"
        audio_filters:  # List of audio filters to apply (optional)
          - name: volume
            value: "2.0"
        constant_bitrate:  # Constant bitrate settings (optional)
          video: 3000k
          audio: 128k
        max_rate: 3200k  # Maximum bitrate (optional)
        buffer_size: 6000k  # Buffer size (optional)
        pixel_format: yuv420p  # Pixel format (optional)
        time_part:  # Time range to trim (optional)
          start_time: "00:00:10.000"
          end_time: "00:00:20.000"
        concat_files:  # List of files to concatenate (optional)
          - source: 'path/to/file1'
          - source: 'path/to/file2'
        hls:  # HLS settings (optional)
          time: 10
          segment_filename: "path/to/segment_%03d.ts"
          master_playlist_name: "master.m3u8"
          playlist_type: "vod"
        variant_stream_map: "v:0,a:0 v:1,a:1 v:2,a:2"
    run_after:  # List of tasks to run before this task (optional)
      - Previous Task Name
    outputs:  # List of output files
      - id: output_id  # Identifier for the output
        overwrite: true  # Whether to overwrite the output file if it exists
        source: 'path/to/output/file'  # Path to the output file
        format: hls  # Output format (optional)
        start_number: 0  # Start number for output files (optional)
        length: 10  # Number of output files (optional)
```
