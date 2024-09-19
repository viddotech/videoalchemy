# Compose File References

## Version

The `version` attribute specifies the version of the `viddo-compose` schema being used.

```yaml
version: 0.1  # The version of the viddo-compose schema
```

## Generate Path

The `generate_path` attribute specifies the path where generated files will be stored.

```yaml
generate_path: "./sample/generated"  # The path where generated files will be stored
```

## Tasks

The `tasks` attribute is a list of tasks to be executed. Each task specifies the command to run, the inputs, outputs, and any codecs or filters to apply.

### Task Parameters

#### Name

The `name` attribute specifies the name of the task.

```yaml
name: My Task  # Name of the task
```

#### Command

The `command` attribute specifies the command to be executed. For FFmpeg, this is typically `ffmpeg`.

```yaml
command: ffmpeg  # Command to be executed
```

#### Inputs

The `inputs` attribute is a list of input files for the task.

```yaml
inputs:
  - id: input_1  # Identifier for the input
    source: 'sample/inputs/SampleVideo_1280x720_30mb.mp4'  # Path to the input file
    output_id: output_1  # Reference to another output as an input
    realtime: false  # Whether the input is real-time
    format: mp4  # Format of the input file
```

- `id`: Identifier for the input. Any non-empty string.
- `source`: Path to the input file. Corresponds to the `-i` parameter in FFmpeg. Any valid file, rtmp, rtsp, etc.
- `output_id`: Reference to another output as an input. Any valid output ID from other tasks that listed in `run_after` attribute.
- `realtime`: Whether the input is real-time. values are `true` or `false`.
- `format`: Format of the input file. available values are:

`3dostr`, `4xm`, `aa`, `aac`, `aax`, `acm`, `act`, `adf`, `adp`, `ads`, `aea`, `afc`, `aix`, `alias_pix`, `amrnb`, `amrwb`, `anm`, `apc`, `ape`, `aqtitle`, `argo_brp`, `asf_o`, `au`, `av1`, `avi`, `avr`, `avs`, `bethsoftvid`, `bfi`, `bfstm`, `bin`, `bink`, `binka`, `brstm`, `c93`, `cdg`, `cdxl`, `cine`, `concat`, `cri_pipe`, `dcstr`, `dds_pipe`, `derf`, `dfa`, `dhav`, `dpx_pipe`, `dsf`, `dsicin`, `dss`, `dvbsub`, `dvbtxt`, `dxa`, `ea`, `ea_cdata`, `epaf`, `exr_pipe`, `flic`, `frm`, `fsb`, `fwse`, `g729`, `gdv`, `gem_pipe`, `genh`, `gif_pipe`, `hca`, `hcom`, `hnm`, `idcin`, `idf`, `iff`, `ifv`, `imf`, `ingenient`, `ipmovie`, `ipu`, `iss`, `iv8`, `ivr`, `j2k_pipe`, `jpeg_pipe`, `jpegls_pipe`, `jpegxl_pipe`, `jv`, `kux`, `lavfi`, `libgme`, `live_flv`, `lmlm4`, `loas`, `luodat`, `lvf`, `lxf`, `mca`, `mcc`, `mgsts`, `mjpeg_2000`, `mlv`, `mm`, `mods`, `moflex`, `mov`, `mp4`, `m4a`, `3gp`, `3g2`, `mj2`, `mpc`, `mpc8`, `mpegtsraw`, `mpegvideo`, `mpl2`, `mpsub`, `msf`, `msnwctcp`, `msp`, `mtaf`, `mtv`, `musx`, `mv`, `mvi`, `mxg`, `nc`, `nistsphere`, `nsp`, `nsv`, `paf`, `pam_pipe`, `pbm_pipe`, `pcx_pipe`, `pfm_pipe`, `pgm_pipe`, `pgmyuv_pipe`, `pgx_pipe`, `phm_pipe`, `photocd_pipe`, `pictor_pipe`, `pjs`, `pmp`, `png_pipe`, `pp_bnk`, `ppm_pipe`, `psd_pipe`, `psxstr`, `pva`, `pvf`, `qcp`, `qdraw_pipe`, `qoi_pipe`, `r3d`, `realtext`, `redspark`, `rl2`, `rpl`, `rsd`, `sami`, `sbg`, `scd`, `sdp`, `sdr2`, `sds`, `sdx`, `ser`, `sga`, `sgi_pipe`, `shn`, `siff`, `simbiosis_imx`, `sln`, `smk`, `smush`, `sol`, `stl`, `subviewer`, `subviewer1`, `sunrast_pipe`, `svag`, `svg_pipe`, `svs`, `tak`, `tedcaptions`, `thp`, `tiertexseq`, `tiff_pipe`, `tmv`, `tty`, `txd`, `ty`, `v210`, `v210x`, `vag`, `vbn_pipe`, `vividas`, `vivo`, `vmd`, `vobsub`, `vpk`, `vplayer`, `vqf`, `wc3movie`, `webp_pipe`, `wsd`, `wsvqa`, `wve`, `x11grab`, `xa`, `xbin`, `xmv`, `xpm_pipe`, `xvag`, `xwd_pipe`, `xwma`, `yop`, `video4linux2`, `v4l2`


