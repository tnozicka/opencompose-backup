package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func NewOpenComposeCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	rootCmd := &cobra.Command{
		Use:   "opencompose",
		Short: "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) artifacts",
		Long:  "opencompose is a tool to transform OpenCompose files into Kubernetes (and OpenShift) definitions\n\nFind more information at https://github.com/tnozicka/opencompose.",
		Run:   runHelp,
		SilenceErrors: true,
		SilenceUsage: true,
	}

	rootCmd.PersistentFlags().Int8P("loglevel", "l", 0, "loglevel")

	rootCmd.AddCommand(NewCmdConvert(out))
	rootCmd.AddCommand(NewCmdValidate(out))
	rootCmd.AddCommand(NewCmdVersion(out))

	return rootCmd
}
