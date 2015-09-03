package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	utilities "xliff-exporter/Utilities"
)

var (
	languageSpecifier = flag.String("language", "all", "specify language to be exported. if specifier not provided, all will be exported")
	pathToProject     = flag.String("projectPath", "pwd", "path to project, when not specified uses current working directory")
	configFilePath    = flag.String("configPath", "pwd", "Path to config file, when not specified uses current working directory")
)

func createAndExecuteCommand(config *utilities.Config, lang string) {

	cmd := exec.Command("xcodebuild", "-exportLocalizations", "-project", config.ProjectName, "-localizationPath", config.PathToExtractedLocalizationFiles, "-exportLanguage", lang)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
		utilities.PrintError(err)
		return
	}

	if output != nil {
		fmt.Println("output")
		utilities.PrintOutput(output)
	}

	err = os.Chdir(config.PathToExtractedLocalizationFiles)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func init() {
	flag.Parse()
}

func main() {

	config := utilities.NewConfigFromPath(*configFilePath)

	if *pathToProject != "pwd" {
		err := os.Chdir(*pathToProject)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(*pathToProject)

	if *languageSpecifier == "all" {
		for _, lang := range config.Languages {
			createAndExecuteCommand(config, lang)
		}
	} else {
		createAndExecuteCommand(config, *languageSpecifier)
	}
}
