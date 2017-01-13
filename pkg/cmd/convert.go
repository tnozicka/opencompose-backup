package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	cmdutil "github.com/tnozicka/opencompose/pkg/cmd/util"
)

var (
	convertExample = `
		# Converts file
		opencompose convert -f opencompose.yaml`
)

func NewCmdConvert(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "convert",
		Short:   "Converts OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Long:    "Converts OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Example: convertExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunConvert(out, cmd)
		},
	}

	cmdutil.AddIOFlags(cmd)

	return cmd
}

func RunConvert(out io.Writer, cmd *cobra.Command) error {
	files, err := cmd.PersistentFlags().GetStringSlice("file")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "files: %#v\n", files)

	loglevel, err := cmd.Flags().GetInt8("loglevel")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "loglevel: %#v\n", loglevel)

	return errors.New("===test error convert===")
}
