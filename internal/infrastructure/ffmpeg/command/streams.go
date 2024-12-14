package command

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GenerateFFMPEGCommand(inst schema.Instruction, allInstructions []schema.Instruction) ([]string, error) {
	var ffmpegCommand []string

	ffmpegCommand = append(ffmpegCommand, inst.Command)

	// Generate input params string
	inputParams, err := generateInputsParams(inst, allInstructions)
	if err != nil {
		return ffmpegCommand, err
	}
	ffmpegCommand = append(ffmpegCommand, inputParams...)

	// Grouping process streams that does not have any Output
	noOutputProcessStreams := findProcessStreamsNoOutput(inst.ProcessStreams)
	for _, processStream := range noOutputProcessStreams {
		ffmpegCommand = append(ffmpegCommand, GenerateProcessStream(processStream, inst)...)
	}

	// Generate Complex Filters
	complexFiltersString, err := GenerateComplexFilterParameters(inst, noOutputProcessStreams)
	ffmpegCommand = append(ffmpegCommand, "-filter_complex", complexFiltersString)

	// Grouping process streams by output ID
	var streamsPerOutput = make(map[string][]schema.ProcessStream)
	if len(inst.Outputs) == 1 {
		streamsPerOutput[inst.Outputs[0].ID] = inst.ProcessStreams
	} else if len(inst.Outputs) > 1 {
		streamsPerOutput, err = MapOutputToProcessStream(inst.ProcessStreams)
		if err != nil {
			return ffmpegCommand, err
		}
	} else {
		return ffmpegCommand, errors.New("no outputs found")
	}

	// Iterate process streams per output ID
	for outputID, processStreams := range streamsPerOutput {

		for _, processStream := range processStreams {
			ffmpegCommand = append(ffmpegCommand, GenerateProcessStream(processStream, inst)...)
		}

		outputParams, outputs, err := generateOutputParamsByID(outputID, inst)
		if err != nil {
			return ffmpegCommand, err
		}

		inst.Outputs = append(inst.Outputs, outputs...)
		ffmpegCommand = append(ffmpegCommand, outputParams...)
	}

	return ffmpegCommand, nil
}

func injectStreamsToTheStream(theStream schema.ProcessStream, inst schema.Instruction) []string {
	var injectedStreams []string

	for _, injectableStreamName := range theStream.InjectStreams {
		for _, s := range inst.ProcessStreams {
			if s.StreamName == injectableStreamName && theStream.StreamName != injectableStreamName {
				var inputIndex int
				if len(inst.Inputs) == 1 {
					inputIndex = 0
				} else {
					inputIndex, _ = getInputByID(s.StreamFrom.InputID, inst.Inputs)
				}
				mapTo := fmt.Sprintf("%d", inputIndex)
				if s.StreamFrom.StreamType != "" {
					shortStreamType := mapStreamTypeToShort(string(s.StreamFrom.StreamType))
					mapTo += fmt.Sprintf(":%s", shortStreamType)
				}
				if s.StreamFrom.StreamTypeIndex != nil {
					mapTo += fmt.Sprintf(":%d", *s.StreamFrom.StreamTypeIndex)
				}
				injectedStreams = append(injectedStreams, fmt.Sprintf("%s", mapTo))
			}
		}
	}
	return injectedStreams
}

func mapParamToStream(param string, streamType string, defaultStreamType string, streamTo *schema.StreamTo, streamFrom *schema.StreamFrom) string {
	mapParamTo := param
	if streamType != "" {
		mapParamTo += ":" + mapStreamTypeToShort(streamType)
	} else if defaultStreamType != "" {
		mapParamTo += ":" + mapStreamTypeToShort(defaultStreamType)
	}
	if streamTo != nil && streamTo.StreamTypeIndex != nil {
		mapParamTo += fmt.Sprintf(":%d", *streamTo.StreamTypeIndex)
	}
	if streamFrom != nil && streamFrom.StreamTypeIndex != nil {
		mapParamTo += fmt.Sprintf(":%d", *streamFrom.StreamTypeIndex)
	}
	return mapParamTo
}

func getStreamIndexByStreamType(streamName string, streamType string, processStreams []schema.ProcessStream) int {
	perStreamTypeCounter := make(map[string]int)
	for _, stream := range processStreams {
		if stream.StreamName == streamName && string(stream.StreamFrom.StreamType) == streamType {
			return perStreamTypeCounter[string(stream.StreamFrom.StreamType)]
		}
		if stream.StreamFrom != nil && stream.StreamFrom.StreamType != "" {
			perStreamTypeCounter[string(stream.StreamFrom.StreamType)]++
		}
	}
	return -1
}

