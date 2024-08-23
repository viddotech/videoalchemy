package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/viddotech/videoalchemy/internal/domain/task/entities"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg"
	"github.com/viddotech/videoalchemy/internal/infrastructure/ffmpeg/compose"
	"log"
	"time"
)

type TaskService struct {
}

func (s *TaskService) CreateTasks(composeFilePath string) ([]entities.Task, error) {
	composeFileData, err := compose.LoadComposeDataFromFile(composeFilePath)
	if err != nil {
		log.Fatal("Error load & validate compose file: ", err)
	}

	tasks := make([]entities.Task, len(composeFileData.Instructions))

	for i, instruction := range composeFileData.Instructions {
		ffmpegCommand, err := ffmpeg.GenerateFFMPEGCommand(instruction)

		if err != nil {
			log.Fatal("Error generate Ffmpeg command : ", err)
		}
		tasks[i] = entities.Task{
			ID:            entities.TaskID(uuid.New().String()),
			Instruction:   instruction,
			FFMPEGCommand: &ffmpegCommand,
			CreatedAt:     time.Now(),
			UpdatedAt:     nil,
		}
		fmt.Println("ffmpeg command : ", *tasks[i].FFMPEGCommand)
	}
	return tasks, nil
}
