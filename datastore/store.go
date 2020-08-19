package datastore

import (
	"os"
	"slam-device-emulator/utils"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Server struct {
		Host       string   `yaml:"host"`
		Equipments []string `yaml:"equipments"`
	} `yaml:"server"`
}

// ReadConfigFile read configuration file
func ReadConfigFile() (cfg *Config) {
	f, err := os.Open("config.yml")
	utils.CheckError(err)
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	utils.CheckError(err)

	return cfg
}
