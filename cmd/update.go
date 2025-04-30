package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/omarisadev/tasks-cli/task"
	"github.com/spf13/cobra"
)

func NewUpdateCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates task description or status",
		Long: `Updates task description or status

	Example:
	task-cli update description 1 'new description'
	task-cli update status 1 2

	1 - TASK_STATUS_DONE
	2 - TASK_STATUS_IN_PROGRESS
	3 - TASK_STATUS_TODO
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return updateTask(args)
		},
	}

	return cmd
}

func updateTask(args []string) error {
	if len(args) < 3 {
		return errors.New("invalid number of args")
	}

	action, value := args[0], args[2]

	id, err := strconv.ParseUint(args[1], 0, 16)
	if err != nil {
		return errors.New("invalid id")
	}

	switch action {
	case "description":
		return task.UpdateTaskDescription(uint16(id), value)

	case "status":
		var status task.TaskStatus

		switch value {
		case "1":
			status = task.TASK_STATUS_DONE
		case "2":
			status = task.TASK_STATUS_IN_PROGRESS
		case "3":
			status = task.TASK_STATUS_TODO

		default:
			fmt.Println("status should be either 1, 2 or 3")
			return errors.New("invalid status value")
		}

		return task.UpdateTaskStatus(uint16(id), status)
	}

	return errors.New("invalid args")
}
