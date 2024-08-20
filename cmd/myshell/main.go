package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/functions"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func handleCommand(command string, args []string) {
  if strings.EqualFold("exit", command) {
    functions.HandleExit(args)
  } else if strings.EqualFold("echo", command){
    functions.HandleEcho(args)
  } else if strings.EqualFold("type", command) {
    functions.HandleType(args)
  } else if strings.EqualFold("pwd", command) {
    functions.HandlePwd()
  } else if strings.EqualFold("cd", command) {
    functions.HandleCd(args)
  } else {
    functions.HandleCommand(command, args)
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
