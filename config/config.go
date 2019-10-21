package config

import (
	"log"
	"path"

	"github.com/riari/moni/util"
	"github.com/spf13/viper"
)

// Initialise sets up config using viper.
func Initialise() {
	home := util.GetUserDir()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(home, ".moni"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
