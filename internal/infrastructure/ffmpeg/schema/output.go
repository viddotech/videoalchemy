package schema

type Output struct {
	ID        string        `validate:"required" yaml:"id"`
	OverWrite bool          `validate:"required" yaml:"overwrite"`
	Source    string        `validate:"required" yaml:"source"`
	StartNum  uint          `validate:"omitempty" yaml:"start_number"`
	Length    uint          `validate:"omitempty" yaml:"length"`
	Format    SelectorField `validate:"omitempty,oneof=3g2 3gp a64 adts amv aptx aptx_hd asf asf_stream ass ast au avi avif avm2 bit caf cavsvideo codec2 codec2raw crc dash data daud dfpwm dirac dnxhd dts dv dvd eac3 f4v fifo fifo_test film_cpk filmstrip fits flac flv framecrc framehash framemd5 g722 g723_1 g726 g726le gif gsm gxf h261 h263 h264 hds hevc hls ico ilbc image2 image2pipe ipod ircam ismv ivf jacosub kvag latm lrc m4v matroska md5 microdvd mjpeg mkvtimestamp_v2 mlp mmf mov mp2 mp3 mp4 mpeg mpeg1video mpeg2video mpegts mpjpeg mulaw mxf mxf_d10 mxf_opatom nut obu oga ogg ogv oma opus oss psp rawvideo rm roq rso rtp rtp_mpegts rtsp sap sbc scc segment smjpeg smoothstreaming sox spdif spx srt stream_segment ssegment streamhash sup svcd swf tee truehd tta ttml u16be u16le u24be u24le u32be u32le u8 uncodedframecrc vc1 vc1test vcd vidc video4linux2 v4l2 vob voc wav webm webm_chunk webm_dash_manifest webp webvtt wsaud wtv wv xv yuv4mpegpipe" yaml:"format"`
}
