package utilities

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
)

func GitInit(path string) bool {

	return executeGitCommand(path, "init")
}

func GitAdd(path string) bool {
	return executeGitCommand(path, "add", ".")
}

func GitCommit(path, commitMessage string) bool {
	return executeGitCommand(path, "commit", "-m", fmt.Sprintf(`"%s"`,commitMessage))
}

func GitPush(path, origin, branch string) bool {
	return executeGitCommand(path, "push", origin, branch)
}
func executeGitCommand(path string, params ... string) bool {

	argString := strings.Join(params, " ")

	err := os.Chdir(path)
	if err != nil {
		PrintError(err)
		return false
	}

	cmd := exec.Command("git",argString)
	output, err := cmd.CombinedOutput()

	if err != nil {
		PrintError(err)
		return false
	}

	if output != nil {
		PrintOutput(output)
	}

	return true
}