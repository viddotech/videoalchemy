package ffmpeg

import (
	"fmt"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"strings"
)

func GenerateFFMPEGCommand(inst schema.Instruction) (string, error) {
	var command []string

	// Add inputs to the command
	for _, input := range inst.Inputs {
		if input.Source != "" {
			command = append(command, "-i", string(input.Source))
		} else if input.OutputID != "" {
			// Reference another output as an input
			command = append(command, "-i", fmt.Sprintf("[%s]", input.OutputID))
		}
	}

	// Add codec options to the command
	for _, codec := range inst.Codec {
		mapperIndex := getInputIndexByCodec(codec, inst.Inputs)
		command = append(command, fmt.Sprintf("-map %d:v", mapperIndex))
		if codec.InputID != "" {
			command = append(command, fmt.Sprintf("[%s]", codec.InputID))
		}
		if codec.CodecName.Video != "" {
			command = append(command, "-c:v", string(codec.CodecName.Video))
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
		if codec.PixFmt != "" {
			command = append(command, "-pix_fmt", string(codec.PixFmt))
		}
		if codec.MaxRate > 0 {
			command = append(command, "-maxrate", fmt.Sprintf("%d", codec.MaxRate))
		}
		if codec.BufferSize > 0 {
			command = append(command, "-bufsize", fmt.Sprintf("%d", codec.BufferSize))
		}
		if codec.ConstantBitrate > 0 {
			command = append(command, "-b:v", fmt.Sprintf("%d", codec.ConstantBitrate))
		}
		if codec.FileSize > 0 {
			command = append(command, "-fs", fmt.Sprintf("%d", codec.FileSize))
		}
		if codec.AudioQuality > 0 {
			command = append(command, "-q:a", fmt.Sprintf("%d", codec.AudioQuality))
		}
		if codec.Pass != "" {
			command = append(command, "-pass", string(codec.Pass))
		}
		if codec.An {
			command = append(command, "-an")
		}
		for _, flag := range codec.MoveFlags {
			command = append(command, "-movflags", string(flag))
		}
		for _, metadata := range codec.MetaData {
			command = append(command, "-metadata", fmt.Sprintf("%s=%s", metadata.Key, metadata.Value))
		}
	}

	// Add outputs to the command
	for _, output := range inst.Outputs {
		if output.Source != "" {
			command = append(command, "-map", fmt.Sprintf("[%s]", output.ID), string(output.Source))
		} else {
			command = append(command, "-map", fmt.Sprintf("[%s]", output.ID))
		}
	}

	result := strings.Join(command, " ")

	return result, nil
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
