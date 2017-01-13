package util

import (
	"github.com/spf13/cobra"
)

func AddIOFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringSliceP("file", "f", []string{}, "Specify alternative OpenCompose file(s)")
}
