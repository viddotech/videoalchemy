package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
	"github.com/viddotech/videoalchemy/internal/infrastructure/compose"
	"github.com/viddotech/videoalchemy/internal/infrastructure/pretty"
)

func RootCommand(taskService services.TaskService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "VideoAlchemy",
		Short: "VideoAlchemy is on of utility of Viddo!",
	}

	cmd.AddCommand(ComposeCommand(taskService))

	return cmd
}

func ComposeCommand(taskService services.TaskService) *cobra.Command {
	var composeFilePath string

	cmd := &cobra.Command{
		Use:   "compose",
		Short: "Manage Media Tasks",
		RunE: func(cmd *cobra.Command, args []string) error {

			composeFileData, err := compose.LoadComposeDataFromFile(composeFilePath)
			if err != nil {
				return nil
			}

			err = taskService.CreateTasks(composeFileData.Instructions, composeFileData.GeneratePath)
			if err != nil {
				return err
			}

			done := taskService.RunTasks(composeFileData.GeneratePath)

			if done {
				pretty.NotifySuccessText("Hooray! All tasks are done 🎉")
			} else {
				pretty.NotifyDangerousText("Oh no! Some tasks have failed 😢")
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&composeFilePath, "compose-file", "f", "./viddo-compose.yaml", fmt.Sprintf("default is %+v", composeFilePath))

	return cmd
}
