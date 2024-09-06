package ffmpeg

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFFMPEGCommand(inst schema.Instruction, allInstructions []schema.Instruction, GeneratePath string) ([]string, error) {
	var command []string

	command = append(command, inst.Command)

	for _, codec := range inst.Codec {
		if codec.InputFrameRate > 0 {
			command = append(command, "-framerate", fmt.Sprintf("%d", codec.InputFrameRate))
		}
		if codec.StreamLoop != 0 {
			command = append(command, "-stream_loop", fmt.Sprintf("%d", codec.StreamLoop))
		}
	}

	// Add inputs to the command
	for _, input := range inst.Inputs {
		if input.RealTime {
			command = append(command, "-re")
		}
		if input.Format != "" {
			command = append(command, "-f", string(input.Format))
		}
		if input.Source != "" {
			command = append(command, "-i", string(input.Source))
		} else if input.OutputID != "" {
			// Reference another output as an input
			output, err := getOutputByRelatedOutputs(input.OutputID, inst.RunAfter, allInstructions)
			if err != nil {
				return command, err
			}
			command = append(command, "-i", output.Source)
		}
	}

	// Add codec options to the command
	for _, codec := range inst.Codec {
		outputMapperIndex := getOutputIndexByCodec(codec, inst.Outputs)

		if len(inst.Inputs) > 1 || (len(inst.Inputs) == 1 && codec.MapInput) {
			inputMapperIndex := getInputIndexByCodec(codec, inst.Inputs)
			command = append(command, "-map", fmt.Sprintf("%d:v", inputMapperIndex))
			command = append(command, "-map", fmt.Sprintf("%d:a", inputMapperIndex))
		}
		if codec.CodecName.Video != "" {
			command = append(command, fmt.Sprintf("-c:v:%d", outputMapperIndex), string(codec.CodecName.Video))
		}
		if codec.CodecName.Audio != "" {
			command = append(command, "-c:a", string(codec.CodecName.Audio))
		}
		if codec.Preset != "" {
			command = append(command, "-preset", string(codec.Preset))
		}
		if codec.Crf > 0 {
			command = append(command, "-crf", fmt.Sprintf("%d", codec.Crf))
		}
		if codec.Profile.Video != "" {
			command = append(command, "-profile:v", string(codec.Profile.Video))
		}
		if codec.Profile.Audio != "" {
			command = append(command, "-profile:a", string(codec.Profile.Audio))
		}
		if codec.Level != "" {
			command = append(command, "-level", string(codec.Level))
		}
		if codec.PixelFormat != "" {
			command = append(command, "-pix_fmt", string(codec.PixelFormat))
		}
		if codec.MaxRate != "" {
			command = append(command, fmt.Sprintf("-maxrate:v:%d", outputMapperIndex), fmt.Sprintf("%s", codec.MaxRate))
		}
		if codec.FrameRate > 0 {
			command = append(command, "-r", fmt.Sprintf("%d", codec.FrameRate))
		}
		if codec.BufferSize != "" {
			command = append(command, fmt.Sprintf("-bufsize:v:%d", outputMapperIndex), fmt.Sprintf("%s", codec.BufferSize))
		}
		if codec.ConstantBitrate != nil {
			if codec.ConstantBitrate.Video != "" {
				command = append(command, fmt.Sprintf("-b:v:%d", outputMapperIndex), fmt.Sprintf("%s", codec.ConstantBitrate.Video))
			}
			if codec.ConstantBitrate.Audio != "" {
				command = append(command, fmt.Sprintf("-b:a:%d", outputMapperIndex), fmt.Sprintf("%s", codec.ConstantBitrate.Audio))
			}
		}
		if codec.FileSize > 0 {
			command = append(command, "-fs", fmt.Sprintf("%d", codec.FileSize))
		}
		if codec.Sync != nil {
			if codec.Sync.Audio > 0 {
				command = append(command, "-async", fmt.Sprintf("%d", codec.Sync.Audio))
			}
			if codec.Sync.Video != "" {
				command = append(command, "-vsync", string(codec.Sync.Video))
			}
		}
		if codec.Frame != nil {
			if codec.Frame.Video > 0 {
				command = append(command, "-frames:v", fmt.Sprintf("%d", codec.Frame.Video))
			}
			if codec.Frame.Audio > 0 {
				command = append(command, "-frames:a", fmt.Sprintf("%d", codec.Frame.Audio))
			}
		}
		if codec.Quality != nil {
			if codec.Quality.Video > 0 {
				command = append(command, "-q:v", fmt.Sprintf("%d", codec.Quality.Video))
			}
			if codec.Quality.Audio > 0 {
				command = append(command, "-q:a", fmt.Sprintf("%d", codec.Quality.Audio))
			}
		}
		if codec.Pass != "" {
			command = append(command, "-pass "+string(codec.Pass))
		}
		if codec.AudioNone {
			command = append(command, "-an")
		}
		if codec.VideoNone {
			command = append(command, "-vn")
		}

		if len(codec.VideoFilters) > 0 {
			var filters []string
			for _, f := range codec.VideoFilters {
				filter := fmt.Sprintf("%s=%s", f.Name, f.Value)
				filters = append(filters, filter)
			}
			command = append(command, "-vf", strings.Join(filters, ","))
		}
		if len(codec.AudioFilters) > 0 {
			var filters []string
			for _, f := range codec.AudioFilters {
				var filter string
				if f.Value != "" {
					filter = fmt.Sprintf("%s=%s", f.Name, f.Value)
				} else {
					filter = fmt.Sprintf("%s", f.Name)
				}
				filters = append(filters, filter)
			}
			command = append(command, "-af", strings.Join(filters, ","))
		}

		if codec.TimePart != nil {
			command = append(command, "-ss", codec.TimePart.StartTime)
			if codec.TimePart.DurationTime != "" {
				command = append(command, "-t", codec.TimePart.DurationTime)
			} else {
				command = append(command, "-to", codec.TimePart.EndTime)
			}
		}

		if len(codec.ConcatFiles) > 0 {
			concatFilePath := fmt.Sprintf("%s/%s/concats.txt", GeneratePath, inst.Name)
			err := CreateConcatInputFiles(concatFilePath, codec.ConcatFiles)
			if err != nil {
				return command, err
			}
			command = append(command, "-f", "concat", "-safe", "0", "-i", concatFilePath)
		}

		for _, flag := range codec.MoveFlags {
			command = append(command, "-movflags", fmt.Sprintf("+%s", string(flag)))
		}
		for _, metadata := range codec.MetaData {
			command = append(command, "-metadata", fmt.Sprintf("%s=%s", metadata.Key, metadata.Value))
		}
		if codec.Shortest {
			command = append(command, "-shortest")
		}
		if codec.Gop > 0 {
			command = append(command, "-g", fmt.Sprintf("%d", codec.Gop))
		}

		if codec.AudioSamplingRate > 0 {
			command = append(command, "-ar", fmt.Sprintf("%d", codec.AudioSamplingRate))
		}

		if codec.HLS != nil {
			if codec.HLS.Time > 0 {
				command = append(command, "-hls_time", fmt.Sprintf("%f", codec.HLS.Time))
			}
			if codec.HLS.ListSize > 0 {
				command = append(command, "-hls_list_size", fmt.Sprintf("%d", codec.HLS.ListSize))
			}
			if codec.HLS.SegmentFilename != "" {
				command = append(command, "-hls_segment_filename", fmt.Sprintf("%s", codec.HLS.SegmentFilename))
			}
			if codec.HLS.PlaylistType != "" {
				command = append(command, "-hls_playlist_type", fmt.Sprintf("%s", codec.HLS.PlaylistType))
			}
			if codec.HLS.SegmentType != "" {
				command = append(command, "-hls_segment_type", fmt.Sprintf("%s", codec.HLS.SegmentType))
			}
			if codec.HLS.Flags != "" {
				command = append(command, "-hls_flags", fmt.Sprintf("%s", codec.HLS.Flags))
			}
			if codec.HLS.MasterPlaylistName != "" {
				command = append(command, "-master_pl_name", fmt.Sprintf("%s", codec.HLS.MasterPlaylistName))
			}
			if codec.HLS.SegmentList != "" {
				command = append(command, "-hls_segment_list", fmt.Sprintf("%s", codec.HLS.SegmentList))
			}
			if codec.HLS.SegmentListSize > 0 {
				command = append(command, "-hls_segment_list_size", fmt.Sprintf("%d", codec.HLS.SegmentListSize))
			}
			if codec.HLS.MaxEntries > 0 {
				command = append(command, "-hls_max_entries", fmt.Sprintf("%d", codec.HLS.MaxEntries))
			}
			if codec.HLS.AllowCache {
				command = append(command, "-hls_allow_cache", "yes")
			}
			if codec.HLS.KeyInfoFile != "" {
				command = append(command, "-hls_key_info_file", fmt.Sprintf("%s", codec.HLS.KeyInfoFile))
			}
			if codec.HLS.KeyURL != "" {
				command = append(command, "-hls_key_url", fmt.Sprintf("%s", codec.HLS.KeyURL))
			}
			if codec.VariantStreamMap != "" {
				command = append(command, "-var_stream_map", fmt.Sprintf("%s", codec.VariantStreamMap))
			}
			if codec.Channels > 0 {
				command = append(command, "-ac", fmt.Sprintf("%d", codec.Channels))
			}
			if codec.ChannelLayout != "" {
				command = append(command, "-channel_layout", fmt.Sprintf("%s", codec.ChannelLayout))
			}
		}
	}

	var newOutputs []schema.Output
	for _, output := range inst.Outputs {
		if output.Length > 0 {
			var startNumber uint
			if output.StartNum > 0 {
				command = append(command, "-start_number", fmt.Sprintf("%d", output.StartNum))
				startNumber = output.StartNum
			} else {
				startNumber = 1
			}
			command = append(command, output.Source)
			dir := filepath.Dir(output.Source)
			err := CreatePath(dir)
			if err != nil {
				return command, err
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
				newOutputs = append(newOutputs, newOutput)
			}
			for _, o := range newOutputs {
				log.Info(fmt.Sprintf("Generated Output -> ID: %s , Source:  %s", o.ID, o.Source))
			}

		} else {
			if output.Format != "" {
				command = append(command, "-f", string(output.Format))
			}
			if output.Source != "" {
				if !strings.Contains(output.Source, "%") {
					dir := filepath.Dir(output.Source)
					err := CreatePath(dir)
					if err != nil {
						return command, err
					}
				}
				command = append(command, output.Source)
			}
			if output.OverWrite {
				command = append(command, "-y")
			}
			newOutputs = append(newOutputs, output)
		}
	}

	// Set the new outputs to inst.Outputs
	inst.Outputs = newOutputs

	return command, nil
}

