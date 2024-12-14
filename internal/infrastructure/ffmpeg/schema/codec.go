package schema

type SelectorField string

type CodecName struct {
	Video SelectorField `validate:"omitempty,oneof=libx264 libx265 mpeg2video libvpx-vp9 gif libvpx libaom-av1 mpeg1video mpeg4 h263 libtheora prores dnxhd libxvid msmpeg4v2 msmpeg4 wmv1 wmv2 vc1 flv rawvideo png bmp jpeg2000 mjpeg huffyuv liblags copy" yaml:"video"`
	Audio SelectorField `validate:"omitempty,oneof=aac ac3 mp3 opus vorbis flac alac pcm_s16le pcm_s24le pcm_s32le pcm_f32le pcm_f64le pcm_mulaw pcm_alaw pcm_s8 pcm_u8 libmp3lame libopus libvorbis copy" yaml:"audio"`
}

type Profile struct {
	Video SelectorField `validate:"omitempty,oneof=baseline main high main10 main12 simple 0 1 2 3" yaml:"video"`
	Audio SelectorField `validate:"omitempty,oneof=aac_low aac_he aac_he_v2 aac_ld aac_eld ac3" yaml:"audio"`
}

type Level struct {
	H264 SelectorField `validate:"omitempty,oneof=1.0 1.1 1.2 1.3 2.0 2.1 2.2 3.0 3.1 3.2 4.0 4.1 4.2 5.0 5.1 5.2" yaml:"h264"`
	H265 SelectorField `validate:"omitempty,oneof=1.0 1.1 2.0 2.1 3.0 3.1 4.0 4.1 5.0 5.1" yaml:"h265"`
}

type MetaDataAttribute struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type TimePart struct {
	StartTime    string `validate:"va_time" yaml:"start_time"`
	DurationTime string `validate:"va_time" yaml:"duration_time"`
	EndTime      string `validate:"va_time" yaml:"end_time"`
}

type ConcatFile struct {
	Source   string `yaml:"source"`
	Duration int    `yaml:"duration"`
	InPoint  int    `yaml:"in_point"`
	OutPoint int    `yaml:"out_point"`
}

type Sync struct {
	Audio uint          `yaml:"audio"`
	Video SelectorField `validate:"oneof=passthrough cfr vfr drop 0" yaml:"video"`
}

type Frame struct {
	Video uint `yaml:"video"`
	Audio uint `yaml:"audio"`
}

type Quality struct {
	Video uint8 `validate:"omitempty,min=1,max=31" yaml:"video"`
	Audio uint8 `validate:"omitempty,min=0,max=9" yaml:"audio"`
}

type ConstantBitrate struct {
	Video string `yaml:"video"`
	Audio string `yaml:"audio"`
}

type StreamFrom struct {
	InputID          string        `validate:"omitempty" yaml:"input_id"`
	StreamType       SelectorField `validate:"omitempty,oneof=video audio subtitle data attachment" yaml:"stream_type"`
	StreamTypeIndex  *uint8        `validate:"excluded_without=StreamType" yaml:"stream_type_index"`
	FilterOutputName string        `validate:"omitempty" yaml:"filter_output_name"`
	StreamName       string        `validate:"omitempty" yaml:"stream_name"`
}

type StreamTo struct {
	OutputID        string `validate:"required" yaml:"output_id"`
	StreamTypeIndex *uint8 `validate:"excluded_without=OutputID" yaml:"stream_type_index"`
}

type MaxRate struct {
	Video string `yaml:"video"`
	Audio string `yaml:"audio"`
}

// ProcessStream represents codec parameters based on FFmpeg's AVCodecContext.
type ProcessStream struct {
	StreamFrom        *StreamFrom         `validate:"stream_from__required,check_refs" yaml:"stream_from"`
	StreamTo          *StreamTo           `validate:"check_refs" yaml:"stream_to"`
	StreamName        string              `validate:"omitempty" yaml:"stream_name"`
	InjectStreams     []string            `validate:"dive" yaml:"inject_streams"`
	CodecName         CodecName           `validate:"omitempty" yaml:"codec_name"`
	Shortest          bool                `yaml:"shortest"`
	Preset            SelectorField       `validate:"omitempty,oneof=veryslow slower slow medium fast faster veryfast superfast altrafast" yaml:"preset"`
	Crf               uint8               `validate:"omitempty,min=0,max=51" yaml:"crf"`
	Profile           Profile             `yaml:"profile"`
	Level             SelectorField       `yaml:"level"`
	PixelFormat       SelectorField       `validate:"omitempty,oneof=yuv420p yuv422p yuv444p yuv420p10le yuv422p10le yuv444p10le yuv420p12le yuv422p12le yuv444p12le rgb24 rgba rgb48le rgba64le gray gray16le nv12 nv21 yuv420p16le yuv422p16le yuv444p16le bgr24 bgra" yaml:"pixel_format"`
	MaxRate           MaxRate             `yaml:"max_rate"`
	BufferSize        string              `yaml:"buffer_size"`
	ConstantBitrate   *ConstantBitrate    `yaml:"constant_bitrate"`
	AudioQuality      uint8               `validate:"omitempty,min=1,max=9" yaml:"audio_quality"`
	Pass              SelectorField       `validate:"omitempty,oneof=1 2" yaml:"pass"`
	AudioNone         bool                `yaml:"audio_none"`
	VideoNone         bool                `yaml:"video_none"`
	MoveFlags         []SelectorField     `validate:"dive,oneof=faststart frag_keyframe frag_custom empty_moov separate_moof omit_tfhd_offset rtphint frag_discont default_base_moof delay_moov negative_cts_offsets disable_chpl write_colr" yaml:"move_flags"`
	MetaData          []MetaDataAttribute `validate:"dive" yaml:"metadata"`
	VideoFilters      []VideoFilter       `validate:"dive,excluded_if=CodecName.Video copy" yaml:"video_filters"`
	AudioFilters      []AudioFilter       `validate:"dive,excluded_if=CodecName.Audio copy" yaml:"audio_filters"`
	TimePart          *TimePart           `validate:"omitnil" yaml:"time_part"`
	Sync              *Sync               `yaml:"sync"`
	Frame             *Frame              `yaml:"frame"`
	Quality           *Quality            `yaml:"quality"`
	FrameRate         uint                `yaml:"framerate"`
	Gop               int                 `validate:"omitempty,min=-1" yaml:"gop_size"`
	AudioSamplingRate uint                `yaml:"audio_sampling_rate"`
	Channels          int                 `yaml:"channels"`
	ChannelLayout     string              `validate:"omitempty,oneof=mono stereo 2.1 3.0 3.1 quad 4.0 4.1 5.0 5.1 6.1 7.0 7.1 hexagonal octagonal surround quadraphonic 5.1(side) 7.1(wide) ambisonic_first_order ambisonic_second_order ambisonic_third_order" yaml:"channel_layout"`
}
