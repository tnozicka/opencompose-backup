package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cmdutil "github.com/tnozicka/opencompose/pkg/cmd/util"
)

const (
	Flag_Distro_Key    = "distro"
	Flag_OutputDir_Key = "output-dir"
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
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Parent().PersistentPreRunE != nil {
				if err := cmd.Parent().PersistentPreRunE(cmd, args); err != nil {
					return err
				}
			}

			// We have to bind Viper in Run because there is only one instance to avoid collisions between subcommands
			cmdutil.AddIOFlagsViper(v, cmd)
			cmdutil.BindViper(v, cmd.PersistentFlags(), Flag_Distro_Key)

			return nil
		},
	}

	cmdutil.AddIOFlags(cmd)
	cmd.PersistentFlags().StringP(Flag_Distro_Key, "d", "kubernetes", "Choose a target distribution")

	return cmd
}

func RunConvert(v *viper.Viper, cmd *cobra.Command, out io.Writer) error {

	//od := v.GetString(cmdutil.Flag_OutputDir_Key)
	dir := v.Get("output-dir")
	fmt.Fprintf(out, "output-dir: %#v\n", dir)

	d := v.Get("distro")
	fmt.Fprintf(out, "distro: %#v\n", d)

	//fs := v.Get(cmdutil.Flag_File_Key)
	//fmt.Fprintf(out, "files: %#v\n", fs)

	//files := v.GetStringSlice(cmdutil.Flag_File_Key)
	//fmt.Fprintf(out, "files: %#v\n", files)
	//if len(files) < 1 {
	//	return NewUsageError(cmd.Help, "You have to specify at least one input file option")
	//}

	return errors.New("===test error convert===")
}
