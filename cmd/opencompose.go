package cmd

import (
	"os"
	"fmt"

	"github.com/tnozicka/opencompose/pkg/cmd"
)

func Run() error {
	cmd := cmd.NewOpenComposeCommand(os.Stdin, os.Stdout, os.Stderr)
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return err
}
