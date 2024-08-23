package main

import (
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
	"log"
	"os"
)

func main() {
	taskService := services.TaskService{}

	composeCommand := RootCommand(taskService)

	if err := composeCommand.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