#### Outputs

The `outputs` attribute is a list of output files for the task.

```yaml
outputs:
  - id: output_1  # Identifier for the output
    overwrite: true  # Whether to overwrite the output file if it exists
    source: 'sample/outputs/output.mp4'  # Path to the output file
    format: mp4  # Format of the output file
    start_number: 0  # Starting number for the output file
    length: 10  # Length of the output file sequence
```

- `id`: Identifier for the output. Any non-empty string.
- `overwrite`: Whether to overwrite the output file if it exists. values are `true` or `false`.
- `source`: Path to the output file. Any valid file path. Any valid file, rtmp, rtsp, etc.
- `format`: Format of the output file. available values are:

`3dostr`, `4xm`, `aa`, `aac`, `aax`, `acm`, `act`, `adf`, `adp`, `ads`, `aea`, `afc`, `aix`, `alias_pix`, `amrnb`, `amrwb`, `anm`, `apc`, `ape`, `aqtitle`, `argo_brp`, `asf_o`, `au`, `av1`, `avi`, `avr`, `avs`, `bethsoftvid`, `bfi`, `bfstm`, `bin`, `bink`, `binka`, `brstm`, `c93`, `cdg`, `cdxl`, `cine`, `concat`, `cri_pipe`, `dcstr`, `dds_pipe`, `derf`, `dfa`, `dhav`, `dpx_pipe`, `dsf`, `dsicin`, `dss`, `dvbsub`, `dvbtxt`, `dxa`, `ea`, `ea_cdata`, `epaf`, `exr_pipe`, `flic`, `frm`, `fsb`, `fwse`, `g729`, `gdv`, `gem_pipe`, `genh`, `gif_pipe`, `hca`, `hcom`, `hnm`, `idcin`, `idf`, `iff`, `ifv`, `imf`, `ingenient`, `ipmovie`, `ipu`, `iss`, `iv8`, `ivr`, `j2k_pipe`, `jpeg_pipe`, `jpegls_pipe`, `jpegxl_pipe`, `jv`, `kux`, `lavfi`, `libgme`, `live_flv`, `lmlm4`, `loas`, `luodat`, `lvf`, `lxf`, `mca`, `mcc`, `mgsts`, `mjpeg_2000`, `mlv`, `mm`, `mods`, `moflex`, `mov`, `mp4`, `m4a`, `3gp`, `3g2`, `mj2`, `mpc`, `mpc8`, `mpegtsraw`, `mpegvideo`, `mpl2`, `mpsub`, `msf`, `msnwctcp`, `msp`, `mtaf`, `mtv`, `musx`, `mv`, `mvi`, `mxg`, `nc`, `nistsphere`, `nsp`, `nsv`, `paf`, `pam_pipe`, `pbm_pipe`, `pcx_pipe`, `pfm_pipe`, `pgm_pipe`, `pgmyuv_pipe`, `pgx_pipe`, `phm_pipe`, `photocd_pipe`, `pictor_pipe`, `pjs`, `pmp`, `png_pipe`, `pp_bnk`, `ppm_pipe`, `psd_pipe`, `psxstr`, `pva`, `pvf`, `qcp`, `qdraw_pipe`, `qoi_pipe`, `r3d`, `realtext`, `redspark`, `rl2`, `rpl`, `rsd`, `sami`, `sbg`, `scd`, `sdp`, `sdr2`, `sds`, `sdx`, `ser`, `sga`, `sgi_pipe`, `shn`, `siff`, `simbiosis_imx`, `sln`, `smk`, `smush`, `sol`, `stl`, `subviewer`, `subviewer1`, `sunrast_pipe`, `svag`, `svg_pipe`, `svs`, `tak`, `tedcaptions`, `thp`, `tiertexseq`, `tiff_pipe`, `tmv`, `tty`, `txd`, `ty`, `v210`, `v210x`, `vag`, `vbn_pipe`, `vividas`, `vivo`, `vmd`, `vobsub`, `vpk`, `vplayer`, `vqf`, `wc3movie`, `webp_pipe`, `wsd`, `wsvqa`, `wve`, `x11grab`, `xa`, `xbin`, `xmv`, `xpm_pipe`, `xvag`, `xwd_pipe`, `xwma`, `yop`, `video4linux2`, `v4l2`
 
