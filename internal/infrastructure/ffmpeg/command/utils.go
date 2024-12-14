package command

import (
	"fmt"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"strings"
)

func getInputByID(id string, inputs []schema.Input) (int, *schema.Input) {
	for index, input := range inputs {
		if id == input.ID {
			return index, &input
		}
	}

	return 0, nil
}

func mapStreamTypeToShort(streamType string) string {
	switch streamType {
	case "video":
		return "v"
	case "audio":
		return "a"
	case "subtitle":
		return "s"
	case "data":
		return "d"
	case "attachment":
		return "t"
	default:
		return ""
	}
}

func findProcessStreamsNoOutput(processStreams []schema.ProcessStream) []schema.ProcessStream {
	var streamsNoOutput []schema.ProcessStream
	for _, processStream := range processStreams {
		if processStream.StreamTo == nil {
			streamsNoOutput = append(streamsNoOutput, processStream)
		}
	}
	return streamsNoOutput
}

func MapOutputToProcessStream(processStreams []schema.ProcessStream) (map[string][]schema.ProcessStream, error) {
	var streamsPerOutput = make(map[string][]schema.ProcessStream)
	for _, processStream := range processStreams {
		if processStream.StreamTo.OutputID != "" {
			streamsPerOutput[processStream.StreamTo.OutputID] = append(streamsPerOutput[processStream.StreamTo.OutputID], processStream)
		}
	}
	return streamsPerOutput, nil
}

