package schema

type SelectorField string

type CodecName struct {
	Video SelectorField `validate:"omitempty,oneof=libx264 libx265 mpep-2 vp9 copy" yaml:"video"`
	Audio SelectorField `validate:"omitempty,oneof=aac ac3 mp3 opus" yaml:"audio"`
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

// CodecSchema represents codec parameters based on FFmpeg's AVCodecContext.
type CodecSchema struct {
	InputID         string              `validate:"omitempty" yaml:"input_id"`
	CodecName       CodecName           `validate:"omitempty" yaml:"codec_name"`
	Preset          SelectorField       `validate:"omitempty,oneof=veryslow slower slow medium fast faster veryfast superfast altrafast" yaml:"preset"`
	Crf             uint8               `validate:"omitempty,min=0,max=51" yaml:"crf"`
	Profile         Profile             `yaml:"profile"`
	Level           SelectorField       `yaml:"level"`
	PixFmt          SelectorField       `validate:"omitempty,oneof=yuv420p yuv422p yuv444p yuv420p10le yuv422p10le yuv444p10le yuv420p12le yuv422p12le yuv444p12le rgb24 rgba rgb48le rgba64le gray gray16le nv12 nv21 yuv420p16le yuv422p16le yuv444p16le bgr24 bgra" yaml:"pix_fmt"`
	MaxRate         uint64              `yaml:"max_rate"`
	BufferSize      uint64              `yaml:"buffer_size"`
	ConstantBitrate uint64              `yaml:"constant_bitrate"`
	FileSize        uint32              `yaml:"file_size"`
	AudioQuality    uint8               `validate:"omitempty,min=1,max=9" yaml:"audio_quality"`
	Pass            SelectorField       `validate:"omitempty,oneof=1 2" yaml:"pass"`
	An              bool                `yaml:"an"`
	MoveFlags       []SelectorField     `validate:"dive,oneof=faststart frag_keyframe frag_custom empty_moov separate_moof omit_tfhd_offset rtphint frag_discont default_base_moof delay_moov negative_cts_offsets disable_chpl write_colr" yaml:"move_flags"`
	MetaData        []MetaDataAttribute `validate:"dive" yaml:"metadata"`

	//SampleRate    int    `yaml:"sample_rate"`    // Sample rate (only for audio codecs)
	//Channels      int    `yaml:"channels"`       // Number of audio channels (only for audio codecs)
	//ChannelLayout uint64 `yaml:"channel_layout"` // Channel layout (only for audio codecs)
	//FrameSize     int    `yaml:"frame_size"`     // Frame size (only for audio codecs)
}