- `start_number`: Starting number for the output file. Any non-negative integer.
- `length`: Length of the output file sequence. Any non-negative integer.


#### Codecs

The `codecs` attribute specifies the codecs and filters to apply.

```yaml
codecs:
  - input_id: input_1  # Identifier for the input
    output_id: output_1  # Identifier for the output
  - codec_name:
      audio: aac  # Audio codec
      video: libx264  # Video codec
    stream_loop: 1  # Number of times to loop the stream
    shortest: true  # Stop encoding when the shortest input stream ends
    preset: medium  # Preset for encoding speed
    crf: 23  # Constant Rate Factor for quality
    profile:
      video: high  # Video profile
      audio: aac_low  # Audio profile
    level: 4.1  # Level for encoding
    pixel_format: yuv420p  # Pixel format
    max_rate: 5000k  # Maximum bitrate
    buffer_size: 10000k  # Buffer size
    constant_bitrate:
      video: 4000k  # Constant video bitrate
      audio: 128k  # Constant audio bitrate
    file_size: 1000000  # Target file size
    audio_quality: 5  # Audio quality
    pass: 2  # Number of encoding passes
    audio_none: false  # No audio encoding
    video_none: false  # No video encoding
    move_flags: [faststart, frag_keyframe]  # Flags for moving the file
    metadata:
      - key: title  # Metadata key
        value: Sample Video  # Metadata value
    video_filters:
      - name: scale  # Video filter name
        value: 1280:720  # Video filter value
    audio_filters:
      - name: volume  # Audio filter name
        value: 1.5  # Audio filter value
    time_part:
      start_time: 00:00:10  # Start time for trimming
      duration_time: 00:00:30  # Duration for trimming
    concat_files:
      - source: part1.mp4  # Source file for concatenation
        duration: 10  # Duration of the segment
        in_point: 0  # In point for the segment
        out_point: 10  # Out point for the segment
    sync:
      audio: 1  # Audio sync method
      video: passthrough  # Video sync method
    frame:
      video: 30  # Video frame rate
      audio: 44100  # Audio sample rate
    quality:
      video: 23  # Video quality
      audio: 5  # Audio quality
    input_framerate: 30  # Input frame rate
    framerate: 30  # Output frame rate
    gop_size: 60  # Group of pictures size
    audio_sampling_rate: 44100  # Audio sampling rate
    hls:
      time: 10  # Segment duration for HLS
      list_size: 5  # Number of segments in the playlist
      segment_filename: segment_%03d.ts  # Segment filename pattern
      playlist_type: vod  # Playlist type
      segment_type: mpegts  # Segment type
      flags: [delete_segments]  # HLS flags
      master_playlist_name: master.m3u8  # Master playlist name
      segment_list: segment_list.m3u8  # Segment list file
      segment_list_size: 5  # Size of the segment list
      max_entries: 10  # Maximum number of entries
      allow_cache: true  # Allow caching
      key_info_file: key_info.txt  # Key info file
      key_url: http://example.com/key  # Key URL
    channels: 2  # Number of audio channels
    channel_layout: stereo  # Audio channel layout
    variant_stream_map: v:0,a:0  # Variant stream map
    map_input: true  # Map input streams
```

