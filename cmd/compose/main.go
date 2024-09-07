package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
	"os"
)

// These variables will be set at build time
var (
	version = "dev"     // default version, if not provided during build
	date    = "unknown" // build date
)

func main() {
	fmt.Printf("Version: %s, Build Date: %s\n", version, date)

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
