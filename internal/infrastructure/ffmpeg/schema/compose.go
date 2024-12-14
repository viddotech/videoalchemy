package schema

type Version string

type Inspector struct {
	Path        string `yaml:"path"`
	CommandType string `validate:"oneof=ffprobe" yaml:"command_type"`
}

type ComposeFileSchema struct {
	Version      Version       `yaml:"version" validate:"required,oneof=1.0 1"`
	GeneratePath string        `yaml:"generate_path" validate:"required"`
	Inspector    Inspector     `yaml:"inspector" validate:"required"`
	Instructions []Instruction `yaml:"tasks" validate:"required,unique=Name,dive"`
}
