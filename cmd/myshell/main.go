package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/functions"
)

func printPrompt() {
  fmt.Print("$ ")
}

func printUnknown(text string) {
  fmt.Printf("%s: command not found\n",text)
}

func cleanInput(text string) string {
  output := strings.TrimSpace(text)
  output = strings.ToLower(output)
  return output
}

func splitInput(text string) (string, []string) {
  stringArray := strings.Split(text, " ")
  return stringArray[0], stringArray[1:]
}

func handleCommand(command string, args []string) {
  if strings.EqualFold("exit", command) {
    functions.HandleExit(args)
  } else if strings.EqualFold("echo", command){
    functions.HandleEcho(args)
  } else if strings.EqualFold("type", command) {
    functions.HandleType(args)
  } else {
    printUnknown(command)
  }
}

func main() {
  reader := bufio.NewScanner(os.Stdin)
  printPrompt()
  for reader.Scan() {
    text := cleanInput(reader.Text())
    command, args := splitInput(text)
    handleCommand(command, args)
    printPrompt()
  }
  fmt.Println()
}
