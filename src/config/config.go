package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	BasePath    string `json:"base_path"`
	LogFilename string `json:"log_filename"`
}

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}
