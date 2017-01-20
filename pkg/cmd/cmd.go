package cmd

import (
	"io"

	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cmdutil "github.com/tnozicka/opencompose/pkg/cmd/util"
)

const (
	Flag_Verbose_Key = "loglevel"
)

func NewOpenComposeCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	v := viper.New()
	v.SetEnvPrefix("opencompose")
	v.AutomaticEnv()

	// Parent command to which all subcommands are added.
	rootCmd := &cobra.Command{
		Use:   "opencompose",
		Short: "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Long:  "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) definitions\n\nFind more information at https://github.com/tnozicka/opencompose.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}
			return errors.New("You have to specify a subcommand.")
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// We have to bind Viper in Run because there is only one instance to avoid collisions between subcommands
			cmdutil.BindViper(v, cmd.PersistentFlags(), Flag_Verbose_Key)

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(NewCmdConvert(v, out))
	rootCmd.AddCommand(NewCmdValidate(v, out))
	rootCmd.AddCommand(NewCmdVersion(v, out))
	rootCmd.AddCommand(NewCmdCompletion(v, out))

	return rootCmd
}
