package compose

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	log "github.com/sirupsen/logrus"
	validate "github.com/viddotech/videoalchemy/internal/infrastructure/compose/validate"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/command"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func LoadComposeDataFromFile(composeFilePath string) (schema.ComposeFileSchema, error) {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	// Set the output to standard output
	log.SetOutput(os.Stdout)

	data, err := ioutil.ReadFile(composeFilePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var composeData schema.ComposeFileSchema
	err = yaml.Unmarshal(data, &composeData)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	composeData = LoadInstruction(composeData)

	validate, err := NewValidator()
	if err != nil {
		log.Fatal(err)
		return schema.ComposeFileSchema{}, err
	}

	uni := ut.New(en.New(), en.New())
	trans, _ := uni.GetTranslator("en")

	// Register default translations for all validation rules
	err = enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal("Error registering translations:", err)
		return schema.ComposeFileSchema{}, err
	}

	err = validate.Struct(composeData)
	if err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			// Generate and print the custom error messages
			for _, e := range GenerateErrorMessages(errs, trans, "") {
				fmt.Println(e)
			}
		}
		return composeData, err
	}

	// TODO: Notify Validated Successfully

	return composeData, err
}

func LoadInstruction(composeData schema.ComposeFileSchema) schema.ComposeFileSchema {
	for _, instruction := range composeData.Instructions {
		for _, input := range instruction.Inputs {
			var source string
			if input.OutputID != "" {
				output, err := command.GetOutputByRelatedOutputs(input.OutputID, instruction.RunAfter, composeData.Instructions)
				if err != nil {
					log.Fatalf("Failed to get output by related outputs: %v", err)
				}
				source = output.Source
			} else {
				source = input.Source
			}
			streamsData, err := ExtractInputStreams(composeData.Inspector, source)
			if err == nil {
				streams := CreateStreams(input, streamsData)
				input.InputStreams = append(input.InputStreams, streams...)
				instruction.InputStreams = append(instruction.InputStreams, streams...)
			}

		}
	}
	return composeData
}

func ExtractInputStreams(inspector schema.Inspector, inputPath string) ([]interface{}, error) {
	if inspector.CommandType != "ffprobe" {
		return nil, nil
	}

	cmd := exec.Command(inspector.Path, "-v", "error", "-show_streams", "-print_format", "json", inputPath)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Warnf("Failed to create stdout pipe: %v", err)
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		log.Warnf("Failed to start ffprobe: %v", err)
		return nil, err
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var output []byte
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line...)
	}

	if err := scanner.Err(); err != nil {
		log.Warnf("Error reading ffprobe output: %v", err)
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		log.Warnf("ffprobe command failed: %v", err)
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		log.Warnf("Failed to unmarshal ffprobe output: %v", err)
		return nil, err
	}

	return result["streams"].([]interface{}), nil
}

func CreateStreams(input schema.Input, streamsData []interface{}) []schema.InputStream {
	var streams []schema.InputStream
	for _, sData := range streamsData {

		var stream schema.InputStream
		data := sData.(map[string]interface{})
		stream.Index = uint8(data["index"].(float64))
		stream.ID = fmt.Sprintf("%s|%d", input.ID, stream.Index)
		stream.Type = schema.SelectorField(data["codec_type"].(string))
		stream.Data = data

		streams = append(streams, stream)
	}

	return streams
}

func GenerateErrorMessages(errs validator.ValidationErrors, trans ut.Translator, parent string) []string {
	var messages []string

	paramMsgStyle := color.RGB(175, 175, 175).Add(color.Bold).SprintFunc()
	messageStyle := color.New(color.FgRed).Add(color.Bold).SprintFunc()

	for _, err := range errs {
		pathString := strings.Replace(err.Namespace(), "ComposeFileSchema.", "", -1)
		namespace := strings.Split(pathString, ".")
		if parent != "" {
			namespace = append([]string{parent}, namespace...)
		}

		// Build the full path for the field
		fieldPath := strings.Join(namespace, ".")

		translatedErr := fmt.Sprintf(messageStyle("Validation Error: %s => %s"), paramMsgStyle(fieldPath), paramMsgStyle(validate.MapErrorTags[err.Tag()]))
		messages = append(messages, translatedErr)
	}

	return messages
}
