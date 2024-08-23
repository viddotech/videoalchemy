package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
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
			_, err := taskService.CreateTasks(composeFilePath)
			if err != nil {
				return err
			}
			//fmt.Printf("Task created: %+v %+v \n", task.Instruction.Name, composeFilePath)
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&composeFilePath, "compose-file", "f", "./viddo-compose.yaml", fmt.Sprintf("default is %+v", composeFilePath))

	return cmd
}
