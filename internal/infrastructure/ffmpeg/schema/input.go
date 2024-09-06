package schema

type Input struct {
	ID       string        `validate:"required" yaml:"id"`
	Source   string        `validate:"required_without=OutputID" yaml:"source"`
	OutputID string        `validate:"required_without=Source" yaml:"output_id"`
	RealTime bool          `yaml:"realtime"`
	Format   SelectorField `validate:"omitempty,oneof=3dostr 4xm aa aac aax acm act adf adp ads aea afc aix alias_pix amrnb amrwb anm apc ape aqtitle argo_brp asf_o au av1 avi avr avs bethsoftvid bfi bfstm bin bink binka brstm c93 cdg cdxl cine concat cri_pipe dcstr dds_pipe derf dfa dhav dpx_pipe dsf dsicin dss dvbsub dvbtxt dxa ea ea_cdata epaf exr_pipe flic frm fsb fwse g729 gdv gem_pipe genh gif_pipe hca hcom hnm idcin idf iff ifv imf ingenient ipmovie ipu iss iv8 ivr j2k_pipe jpeg_pipe jpegls_pipe jpegxl_pipe jv kux lavfi libgme live_flv lmlm4 loas luodat lvf lxf mca mcc mgsts mjpeg_2000 mlv mm mods moflex mov mp4 m4a 3gp 3g2 mj2 mpc mpc8 mpegtsraw mpegvideo mpl2 mpsub msf msnwctcp msp mtaf mtv musx mv mvi mxg nc nistsphere nsp nsv paf pam_pipe pbm_pipe pcx_pipe pfm_pipe pgm_pipe pgmyuv_pipe pgx_pipe phm_pipe photocd_pipe pictor_pipe pjs pmp png_pipe pp_bnk ppm_pipe psd_pipe psxstr pva pvf qcp qdraw_pipe qoi_pipe r3d realtext redspark rl2 rpl rsd sami sbg scd sdp sdr2 sds sdx ser sga sgi_pipe shn siff simbiosis_imx sln smk smush sol stl subviewer subviewer1 sunrast_pipe svag svg_pipe svs tak tedcaptions thp tiertexseq tiff_pipe tmv tty txd ty v210 v210x vag vbn_pipe vividas vivo vmd vobsub vpk vplayer vqf wc3movie webp_pipe wsd wsvqa wve x11grab xa xbin xmv xpm_pipe xvag xwd_pipe xwma yop video4linux2 v4l2" yaml:"format"`
}
