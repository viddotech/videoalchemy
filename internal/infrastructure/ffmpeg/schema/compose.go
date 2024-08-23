package schema

type Version string

type ComposeFileSchema struct {
	Version      Version       `yaml:"version" validate:"required"`
	Instructions []Instruction `yaml:"tasks" validate:"required,dive"`
}
