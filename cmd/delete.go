package cmd

import (
	"errors"
	"strconv"

	"github.com/omarisadev/tasks-cli/task"
	"github.com/spf13/cobra"
)

func NewDeleteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes task by id",
		Long: `Deletes task by id

	Example:
	task-cli delete 1
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteTask(args)
		},
	}

	return cmd
}

func deleteTask(args []string) error {
	if len(args) == 0 {
		return errors.New("no id was provided")
	}

	id, err := strconv.ParseUint(args[0], 0, 16)

	if err != nil {
		return errors.New("invalid id")
	}

	return task.DeleteTask(uint16(id))
}
