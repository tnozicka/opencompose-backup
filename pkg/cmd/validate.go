package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

var (
	validateExample = `
		# Print validate information
		opencompose validate`
)

func NewCmdValidate(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate",
		Short:   "Print validate information",
		Long:    "Print validate information",
		Example: validateExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunValidate(out, cmd)
		},
	}
	return cmd
}

func RunValidate(out io.Writer, cmd *cobra.Command) error {
	return errors.New("===test error validate===")
}
