package cmd

import (
	"fmt"
	"os"

	"github.com/tnozicka/opencompose/pkg/cmd"
)

func Run() error {
	command := cmd.NewOpenComposeCommand(os.Stdin, os.Stdout, os.Stderr)
	err := command.Execute()
	if err != nil {
		switch e := err.(type) {
		case cmd.UsageError:
			fmt.Printf("Usage ERROR: %s\n\n", e)
			e.Help()
		default:
			fmt.Printf("Error: %s\n", err)
		}
	}
	return err
}
