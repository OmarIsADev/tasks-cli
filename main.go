package main

import (
	"github.com/omarisadev/tasks-cli/cmd"

	"fmt"
)

func main() {
	rootCmd := cmd.NewRootCMD()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Err: " + err.Error())
	}
}
