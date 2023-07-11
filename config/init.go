package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var config Configuration

func GetConfig() *Configuration {
	return &config
}

// load configuration yaml file.
func InitConfig(configPath string) error {
	f, err := os.ReadFile(configPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return err
	}
	fmt.Println("Loading config finished")

	return nil
}
