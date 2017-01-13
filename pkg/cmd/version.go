package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/tnozicka/opencompose/pkg/version"
)

var (
	versionExample = `
		# Print version information
		kubectl version`
)

func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print version information",
		Long:    "Print version information",
		Example: versionExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunVersion(out, cmd)
		},
	}
	return cmd
}

func RunVersion(out io.Writer, cmd *cobra.Command) error {
	info := version.Get()
	fmt.Fprintln(out, info)
	return nil
}
