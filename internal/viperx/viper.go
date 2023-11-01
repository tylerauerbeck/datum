package viperx

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// MustBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func MustBindFlag(v *viper.Viper, name string, flag *pflag.Flag) {
	err := v.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}