func CreateConcatInputFiles(inputFilePath string, files []schema.ConcatFile) error {
	var concatFile string
	for _, file := range files {
		absSource, err := filepath.Abs(file.Source)
		if err != nil {
			return err
		}
		concatFile += fmt.Sprintf("file '%s'\n", absSource)
		if file.Duration > 0 {
			concatFile += fmt.Sprintf("duration %d\n", file.Duration)
		}
		if file.InPoint > 0 {
			concatFile += fmt.Sprintf("in_point %d\n", file.InPoint)
		}
		if file.OutPoint > 0 {
			concatFile += fmt.Sprintf("out_point %d\n", file.OutPoint)
		}
	}
	absPath, err := filepath.Abs(inputFilePath)
	if err != nil {
		return err
	}
	dir := filepath.Dir(absPath)
	err = CreatePath(dir)
	if err != nil {
		return err
	}
	err = os.WriteFile(inputFilePath, []byte(concatFile), 0600)
	if err != nil {
		return err
	}
	return nil
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

func getInputIndexByCodec(codec schema.CodecSchema, inputs []schema.Input) int {
	if codec.InputID != "" {
		for index, input := range inputs {
			if codec.InputID == input.ID {
				return index
			}
		}
	}
	return 0
}

func getOutputIndexByCodec(codec schema.CodecSchema, outputs []schema.Output) int {
	if codec.InputID != "" {
		for index, output := range outputs {
			if codec.OutputID == output.ID {
				return index
			}
		}
	}
	return 0
}

func getOutputByRelatedOutputs(outputID string, runAfters []string, instructions []schema.Instruction) (schema.Output, error) {
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
