package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Flag_LogLevel_Key = "loglevel"
)

func runHelp(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func NewOpenComposeCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	rootCmd := &cobra.Command{
		Use:           "opencompose",
		Short:         "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Long:          "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) definitions\n\nFind more information at https://github.com/tnozicka/opencompose.",
		RunE:          runHelp,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	v := viper.New()
	v.SetEnvPrefix("opencompose")
	v.AutomaticEnv()

	rootCmd.PersistentFlags().Int8P(Flag_LogLevel_Key, "l", 1, "set loglevel")
	v.BindPFlag(Flag_LogLevel_Key, rootCmd.PersistentFlags().Lookup(Flag_LogLevel_Key))

	rootCmd.AddCommand(NewCmdConvert(v, out))
	rootCmd.AddCommand(NewCmdValidate(v, out))
	rootCmd.AddCommand(NewCmdVersion(v, out))

	return rootCmd
}
