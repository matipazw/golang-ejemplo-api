package config

import (
	"log"

	"github.com/spf13/viper"
)

var Provider = func() *Configuration {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &configuration
}
