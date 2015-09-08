package utilities

import (
	"fmt"
	"os"
	"os/exec"
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

func (xcc *XcodeCaller)ExportLanguages() (bool, error) {

	return true, nil
}

func createAndExecuteCommand(config *Config, lang string) {

	cmd := exec.Command("xcodebuild", "-exportLocalizations", "-project", config.ProjectName, "-localizationPath", config.P, "-exportLanguage", lang)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	err = os.Chdir(config.LocalizationsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

}