func GenerateProcessStream(processStream schema.ProcessStream, inst schema.Instruction) []string {
	var ffmpegCommand []string
	if !isStreamInjectedToOtherStreams(processStream.StreamName, inst) {
		selectedStream := selectStream(processStream, inst)
		if selectedStream != "" {
			ffmpegCommand = append(ffmpegCommand, "-map", selectedStream)
		}
	}

	injectedStreamParams := injectStreamsToTheStream(processStream, inst)
	if len(injectedStreamParams) > 0 {
		ffmpegCommand = append(ffmpegCommand, "-map")
		ffmpegCommand = append(ffmpegCommand, injectedStreamParams...)
	}

	streamType := ""
	if processStream.StreamFrom != nil {
		streamType = string(processStream.StreamFrom.StreamType)
	}

	if processStream.CodecName.Video != "" {
		videoCodec := mapParamToStream("-c", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoCodec, string(processStream.CodecName.Video))
	}
	if processStream.CodecName.Audio != "" {
		audioCodec := mapParamToStream("-c", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, audioCodec, string(processStream.CodecName.Audio))
	}
	if processStream.Preset != "" {
		videoPreset := mapParamToStream("-preset", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoPreset, string(processStream.Preset))
	}
	if processStream.Crf > 0 {
		videoCrf := mapParamToStream("-crf", "", "", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoCrf, fmt.Sprintf("%d", processStream.Crf))
	}
	if processStream.Profile.Video != "" {
		videoProfile := mapParamToStream("-profile", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoProfile, string(processStream.Profile.Video))
	}
	if processStream.Profile.Audio != "" {
		audioProfile := mapParamToStream("-profile", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, audioProfile, string(processStream.Profile.Audio))
	}
	if processStream.Level != "" {
		videoLevel := mapParamToStream("-level", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoLevel, string(processStream.Level))
	}
	if processStream.PixelFormat != "" {
		videoPixelFormat := mapParamToStream("-pix_fmt", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoPixelFormat, string(processStream.PixelFormat))
	}
	if processStream.MaxRate.Video != "" {
		videoMaxRate := mapParamToStream("-maxrate", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoMaxRate, processStream.MaxRate.Video)
	}
	if processStream.MaxRate.Audio != "" {
		audioMaxRate := mapParamToStream("-maxrate", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, audioMaxRate, processStream.MaxRate.Audio)
	}
	if processStream.FrameRate > 0 {
		videoFrameRate := mapParamToStream("-r", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoFrameRate, fmt.Sprintf("%d", processStream.FrameRate))
	}
	if processStream.BufferSize != "" {
		videoBufferSize := mapParamToStream("-bufsize", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
		ffmpegCommand = append(ffmpegCommand, videoBufferSize, processStream.BufferSize)
	}
	if processStream.ConstantBitrate != nil {
		if processStream.ConstantBitrate.Video != "" {
			videoConstantBitrate := mapParamToStream("-b", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, videoConstantBitrate, processStream.ConstantBitrate.Video)
		}
		if processStream.ConstantBitrate.Audio != "" {
			audioConstantBitrate := mapParamToStream("-b", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, audioConstantBitrate, fmt.Sprintf("%s", processStream.ConstantBitrate.Audio))
		}
	}
	if processStream.Sync != nil {
		if processStream.Sync.Audio > 0 {
			ffmpegCommand = append(ffmpegCommand, "-async", fmt.Sprintf("%d", processStream.Sync.Audio))
		}
		if processStream.Sync.Video != "" {
			ffmpegCommand = append(ffmpegCommand, "-vsync", string(processStream.Sync.Video))
		}
	}
	if processStream.Frame != nil {
		if processStream.Frame.Video > 0 {
			videoFrame := mapParamToStream("-frames", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, videoFrame, fmt.Sprintf("%d", processStream.Frame.Video))
		}
		if processStream.Frame.Audio > 0 {
			audioFrame := mapParamToStream("-frames", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, audioFrame, fmt.Sprintf("%d", processStream.Frame.Audio))
		}
	}
	if processStream.Quality != nil {
		if processStream.Quality.Video > 0 {
			videoQuality := mapParamToStream("-q", streamType, "video", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, videoQuality, fmt.Sprintf("%d", processStream.Quality.Video))
		}
		if processStream.Quality.Audio > 0 {
			audioQuality := mapParamToStream("-q", streamType, "audio", processStream.StreamTo, processStream.StreamFrom)
			ffmpegCommand = append(ffmpegCommand, audioQuality, fmt.Sprintf("%d", processStream.Quality.Audio))
		}
	}
	if processStream.Pass != "" {
		ffmpegCommand = append(ffmpegCommand, "-pass "+string(processStream.Pass))
	}
	if processStream.AudioNone {
		ffmpegCommand = append(ffmpegCommand, "-an")
	}
	if processStream.VideoNone {
		ffmpegCommand = append(ffmpegCommand, "-vn")
	}

	if len(processStream.VideoFilters) > 0 {
		var filters []string
		for _, f := range processStream.VideoFilters {
			filter := fmt.Sprintf("%s=%s", f.Name, f.Value)
			filters = append(filters, filter)
		}
		ffmpegCommand = append(ffmpegCommand, "-vf", strings.Join(filters, ","))
	}
	if len(processStream.AudioFilters) > 0 {
		var filters []string
		for _, f := range processStream.AudioFilters {
			var filter string
			if f.Value != "" {
				filter = fmt.Sprintf("%s=%s", f.Name, f.Value)
			} else {
				filter = fmt.Sprintf("%s", f.Name)
			}
			filters = append(filters, filter)
		}
		ffmpegCommand = append(ffmpegCommand, "-af", strings.Join(filters, ","))
	}

	if processStream.TimePart != nil {
		ffmpegCommand = append(ffmpegCommand, "-ss", processStream.TimePart.StartTime)
		if processStream.TimePart.DurationTime != "" {
			ffmpegCommand = append(ffmpegCommand, "-t", processStream.TimePart.DurationTime)
		} else {
			ffmpegCommand = append(ffmpegCommand, "-to", processStream.TimePart.EndTime)
		}
	}

	for _, flag := range processStream.MoveFlags {
		ffmpegCommand = append(ffmpegCommand, "-movflags", fmt.Sprintf("+%s", string(flag)))
	}
	for _, metadata := range processStream.MetaData {
		ffmpegCommand = append(ffmpegCommand, "-metadata", fmt.Sprintf("%s=%s", metadata.Key, metadata.Value))
	}
	if processStream.Shortest {
		ffmpegCommand = append(ffmpegCommand, "-shortest")
	}
	if processStream.Gop > 0 {
		ffmpegCommand = append(ffmpegCommand, "-g", fmt.Sprintf("%d", processStream.Gop))
	}

	if processStream.AudioSamplingRate > 0 {
		ffmpegCommand = append(ffmpegCommand, "-ar", fmt.Sprintf("%d", processStream.AudioSamplingRate))
	}
	return ffmpegCommand
}
