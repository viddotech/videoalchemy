package schema

type Input struct {
	ID       string       `validate:"required" yaml:"id"`
	Source   FFMPEGSource `validate:"required_without=OutputID" yaml:"source"`
	OutputID string       `validate:"required_without=Source" yaml:"output_id"`
}
