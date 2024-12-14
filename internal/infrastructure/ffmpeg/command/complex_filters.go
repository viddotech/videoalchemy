package command

import (
	"fmt"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"strings"
)

func GenerateComplexFilterParameters(inst schema.Instruction, noOutputProcessStreams []schema.ProcessStream) (string, error) {
	var ffmpegString []string
	for _, filter := range inst.ComplexFilters {
		filterString := ""

		// Generate strings of inputs by Input ID
		if filter.StreamFrom.InputID != "" && filter.StreamFrom.StreamType != "" {
			inputIndex, _ := getInputByID(filter.StreamFrom.InputID, inst.Inputs)
			streamType := mapStreamTypeToShort(string(filter.StreamFrom.StreamType))
			filterString += fmt.Sprintf("[%d:%s]", inputIndex, streamType)
		}

		// Generate strings of inputs By output of other complex filters
		if filter.StreamFrom.FilterOutputName != "" {
			filterString += fmt.Sprintf("[%s]", filter.StreamFrom.FilterOutputName)
		}

		// Generate strings of inputs by Stream name
		if filter.StreamFrom.StreamName != "" {
			for _, stream := range noOutputProcessStreams {
				if stream.StreamName == filter.StreamFrom.StreamName {
					//	TODO
				}
			}
		}

		var filterItemsString []string
		for _, filterItem := range filter.Filters {
			filterItemsString = append(filterItemsString, fmt.Sprintf("%s=%s", filterItem.Name, filterItem.Value))
		}
		filterString += strings.Join(filterItemsString, ",")

		// Generate strings of Outputs
		if len(filter.OutputsName) > 0 {
			for _, filterOutput := range filter.OutputsName {
				filterString += fmt.Sprintf("[%s]", filterOutput)
			}
		}

		ffmpegString = append(ffmpegString, filterString)
	}

	return strings.Join(ffmpegString, ";"), nil
}
