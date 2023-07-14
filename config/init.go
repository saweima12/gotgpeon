package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var config Configuration
var ignoreWordMap map[string]struct{} = make(map[string]struct{})
var textLang TextLang

// Get Configuration object.
func GetConfig() *Configuration {
	return &config
}

// Get Ignore wordmap.
func GetIgnoreWordMap() map[string]struct{} {
	return ignoreWordMap
}

func GetTextLang() *TextLang {
	return &textLang
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

	if len(config.IgnoreWord) > 0 {
		fmt.Println("Loading ignore words.")
		for _, word := range config.IgnoreWord {
			ignoreWordMap[word] = struct{}{}
		}

		fmt.Println("IgnoreWordMap: ", ignoreWordMap)
	}

	fmt.Println("Loading config finished")
	return nil
}

// Loading textlang data from langPath.
func InitTextlang(langPath string) error {
	f, err := os.ReadFile(langPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, &textLang)
	if err != nil {
		return err
	}

	fmt.Println("Loading language file [" + langPath + "] finished.")
	return nil
}
