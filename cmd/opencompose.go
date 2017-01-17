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
		fmt.Printf("ERROR: %s\n", err)
	}
	return err
}
