package schema

type ComplexFilterItem struct {
	Name  string `validate:"oneof=select scale crop transpose vflip hflip rotate fps pad setpts eq hue brightness contrast gamma sharpness unsharp colorbalance lut colorchannelmixer overlay blend fade split tile geq noise negate curves boxblur gblur edgedetect vignette fade subtitles drawtext yadif bwdif pp null settb ass aselect volume lowpass highpass loudnorm adelay aecho atempo areverse" yaml:"name"`
	Value string `yaml:"value"`
}

type ComplexFilter struct {
	StreamFrom  *StreamFrom         `validate:"stream_from__required,check_refs" yaml:"stream_from"`
	OutputsName []string            `yaml:"outputs_name"`
	Filters     []ComplexFilterItem `yaml:"filters"`
}
