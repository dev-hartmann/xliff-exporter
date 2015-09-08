package utilities

import (
	"fmt"
	"os"
)

type XcodeCaller struct {
	Config *Config
}

func NewXcodeCaller(config *Config) *XcodeCaller {
	return &XcodeCaller{
		Config: config,
	}
}

func (xcc *XcodeCaller) ChangeToProjectDir(path string) (string, error) {

	if path != "pwd" {
		err := os.Chdir(path)

		if err != nil {
			fmt.Println(err.Error())
			return path, err
		}
	}

	return path, nil
}
