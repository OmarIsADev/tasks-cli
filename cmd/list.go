package cmd

import (
	"github.com/omarisadev/tasks-cli/task"
	"github.com/spf13/cobra"
)

func NewListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List saved tasks",
		Long: `List all tasks. You can filter tasks by status

	Example:
	task-cli list todo
	task-cli list in-progress
	task-cli list done
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listTasks(args)
		},
	}

	return cmd
}

func listTasks(args []string) error {
	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.ListTasks(status)
	}

	return task.ListTasks("all")
}