- `input_id`: Identifier for the input. not required. using for map codec to one of inputs.
- `output_id`: Identifier for the output. not required. using for map codec to one of outputs.
- `codec_name`: Specifies the codec to use. Corresponds to the `-c:v` and `-c:a` parameters in FFmpeg.
  - `audio`: Corresponds to the `-c:a` parameter in FFmpeg. Available values: `aac`, `ac3`, `mp3`, `opus`, `vorbis`, `flac`, `alac`, `pcm_s16le`, `pcm_s24le`, `pcm_s32le`, `pcm_f32le`, `pcm_f64le`, `pcm_mulaw`, `pcm_alaw`, `pcm_s8`, `pcm_u8`, `libmp3lame`, `libopus`, `libvorbis`, `copy`.
  - `video`: Corresponds to the `-c:v` parameter in FFmpeg. Available values: `libx264`, `libx265`, `mpeg2video`, `libvpx-vp9`, `gif`, `libvpx`, `libaom-av1`, `mpeg1video`, `mpeg4`, `h263`, `libtheora`, `prores`, `dnxhd`, `libxvid`, `msmpeg4v2`, `msmpeg4`, `wmv1`, `wmv2`, `vc1`, `flv`, `rawvideo`, `png`, `bmp`, `jpeg2000`, `mjpeg`, `huffyuv`, `liblags`, `copy`.
- `stream_loop`: Number of times to loop the stream. Corresponds to the `-stream_loop` parameter in FFmpeg. Any integer greater than or equal to -1.
- `shortest`: Stop encoding when the shortest input stream ends. Corresponds to the `-shortest` parameter in FFmpeg. Values are `true` or `false`.
- `preset`: Preset for encoding speed. Corresponds to the `-preset` parameter in FFmpeg. Available values: `veryslow`, `slower`, `slow`, `medium`, `fast`, `faster`, `veryfast`, `superfast`, `ultrafast`.
- `crf`: Constant Rate Factor for quality. Corresponds to the `-crf` parameter in FFmpeg. Any integer between 0 and 51.
- `profile`: Specifies the profile to use.
  - `video`: Video profile. Corresponds to the `-profile:v` parameter in FFmpeg. Available values: `baseline`, `main`, `high`, `main10`, `main12`, `simple`, `0`, `1`, `2`, `3`.
  - `audio`: Audio profile. Corresponds to the `-profile:a` parameter in FFmpeg. Available values: `aac_low`, `aac_he`, `aac_he_v2`, `aac_ld`, `aac_eld`, `ac3`.
