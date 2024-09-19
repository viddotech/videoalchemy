package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/viddotech/videoalchemy/internal/domain/task/services"
	"github.com/viddotech/videoalchemy/internal/infrastructure/compose"
	"github.com/viddotech/videoalchemy/internal/infrastructure/pretty"
	"os"
)

// These variables will be set at build time
var (
	version = "dev"     // default version, if not provided during build
	date    = "unknown" // build date
)

func RootCommand(taskService services.TaskService) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "VideoAlchemy",
		Short: "Simplify your media workflows with ease!",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				return
			}
		},
	}

	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number of VideoAlchemy")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("Version: %s\n", version)
			os.Exit(0)
		}
	}

	rootCmd.AddCommand(ComposeCommand(taskService))

	return rootCmd
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
