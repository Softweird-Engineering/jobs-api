package config

import (
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once     sync.Once
	instance Configuration
)

var (
	config_filename = os.Getenv("CONFIG_FILENAME")
	config_path     = os.Getenv("CONFIG_PATH")
)

func Config() Configuration {
	once.Do(func() {
		instance = loadConfig()
	})

	return instance
}

func loadConfig() Configuration {
	var config_instance Configuration

	// Set the file name, path, type of the configurations file
	viper.SetConfigName(config_filename)
	viper.AddConfigPath(config_path)
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
