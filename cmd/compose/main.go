package main

import (
	"github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
	"os"
)

func main() {

	taskService := services.TaskService{}

	logrus.SetFormatter(&VideoAlchemyLogFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp: true,
		},
	})

	composeCommand := RootCommand(taskService)

	if err := composeCommand.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
