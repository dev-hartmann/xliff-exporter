package utilities

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

type Config struct {
	Languages         []string
	ProjectName       string
	LocalizationsPath string
	ProjectPath       string
}

func NewConfigFromPath(path string) *Config {
	var config Config

	if path == "pwd" {
		wdString, err := os.Getwd()

		if err != nil {
			fmt.Errorf("Could not open current working directory, error:%v", err.Error())
		}

		path = strings.Join([]string{wdString, "test.toml"}, "/")
	} else {
		path = strings.Join([]string{path, "test.toml"}, "/")
	}

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil

	}

	return &config
}
