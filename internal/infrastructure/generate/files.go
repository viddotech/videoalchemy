package generate

import (
	"fmt"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/schema"
	"os"
)

func CreateConcatFilesList(files []schema.ConcatFile, filePath, InstructionName string) error {
	// Create a list of files to concatenate
	var concatFile string
	for _, file := range files {
		concatFile += fmt.Sprintf("file '%s'\n", file.Source)
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
	err := os.WriteFile(filePath, []byte(concatFile), 0600)
	if err != nil {
		return err
	}
	return nil
}
