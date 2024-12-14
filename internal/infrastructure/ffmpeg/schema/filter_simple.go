package schema

type VideoFilter struct {
	Name  string `validate:"oneof=select scale crop transpose vflip hflip rotate fps pad setpts eq hue brightness contrast gamma sharpness unsharp colorbalance lut colorchannelmixer overlay blend fade split tile geq noise negate curves boxblur gblur edgedetect vignette fade subtitles drawtext yadif bwdif pp null settb ass" yaml:"name"`
	Value string `yaml:"value"`
}

type AudioFilter struct {
	Name  string `validate:"oneof=aselect volume lowpass highpass loudnorm adelay aecho atempo areverse" yaml:"name"`
	Value string `yaml:"value"`
}
