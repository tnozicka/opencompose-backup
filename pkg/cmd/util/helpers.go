package util

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

const (
	Flag_File_Key = "file"
)

func UsageError(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '%s -h' for help and examples.", msg, cmd.CommandPath())
}

func AddIOFlags(v *viper.Viper, cmd *cobra.Command) {
	cmd.PersistentFlags().StringSliceP(Flag_File_Key, "f", []string{}, "Specify alternative OpenCompose file(s)")
	v.BindPFlag(Flag_File_Key, cmd.PersistentFlags().Lookup(Flag_File_Key))
}
