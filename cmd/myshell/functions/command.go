package functions

import (
	"fmt"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func HandleCommand(command string, args []string) {
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		utils.PrintUnknown(command)
		return
	}

	outputStr := string(output)

	fmt.Print(outputStr)
}