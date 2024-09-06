package schema

type Instruction struct {
	Name             string        `validate:"required" yaml:"name"`
	Inputs           []Input       `validate:"dive" yaml:"inputs"`
	Codec            []CodecSchema `validate:"dive" yaml:"codecs"`
	Outputs          []Output      `validate:"dive" yaml:"outputs"`
	FormatterOutputs []Output
	Command          string   `validate:"required" yaml:"command"`
	RunAfter         []string `validate:"dive,omitempty" yaml:"run_after"`
}
