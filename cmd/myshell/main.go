package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/functions"
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

func handleCommand(command string, args []string) {
  if strings.EqualFold("exit", command) {
    functions.HandleExit(args)
  } else if strings.EqualFold("echo", command){
    functions.HandleEcho(args)
  } else if strings.EqualFold("type", command) {
    functions.HandleType(args)
  } else {
    HandleCommand(command, args)
  }
}

func main() {
  reader := bufio.NewScanner(os.Stdin)
  utils.PrintPrompt()
  for reader.Scan() {
    text := utils.CleanInput(reader.Text())
    command, args := utils.SplitInput(text)
    handleCommand(command, args)
    utils.PrintPrompt()
  }
  fmt.Println()
}
