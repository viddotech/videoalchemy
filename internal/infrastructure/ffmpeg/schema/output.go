package schema

type Output struct {
	ID     string       `validate:"required" yaml:"id"`
	Source FFMPEGSource `yaml:"source" validate:"required"`
}
