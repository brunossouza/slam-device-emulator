package datastore

import (
	"io/ioutil"
	"log"
	"os"
	"slam-device-emulator/utils"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
	Equipments []string `yaml:"equipments"`
}

// Tokens struct
type Tokens struct {
	Tokens []string `yaml:"tokens"`
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

// SaveTokensFile read configuration file
func SaveTokensFile(token string) {
	tokens := ReadTokens()

	tokens.Tokens = append(tokens.Tokens, token)

	d, err := yaml.Marshal(&tokens)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile("tokens.yml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadTokens read configuration file
func ReadTokens() (tokens *Tokens) {

	if !utils.FileExists("tokens.yml") {
		return &Tokens{}
	}

	f, err := os.Open("tokens.yml")
	utils.CheckError(err)
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&tokens)
	utils.CheckError(err)

	return tokens
}
