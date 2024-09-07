package schema

type Version string

type ComposeFileSchema struct {
	Version      Version       `yaml:"version" validate:"required,oneof=1.0 1"`
	GeneratePath string        `yaml:"generate_path" validate:"required"`
	Instructions []Instruction `yaml:"tasks" validate:"required,unique=Name,dive"`
}
