package util

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Flag_File_Key = "file"
)

func AddIOFlags(v *viper.Viper, cmd *cobra.Command) {
	cmd.PersistentFlags().StringSliceP(Flag_File_Key, "f", []string{}, "Specify alternative OpenCompose file(s)")
	v.BindPFlag(Flag_File_Key, cmd.PersistentFlags().Lookup(Flag_File_Key))
}
