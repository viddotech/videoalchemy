package compose

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"time"
)

func NewValidator() (*validator.Validate, error) {
	validate := validator.New()

	err := validate.RegisterValidation("time", validateTime)
	if err != nil {
		return nil, err
	}

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

func validateTime(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	if timeStr != "" {
		format := "15:04:05.000" // Format for HH:MM:SS
		_, err := time.Parse(format, timeStr)
		return err == nil
	}
	return true
}
