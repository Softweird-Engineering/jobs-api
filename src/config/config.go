package config

import (
	"github.com/spf13/viper"
	"sync"
	"strings"
	log "github.com/sirupsen/logrus"
)

var (
	once sync.Once
	instance Configuration
)

const config_filename = "config"

func Config(config_path string) Configuration {
	once.Do(func() {
		instance = loadConfig(config_path)
	})

	return instance
}

func loadConfig(path string) Configuration {
	var config_instance Configuration

	// Set the file name, path, type of the configurations file
	viper.SetConfigName(config_filename)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	// Getting evironment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"Error": err,
		}).Error("Error reading config file")
	}

	err := viper.Unmarshal(&config_instance)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err,
			}).Error("Unable to decode into struct")
	}

	return config_instance
}
