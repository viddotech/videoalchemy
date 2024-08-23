package compose

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	log "github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/infrastructure/cli"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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

	validate, err := cli.NewValidator()
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
				log.Error(e)
			}
		}
		return composeData, err
	}

	// Use the color package to print a colorful success message
	success := color.New(color.FgGreen, color.Bold).PrintfFunc()
	success("🎉 Validation passed successfully!\n")

	return composeData, err
}

func GenerateErrorMessages(errs validator.ValidationErrors, trans ut.Translator, parent string) []string {
	var messages []string

	for _, err := range errs {
		pathString := strings.Replace(err.Namespace(), "ComposeFileSchema.", "", -1)
		namespace := strings.Split(pathString, ".")
		if parent != "" {
			namespace = append([]string{parent}, namespace...)
		}

		// Build the full path for the field
		fieldPath := strings.Join(namespace, ".")
		messages = append(messages, fmt.Sprintf("Validation failed on '%s': %s", fieldPath, err.Translate(trans)))
	}

	return messages
}