- `level`: Level for encoding. Corresponds to the `-level` parameter in FFmpeg. Available values: `1.0`, `1.1`, `1.2`, `1.3`, `2.0`, `2.1`, `2.2`, `3.0`, `3.1`, `3.2`, `4.0`, `4.1`, `4.2`, `5.0`, `5.1`, `5.2`.
- `pixel_format`: Pixel format. Corresponds to the `-pix_fmt` parameter in FFmpeg. Available values: `yuv420p`, `yuv422p`, `yuv444p`, `yuv420p10le`, `yuv422p10le`, `yuv444p10le`, `yuv420p12le`, `yuv422p12le`, `yuv444p12le`, `rgb24`, `rgba`, `rgb48le`, `rgba64le`, `gray`, `gray16le`, `nv12`, `nv21`, `yuv420p16le`, `yuv422p16le`, `yuv444p16le`, `bgr24`, `bgra`.
- `max_rate`: Maximum bitrate. Corresponds to the `-maxrate` parameter in FFmpeg. Any valid bitrate string.
- `buffer_size`: Buffer size. Corresponds to the `-bufsize` parameter in FFmpeg. Any valid buffer size string.
- `constant_bitrate`: Specifies the constant bitrate.
  - `video`: Constant video bitrate. Corresponds to the `-b:v` parameter in FFmpeg. Any valid bitrate string.
  - `audio`: Constant audio bitrate. Corresponds to the `-b:a` parameter in FFmpeg. Any valid bitrate string.
- `file_size`: Target file size. Corresponds to the `-fs` parameter in FFmpeg. Any non-negative integer.
- `audio_quality`: Audio quality. Corresponds to the `-q:a` parameter in FFmpeg. Any integer between 0 and 9.
- `pass`: Number of encoding passes. Corresponds to the `-pass` parameter in FFmpeg. Available values: `1`, `2`.
- `audio_none`: No audio encoding. Corresponds to the `-an` parameter in FFmpeg. Values are `true` or `false`.
- `video_none`: No video encoding. Corresponds to the `-vn` parameter in FFmpeg. Values are `true` or `false`.
- `move_flags`: Flags for moving the file. Corresponds to the `-movflags` parameter in FFmpeg. Available values: `faststart`, `frag_keyframe`, `frag_custom`, `empty_moov`, `separate_moof`, `omit_tfhd_offset`, `rtphint`, `frag_discont`, `default_base_moof`, `delay_moov`, `negative_cts_offsets`, `disable_chpl`, `write_colr`.
- `metadata`: Metadata attributes.
  - `key`: Metadata key. Corresponds to the `-metadata` parameter in FFmpeg. Any non-empty string.
  - `value`: Metadata value. Corresponds to the `-metadata` parameter in FFmpeg. Any non-empty string.
- `video_filters`: Video filters to apply.
  - `name`: Name of the filter. Corresponds to the `-vf` parameter in FFmpeg. Any non-empty string.
  - `value`: Value of the filter. Corresponds to the `-vf` parameter in FFmpeg. Any non-empty string.
- `audio_filters`: Audio filters to apply.
  - `name`: Name of the filter. Corresponds to the `-af` parameter in FFmpeg. Any non-empty string.
  - `value`: Value of the filter. Corresponds to the `-af` parameter in FFmpeg. Any non-empty string.
- `time_part`: Specifies the time range to trim.
  - `start_time`: Start time for trimming. Corresponds to the `-ss` parameter in FFmpeg. Any valid time string.
  - `duration_time`: Duration for trimming. Corresponds to the `-t` parameter in FFmpeg. Any valid time string.
- `concat_files`: Files to concatenate.
  - `source`: Source file for concatenation. Corresponds to the `-i` parameter in FFmpeg. Any valid file path.
  - `duration`: Duration of the segment. Corresponds to the `-t` parameter in FFmpeg. Any non-negative integer.
  - `in_point`: In point for the segment. Corresponds to the `-ss` parameter in FFmpeg. Any non-negative integer.
  - `out_point`: Out point for the segment. Corresponds to the `-to` parameter in FFmpeg. Any non-negative integer.
- `sync`: Sync methods.
  - `audio`: Audio sync method. Corresponds to the `-async` parameter in FFmpeg. Any non-negative integer.
  - `video`: Video sync method. Corresponds to the `-vsync` parameter in FFmpeg. Available values: `passthrough`, `cfr`, `vfr`, `drop`, `0`.