func generateInputsParams(inst schema.Instruction, allInstructions []schema.Instruction) ([]string, error) {
	var inputs []string
	for _, input := range inst.Inputs {
		if input.FrameRate > 0 {
			inputs = append(inputs, "-framerate", fmt.Sprintf("%d", input.FrameRate))
		}
		if input.StreamLoop != 0 {
			inputs = append(inputs, "-stream_loop", fmt.Sprintf("%d", input.StreamLoop))
		}
		if input.RealTime {
			inputs = append(inputs, "-re")
		}
		if input.Format != "" {
			inputs = append(inputs, "-f", string(input.Format))
		}
		if input.SafePath != nil {
			if *input.SafePath {
				inputs = append(inputs, "-safe", "1")
			} else {
				inputs = append(inputs, "-safe", "0")
			}
		}
		if input.Source != "" {
			inputs = append(inputs, "-i", input.Source)
		} else if input.OutputID != "" {
			// Reference another output as an input
			output, err := GetOutputByRelatedOutputs(input.OutputID, inst.RunAfter, allInstructions)
			if err != nil {
				return inputs, err
			}
			inputs = append(inputs, "-i", output.Source)
		}
	}
	return inputs, nil
}

func isStreamInjectedToOtherStreams(streamName string, inst schema.Instruction) bool {
	if streamName == "" {
		return false
	}
	for _, processStream := range inst.ProcessStreams {
		for _, injected := range processStream.InjectStreams {
			if injected == streamName && processStream.StreamName != streamName {
				return true
			}
		}
	}
	return false
}

func generateOutputParamsByID(outputID string, inst schema.Instruction) ([]string, []schema.Output, error) {
	var params []string
	var outputs []schema.Output
	for _, output := range inst.Outputs {
		if output.ID == outputID {
			if output.Length > 0 {
				var startNumber uint
				if output.StartNum > 0 {
					params = append(params, "-start_number", fmt.Sprintf("%d", output.StartNum))
					startNumber = output.StartNum
				} else {
					startNumber = 1
				}
				params = append(params, output.Source)
				dir := filepath.Dir(output.Source)
				err := CreatePath(dir)
				if err != nil {
					return params, outputs, err
				}

				inst.FormatterOutputs = append(inst.FormatterOutputs, output)

				// Generate new output with formatted source value
				for i := startNumber; i < output.Length; i++ {
					newOutput := schema.Output{
						ID:        fmt.Sprintf(output.ID, i),
						OverWrite: output.OverWrite,
						Source:    fmt.Sprintf(output.Source, i),
						StartNum:  0,
						Length:    0,
					}
					outputs = append(outputs, newOutput)
				}
				for _, o := range outputs {
					log.Info(fmt.Sprintf("Generated Output -> ID: %s , Source:  %s", o.ID, o.Source))
				}

			} else {

				if output.FileSize > 0 {
					params = append(params, "-fs", fmt.Sprintf("%d", output.FileSize))
				}
				if output.HLS != nil {
					if output.HLS.Time > 0 {
						params = append(params, "-hls_time", fmt.Sprintf("%f", output.HLS.Time))
					}
					if output.HLS.ListSize > 0 {
						params = append(params, "-hls_list_size", fmt.Sprintf("%d", output.HLS.ListSize))
					}
					if output.HLS.SegmentFilename != "" {
						params = append(params, "-hls_segment_filename", fmt.Sprintf("%s", output.HLS.SegmentFilename))
					}
					if output.HLS.PlaylistType != "" {
						params = append(params, "-hls_playlist_type", fmt.Sprintf("%s", output.HLS.PlaylistType))
					}
					if output.HLS.SegmentType != "" {
						params = append(params, "-hls_segment_type", fmt.Sprintf("%s", output.HLS.SegmentType))
					}
					if output.HLS.Flags != "" {
						params = append(params, "-hls_flags", fmt.Sprintf("%s", output.HLS.Flags))
					}
					if output.HLS.MasterPlaylistName != "" {
						params = append(params, "-master_pl_name", fmt.Sprintf("%s", output.HLS.MasterPlaylistName))
					}
					if output.HLS.SegmentList != "" {
						params = append(params, "-hls_segment_list", fmt.Sprintf("%s", output.HLS.SegmentList))
					}
					if output.HLS.SegmentListSize > 0 {
						params = append(params, "-hls_segment_list_size", fmt.Sprintf("%d", output.HLS.SegmentListSize))
					}
					if output.HLS.MaxEntries > 0 {
						params = append(params, "-hls_max_entries", fmt.Sprintf("%d", output.HLS.MaxEntries))
					}
					if output.HLS.AllowCache {
						params = append(params, "-hls_allow_cache", "yes")
					}
					if output.HLS.KeyInfoFile != "" {
						params = append(params, "-hls_key_info_file", fmt.Sprintf("%s", output.HLS.KeyInfoFile))
					}
					if output.HLS.KeyURL != "" {
						params = append(params, "-hls_key_url", fmt.Sprintf("%s", output.HLS.KeyURL))
					}
				}

				if output.Format == "hls" || output.Format == "dash" {
					streamsPerOutput, err := MapOutputToProcessStream(inst.ProcessStreams)
					if err != nil {
						return params, outputs, err
					}
					processStreams := streamsPerOutput[output.ID]
					varStreamMap := generateVarStreamMap(processStreams)
					params = append(params, "-var_stream_map", varStreamMap)
				}

				if output.Format != "" {
					params = append(params, "-f", string(output.Format))
				}
				if output.Source != "" {
					if !strings.Contains(output.Source, "%") {
						dir := filepath.Dir(output.Source)
						err := CreatePath(dir)
						if err != nil {
							return params, outputs, err
						}
					}
					params = append(params, output.Source)
				}
				if output.OverWrite {
					params = append(params, "-y")
				}
				outputs = append(outputs, output)
			}
		}
	}

	return params, outputs, nil
}

