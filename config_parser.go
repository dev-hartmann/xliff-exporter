package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Languages                        []string
	ProjectName                      string
	PathToExtractedLocalizationFiles string
}

func NewConfigFromPath(path string) *Config {
	var config Config

	if path == "pwd" {
		wdString, err := os.Getwd()

		if err != nil {
			//Logger.Fatal(err.Error())
		}

		path = fmt.Sprint(wdString, "/test.toml")
		//utilities.Logger.Info("Using:",path," as config file")
	}

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil

	}

	return &config
}