- `frame`: Frame rates.
  - `video`: Video frame rate. Corresponds to the `-r` parameter in FFmpeg. Any non-negative integer.
  - `audio`: Audio sample rate. Corresponds to the `-ar` parameter in FFmpeg. Any non-negative integer.
- `quality`: Quality settings.
  - `video`: Video quality. Corresponds to the `-q:v` parameter in FFmpeg. Any integer between 1 and 31.
  - `audio`: Audio quality. Corresponds to the `-q:a` parameter in FFmpeg. Any integer between 0 and 9.
- `input_framerate`: Input frame rate. Corresponds to the `-r` parameter in FFmpeg. Any non-negative integer.
- `framerate`: Output frame rate. Corresponds to the `-r` parameter in FFmpeg. Any non-negative integer.
- `gop_size`: Group of pictures size. Corresponds to the `-g` parameter in FFmpeg. Any integer greater than or equal to -1.
- `audio_sampling_rate`: Audio sampling rate. Corresponds to the `-ar` parameter in FFmpeg. Any non-negative integer.
- `hls`: HLS options.
  - `time`: Segment duration for HLS. Corresponds to the `-hls_time` parameter in FFmpeg. Any non-negative float.
  - `list_size`: Number of segments in the playlist. Corresponds to the `-hls_list_size` parameter in FFmpeg. Any non-negative integer.
  - `segment_filename`: Segment filename pattern. Corresponds to the `-hls_segment_filename` parameter in FFmpeg. Any valid file name.
  - `playlist_type`: Playlist type. Corresponds to the `-hls_playlist_type` parameter in FFmpeg. Available values: `event`, `vod`.
  - `segment_type`: Segment type. Corresponds to the `-hls_segment_type` parameter in FFmpeg. Any valid segment type.
  - `flags`: HLS flags. Corresponds to the `-hls_flags` parameter in FFmpeg. Any valid flag.
  - `master_playlist_name`: Master playlist name. Corresponds to the `-master_pl_name` parameter in FFmpeg. Any valid file name.
  - `segment_list`: Segment list file. Corresponds to the `-hls_segment_list` parameter in FFmpeg. Any valid file path.
  - `segment_list_size`: Size of the segment list. Corresponds to the `-hls_segment_list_size` parameter in FFmpeg. Any non-negative integer.
  - `max_entries`: Maximum number of entries. Corresponds to the `-hls_segment_list_size` parameter in FFmpeg. Any non-negative integer.
  - `allow_cache`: Allow caching. Corresponds to the `-hls_allow_cache` parameter in FFmpeg. Values are `true` or `false`.
  - `key_info_file`: Key info file. Corresponds to the `-hls_key_info_file` parameter in FFmpeg. Any valid file path.
  - `key_url`: Key URL. Corresponds to the `-hls_key_url` parameter in FFmpeg. Any valid URL.
- `channels`: Number of audio channels. Corresponds to the `-ac` parameter in FFmpeg. Any non-negative integer.
- `channel_layout`: Audio channel layout. Corresponds to the `-channel_layout` parameter in FFmpeg. Available values: `mono`, `stereo`, `2.1`, `3.0`, `3.1`, `quad`, `4.0`, `4.1`, `5.0`, `5.1`, `6.1`, `7.0`, `7.1`, `hexagonal`, `octagonal`, `surround`, `quadraphonic`, `5.1(side)`, `7.1(wide)`, `ambisonic_first_order`, `ambisonic_second_order`, `ambisonic_third_order`.
- `variant_stream_map`: Variant stream map. Corresponds to the `-var_stream_map` parameter in FFmpeg. Any valid stream map.
- `map_input`: Map input streams. Corresponds to the `-map` parameter in FFmpeg. Values are `true` or `false`.

#### Run After

The `run_after` attribute specifies a list of tasks to run before this task.

```yaml
run_after:
  - Basic Video Conversion
```