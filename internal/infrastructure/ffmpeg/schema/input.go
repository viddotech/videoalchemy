package schema

type InputStream struct {
	Index uint8
	ID    string
	Type  SelectorField
	Data  map[string]interface{}
}

type Input struct {
	ID           string        `validate:"required" yaml:"id"`
	Source       string        `validate:"required_without=OutputID" yaml:"source"`
	OutputID     string        `validate:"required_without=Source" yaml:"output_id"`
	RealTime     bool          `yaml:"realtime"`
	Format       SelectorField `validate:"omitempty,oneof=3dostr 4xm aa aac aax ac3 acm act adf adp ads aea afc aiff aix alias_pix amr amrnb amrwb anm apc ape apng aptx aptx_hd aqtitle argo_brp asf asf_o ass ast au av1 avi avm2 avr avs bethsoftvid bfi bfstm bin bink binka bit bmp_pipe brstm c93 caf cavsvideo cdg cdxl cine codec2 codec2raw concat ffconcat crc dash data daud dds_pipe derf dfa dhav dirac dnxhd dpx_pipe dsf dsicin dss dts dv dvbsub dvbtxt dxa ea ea_cdata eac3 epaf exr_pipe f32be f32le f64be f64le ffmetadata fifo fifo_test film_cpk flic flv frm fsb g722 g723_1 g726 g726le g729 gdv genh gif_pipe gsm gxf h261 h263 h264 hca hcom hevc hls hnm ico_pipe idcin idf iec61883 iff ifv ilbc image2 image2pipe ingenient ipmovie ipu ircam iss iv8 ivf ivr j2k_pipe jacosub jpeg_pipe jpegls_pipe jpegxl_pipe jv kux latm lavfi libgme live_flv lmlm4 loas lrc luodat lvf lxf m4v mca mcc md5 mgsts microdvd mjpeg_2000 mjpeg mlp mlv mm mmf mods moflex mov mp2 mp3 mp4 m4a 3gp 3g2 mj2 mpc mpc8 mpeg mpeg1video mpeg2video mpegts mpegtsraw mpegvideo mpl2 mpsub msf msnwctcp msp mtaf mtv musx mv mvi mxg nc nistsphere nsp nsv null nut obu ogg oma paf pam_pipe pbm_pipe pcm_alaw pcm_f32be pcm_f32le pcm_f64be pcm_f64le pcm_mulaw pcm_s16be pcm_s16le pcm_s24be pcm_s24le pcm_s32be pcm_s32le pcm_s8 pcm_u16be pcm_u16le pcm_u24be pcm_u24le pcm_u32be pcm_u32le pcm_u8 pjs pmp png_pipe pp_bnk ppm_pipe prores psd_pipe psxstr pva pvf qcp qdraw_pipe qoi_pipe r3d rawvideo realtext redspark rl2 rm roq rpl rsd rso rtp rtp_mpegts rtsp s16be s16le s24be s24le s32be s32le sami sap sbc sbg scc scd sdp sdr2 sds sdx ser sga sgi_pipe shn siff simbiosis_imx sln smk smush sol sox spdif srt stl stream_segment subviewer subviewer1 sunrast_pipe sup svag svg_pipe svs swf tak tedcaptions thp tiertexseq tiff_pipe tmv truehd tta tty txd ty v210 v210x vag vbn_pipe vc1 vc1test vivo vmd vobsub voc vpk vplayer vqf w64 wav wc3movie webm webm_chunk webm_dash_manifest webp_pipe webvtt wsaud wsd wsvqa wtv wve xa xbin xmv xpm_pipe xvag xwma yop yuv4mpegpipe zmbv video4linux2 v4l2 vfwcap" yaml:"format"`
	FrameRate    uint          `yaml:"framerate"`
	StreamLoop   int           `validate:"omitempty,min=-1" yaml:"stream_loop"`
	SafePath     *bool         `validate:"omitempty" yaml:"safe_path"`
	InputStreams []InputStream
}
