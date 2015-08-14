package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	languageSpecifier = flag.String("language", "all", "specify language to be exported. if specifier not provided, all will be exported")
	pathToProject     = flag.String("projectPath", "pwd", "path to project, when not specified uses current working directory")
	configFilePath    = flag.String("configPath", "pwd", "Path to config file, when not specified uses current working directory")
)

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func createAndExecuteCommand(config *Config, lang string) {

	cmd := exec.Command("xcodebuild", "-exportLocalizations", "-project", config.ProjectName, "-localizationPath", config.PathToExtractedLocalizationFiles, "-exportLanguage", lang)
	fmt.Println("-> cmd:", cmd)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
		printError(err)
	}

	if output != nil {
		fmt.Println("output")
		printOutput(output)
	}

}

func createArgFromStrings(argPrefix, argValue string) string {
	return fmt.Sprintf(argPrefix, argValue)
}

func init() {
	flag.Parse()
}

func main() {

	config := NewConfigFromPath(*configFilePath)

	if *pathToProject != "pwd" {
		fmt.Println(*pathToProject, "<- inside unequals pwd")
		pathProjectArg := createArgFromStrings("%q", *pathToProject)
		err := os.Chdir(pathProjectArg)

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
