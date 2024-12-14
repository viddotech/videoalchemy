package compose

import (
	"github.com/go-playground/validator/v10"
	vavalidate "github.com/viddotech/videoalchemy/internal/infrastructure/compose/validate"
	"reflect"
	"strings"
)

func NewValidator() (*validator.Validate, error) {
	validate := validator.New()

	for tag, function := range vavalidate.VideoAlchemyValidatorFunc {
		err := validate.RegisterValidation(tag, function)
		if err != nil {
			return nil, err
		}
	}

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
