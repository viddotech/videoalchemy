package schema

type Version string

type ComposeFileSchema struct {
	Version      Version       `yaml:"version" validate:"required"`
	GeneratePath string        `yaml:"generate_path" validate:"required"`
	Instructions []Instruction `yaml:"tasks" validate:"required,unique=Name,dive"`
}
