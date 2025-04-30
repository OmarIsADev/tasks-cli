package cmd

import "github.com/spf13/cobra"

func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "task managment cli tool",
		Long:  "task managing cli tool. It allows create, list, remove, modify your tasks",
	}

	cmd.AddCommand(NewAddCMD())
	cmd.AddCommand(NewListCMD())
	cmd.AddCommand(NewDeleteCMD())
	cmd.AddCommand(NewUpdateCMD())

	return cmd
}
