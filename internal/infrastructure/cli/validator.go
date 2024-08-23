package cli

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func NewValidator() (*validator.Validate, error) {
	validate := validator.New()

	// TODO: Implement custom validator to define output if needed output.
	// TODO: Implement custom validator to if instruction will have multiple output must be define every output.
	// TODO: Implement custom validator to if defined input from other instruction output must be define current task as run_after.

	// Set up the validator to use the YAML tag name instead of the field name
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate, nil
}
