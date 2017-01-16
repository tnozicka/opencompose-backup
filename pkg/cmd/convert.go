package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	cmdutil "github.com/tnozicka/opencompose/pkg/cmd/util"
	"github.com/spf13/viper"
)

var (
	convertExample = `  # Converts file
  opencompose convert -f opencompose.yaml`
)

func NewCmdConvert(v *viper.Viper, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "convert",
		Short:   "Converts OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Long:    "Converts OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Example: convertExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunConvert(v, cmd, out)
		},
	}

	cmdutil.AddIOFlags(v, cmd)

	return cmd
}

func RunConvert(v *viper.Viper, cmd *cobra.Command, out io.Writer) error {
	files := v.GetStringSlice(cmdutil.Flag_File_Key)
	if len(files) < 1 {
		return NewUsageError(cmd.Help, "You have to specify at least one input file option")
	}
	fmt.Fprintf(out, "files: %#v\n", files)


	return errors.New("===test error convert===")
}
