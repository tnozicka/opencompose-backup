package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	validateExample = `
		# Print validate information
		opencompose validate`
)

func NewCmdValidate(v *viper.Viper, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate",
		Short:   "Print validate information",
		Long:    "Print validate information",
		Example: validateExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunValidate(v, cmd, out)
		},
	}
	return cmd
}

func RunValidate(v *viper.Viper, cmd *cobra.Command, out io.Writer) error {
	return errors.New("===test error validate===")
}
