package cmd

import (
	"errors"

	"github.com/omarisadev/tasks-cli/task"
	"github.com/spf13/cobra"
)

func NewAddCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Adds new task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addTask(args)
		},
	}

	return cmd
}

func addTask(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	return task.AddTask(args[0])
}