func selectStream(processStream schema.ProcessStream, inst schema.Instruction) string {
	if processStream.StreamFrom.FilterOutputName != "" {
		return fmt.Sprintf("[%s]", processStream.StreamFrom.FilterOutputName)
	}

	var inputID string
	if len(inst.Inputs) == 1 {
		inputID = inst.Inputs[0].ID
	} else if processStream.StreamFrom.InputID != "" {
		inputID = processStream.StreamFrom.InputID
	} else {
		return ""
	}

	inputIndex, _ := getInputByID(inputID, inst.Inputs)
	connector := ""

	connector += strconv.Itoa(inputIndex)

	if processStream.StreamFrom.StreamType != "" {
		connector += ":" + mapStreamTypeToShort(string(processStream.StreamFrom.StreamType))
	}
	if processStream.StreamFrom.StreamTypeIndex != nil {
		connector += fmt.Sprintf(":%d", *processStream.StreamFrom.StreamTypeIndex)
	}
	return connector
}

func generateVarStreamMap(streams []schema.ProcessStream) string {
	var params []string
	for _, stream := range streams {
		if len(stream.InjectStreams) > 0 {
			params = append(params, generateSingleStreamMap(stream, streams))
		}
	}
	varStreamMap := strings.Join(params, " ")
	return varStreamMap
}

func generateSingleStreamMap(processStream schema.ProcessStream, streams []schema.ProcessStream) string {
	var params []string
	index := getStreamIndexByStreamType(processStream.StreamName, string(processStream.StreamFrom.StreamType), streams)
	short := mapStreamTypeToShort(string(processStream.StreamFrom.StreamType))
	params = append(params, fmt.Sprintf("%s:%d", short, index))
	for _, injectableStreamName := range processStream.InjectStreams {
		injectedStream, err := getStreamByStreamName(injectableStreamName, streams)
		if err != nil {
			log.Fatal(err)
		}
		shortStreamType := mapStreamTypeToShort(string(injectedStream.StreamFrom.StreamType))
		params = append(params, fmt.Sprintf("%s:%d", shortStreamType, index))
	}
	return strings.Join(params, ",")
}

func getStreamByStreamName(streamName string, processStreams []schema.ProcessStream) (schema.ProcessStream, error) {
	for _, processStream := range processStreams {
		if processStream.StreamName == streamName {
			return processStream, nil
		}
	}
	return schema.ProcessStream{}, errors.New("stream not found")
}

func CreatePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0750)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetOutputByRelatedOutputs(outputID string, runAfters []string, instructions []schema.Instruction) (schema.Output, error) {
	var result schema.Output

	if outputID != "" {
		for _, runAfter := range runAfters {
			for _, instruction := range instructions {
				if runAfter == instruction.Name {
					output, err := getOutputById(outputID, instruction.Outputs)
					if err == nil {
						return output, nil
					}
					if len(instruction.FormatterOutputs) > 0 {
						o, er := getOutputById(outputID, instruction.FormatterOutputs)
						if er == nil {
							return o, nil
						}
					}
				}
			}
		}
	}

	return result, errors.New(fmt.Sprintf("output id %s not found", outputID))
}

func getOutputById(outputID string, outputs []schema.Output) (schema.Output, error) {
	var result schema.Output
	if outputID != "" {
		for _, output := range outputs {
			if outputID == output.ID {
				return output, nil
			}
		}
	}

	return result, errors.New(fmt.Sprintf("output id %s not found", outputID))
}
