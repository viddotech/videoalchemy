package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"time"
)

var VideoAlchemyValidatorFunc = map[string]validator.Func{
	REQUIRED_STREAM_FROM: validateStreamFromRequired,
	VA_TIME:              validateVATime,
	CHECK_REFERENCES:     checkRefs,
}

func validateStreamFromRequired(fl validator.FieldLevel) bool {
	streamFrom, ok := fl.Field().Interface().(schema.StreamFrom)
	if !ok {
		return false
	}
	if streamFrom.InputID == "" && streamFrom.FilterOutputName == "" && streamFrom.StreamName == "" {
		return false
	}
	return true
}

func validateVATime(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	if timeStr != "" {
		format := "15:04:05.000" // Format for HH:MM:SS
		_, err := time.Parse(format, timeStr)
		return err == nil
	}
	return true
}

func checkRefs(fl validator.FieldLevel) bool {

	composeFile, ok := fl.Top().Interface().(schema.ComposeFileSchema)
	if !ok {
		// If the parent is not of type ComposeFileSchema, we can't proceed
		return false
	}

	fieldInterface := fl.Field().Interface()

	for _, instruction := range composeFile.Instructions {

		if streamFrom, ok := fieldInterface.(schema.StreamFrom); ok { // check references of stream from parameters
			// Check if inputID references a valid input within the same instruction
			for _, in := range instruction.Inputs {
				if in.ID == streamFrom.InputID {
					return true
				}
			}

			// Check if streamName references a valid process stream within the same instruction
			for _, stream := range instruction.ProcessStreams {
				if stream.StreamName == streamFrom.StreamName {
					return true
				}
			}

			// Check if filterOutputName references a complex filter within the same instruction
			for _, complexFilter := range instruction.ComplexFilters {
				for _, outputName := range complexFilter.OutputsName {
					if outputName == streamFrom.FilterOutputName {
						return true
					}
				}
			}

		} else if streamTo, ok := fieldInterface.(schema.StreamTo); ok { // check references of stream from parameters
			// Check if outputID references a valid output within the same instruction
			for _, out := range instruction.Outputs {
				if out.ID == streamTo.OutputID {
					return true
				}
			}
		}

	}

	return false

}
