package config

import (
	"log"
	"os"
	"path"

	"github.com/riari/moni/util"
	"github.com/spf13/viper"
)

// Initialise sets up config using viper.
func Initialise() *viper.Viper {
	home := util.GetUserDir()
	configFile := "config.yml"
	configPath := path.Join(home, ".moni")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.Mkdir(configPath, os.ModePerm)
	}

	os.OpenFile(path.Join(configPath, configFile), os.O_RDWR|os.O_CREATE, 0644)

	config := viper.New()

	config.SetConfigType("yaml")
	config.SetConfigFile(path.Join(configPath, configFile))

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config
}
