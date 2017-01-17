package util

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	flag "github.com/spf13/pflag"
)

const (
	Flag_File_Key = "file"
	Flag_OutputDir_Key = "output-dir"
)

func UsageError(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '%s -h' for help and examples.", msg, cmd.CommandPath())
}

func BindViperNames(v *viper.Viper, fs *flag.FlagSet, viperName string, cobraName string) {
	// errors here are mistakea in the code and cobra will panic in similar conditions; let's not handle it differently here right now

	flag := fs.Lookup(cobraName)
	if flag == nil {
		panic(fmt.Sprintf("Viper can't bind flag: %s", cobraName))
	}

	err := v.BindPFlag(viperName, flag)
	if err != nil {
		panic(err)
	}
}

func BindViper(v *viper.Viper, fs *flag.FlagSet, name string) {
	BindViperNames(v, fs, name, name)
}

func AddIOFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(Flag_OutputDir_Key, "o", "./", "Specify output directory for genrated Kubernetes (and OpenShift) definitions")
	cmd.PersistentFlags().StringSliceP(Flag_File_Key, "f", []string{}, "Specify alternative OpenCompose file(s)")
}

func AddIOFlagsViper(v *viper.Viper, cmd *cobra.Command) {
	BindViper(v, cmd.PersistentFlags(), Flag_OutputDir_Key)
	BindViper(v, cmd.PersistentFlags(), Flag_File_Key)
}
