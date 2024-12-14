package schema

type Instruction struct {
	Name             string          `validate:"required" yaml:"name"`
	Inputs           []Input         `validate:"dive" yaml:"inputs"`
	ProcessStreams   []ProcessStream `validate:"unique=StreamName,dive" yaml:"streams"`
	Outputs          []Output        `validate:"dive" yaml:"outputs"`
	ComplexFilters   []ComplexFilter `validate:"dive" yaml:"complex_filters"`
	InputStreams     []InputStream
	FormatterOutputs []Output
	Command          string   `validate:"required" yaml:"command"`
	RunAfter         []string `validate:"dive,omitempty" yaml:"run_after"`
}
