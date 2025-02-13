# Guides

#### Introduction

This guide will help you understand how to use the `viddo-compose.yaml` file to define and execute video processing tasks using FFmpeg. The `viddo-compose.yaml` file is a configuration file that specifies the commands, inputs, outputs, and codecs/filters to apply for each task.

#### Table of Contents

1. [Version](#version)
2. [Generate Path](#generate-path)
3. [Tasks](#tasks)
    - [Basic Video Conversion](#basic-video-conversion)
    - [Extracting Audio from Video](#extracting-audio-from-video)
    - [Resizing Video](#resizing-video)
    - [Trimming Videos](#trimming-videos)
    - [Combining Videos](#combining-videos)
    - [Extract Images from Video](#extract-images-from-video)
    - [Creating a Video from Images](#creating-a-video-from-images)
    - [Convert Video to Multi-Bitrate HLS Format](#convert-video-to-multi-bitrate-hls-format)

#### Version

Compose file version is different from the VideoAlchemy version. it's used to define the schema version of the `viddo-compose.yaml` file.

```yaml
version: 1
```

#### Generate Path

The path where containing the generated ffmpeg script, ffmpeg logs, etc.

```yaml
generate_path: "./sample/generated"
```

#### Tasks

The `tasks` attribute is a list of tasks to be executed. Each task specifies the command to run, the inputs, outputs, and any codecs or filters to apply.

##### Basic Video Conversion

```yaml
- name: Basic Video Conversion  # Name of the task
  command: ffmpeg  # Command to be executed
  inputs:  # List of input files
    - id: input_1  # Identifier for the input
      source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'  # Path to the input file
  streams:
    - stream_from:
        input_id: input_1
      stream_to:
        output_id: output_avi
  outputs:  # List of output files
    - id: output_avi  # Identifier for the output
      overwrite: true  # Whether to overwrite the output file if it exists
      source: 'sample/outputs/1280x720_30mb_output.avi'  # Path to the output file
```

##### Extracting Audio from Video

```yaml
- name: Extracting Audio from Video
  command: ffmpeg
  inputs:
    - id: input_2
      output_id: output_avi  # Reference to the output of a previous task
  streams:  # List of streams to apply
    - codec_name:
        audio: copy  # Copy the audio codec
      video_none: true  # No video codec
      stream_from:
        input_id: input_2
      stream_to:
        output_id: output_only_audio
  run_after:  # List of tasks to run before this task
    - Basic Video Conversion
  outputs:
    - id: output_only_audio
      overwrite: true
      source: "sample/outputs/1280x720_30mb_output_only_audio.mp3"
```

##### Resizing Video

```yaml
- name: Resizing Video
  command: ffmpeg
  inputs:
    - id: input_3
      source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
  streams:
    - video_filters:  # List of video filters to apply
        - name: scale
          value: "720:480:flags=lanczos"
      stream_from:
        input_id: input_3
      stream_to:
        output_id: resized_output.mp4
  outputs:
    - id: resized_output.mp4
      overwrite: true
      source: 'sample/outputs/resized.mp4'
```

##### Trimming Videos

```yaml
- name: Trimming Videos
  command: ffmpeg
  inputs:
    - id: input_4
      source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
  streams:
    - time_part:  # Time range to trim
        start_time: "00:00:10.000"
        end_time: "00:00:20.000"
      stream_from:
        input_id: input_4
      stream_to:
        output_id: trim_video
  outputs:
    - id: trim_video
      overwrite: true
      source: 'sample/outputs/trim.mp4'
```

##### Extract Images from Video

```yaml
- name: Extract images from video
  command: ffmpeg
  inputs:
    - id: input_5
      source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'
  streams:
    - video_filters:
      - name: fps
        value: 1
      stream_from: 
        input_id: input_5
      stream_to:
        output_id: extract_image_%04d
  outputs:
    - id: extract_image_%04d
      start_number: 0
      length: 10
      overwrite: true
      source: 'sample/outputs/extracted/%04d.png'
```

##### Creating a Video from Images

```yaml
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
      stream_from:
        input_id: input_6
      stream_to:
        output_id: video_from_images
  run_after:
    - Extract images from video
  outputs:
    - id: video_from_images
      overwrite: true
      source: 'sample/outputs/video_from_images.mp4'
```

##### Convert Video to Multi-Bitrate HLS Format

```yaml
  - name: Convert video to multi-bitrate HLS format
    command: ffmpeg
    inputs:
       - id: input_7
         source: 'sample/inputs/sample-10s.mp4'
    streams:
       - stream_from:
            input_id: input_7
            stream_type: audio
         stream_name: "my-audio"
         codec_name:
            audio: aac
         constant_bitrate:
            audio: 128k
         audio_sampling_rate: 48000
         stream_to:
            output_id: playlist
            groups: [playlist_1, playlist_2, playlist_3]
       - stream_from:
            input_id: input_7
            stream_type: video
         stream_name: video_1
         inject_streams: ["my-audio"]
         video_filters:
            - name: scale
              value: "-2:720"
         constant_bitrate:
            video: 3000k
         max_rate:
            video: 3200k
         buffer_size: 6000k
         stream_to:
            output_id: playlist
            groups: playlist_1
       - stream_from:
            input_id: input_7
            stream_type: video
         stream_name: video_2
         inject_streams: ["my-audio"]
         video_filters:
            - name: scale
              value: "-2:480"
         constant_bitrate:
            video: 1500k
         max_rate:
            video: 1600k
         buffer_size: 3000k
         stream_to:
            output_id: playlist
            groups: playlist_2
       - stream_from:
            input_id: input_7
            stream_type: video
         stream_name: video_3
         inject_streams: ["my-audio"]
         video_filters:
            - name: scale
              value: "-2:360"
         constant_bitrate:
            video: 800k
         max_rate:
            video: 900k
         buffer_size: 1800k
         stream_to:
            output_id: playlist
            groups: playlist_3
    outputs:
       - id: playlist
         overwrite: true
         format: hls
         source: 'sample/outputs/1.0/hls/%v/playlist.m3u8'
         hls:
            time: 10
            segment_filename: "sample/outputs/1.0/hls/%v/segment_%03d.ts"
            master_playlist_name: "master.m3u8"
            playlist_type: "vod"
```