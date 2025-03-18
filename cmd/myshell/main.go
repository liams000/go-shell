package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/functions"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
	"os"
)

var commands = map[string]func([]string){
	"exit": functions.HandleExit,
	"echo": functions.HandleEcho,
	"type": functions.HandleType,
	"pwd":  functions.HandlePwd,
	"cd":   functions.HandleCd,
}

func handleCommand(command string, args []string) {
	if callback, exists := commands[command]; exists {
		callback(args)
		return
	}

	// Fallback function if none is found
	functions.HandleCommand(command, args)
}

func main() {
	// initialises a reader for use to read inputs from the cli
	reader := bufio.NewScanner(os.Stdin)
	// print a $
	utils.PrintPrompt()
	for reader.Scan() {
		// this cleans up any excess spaces in the input
		text := utils.CleanInput(reader.Text())
		// split the input into pieces we can ref where the 0 index is the command all the time
		command, args := utils.SplitInput(text)
		// self-explanatory
		handleCommand(command, args)
		utils.PrintPrompt()
	}
	fmt.Println()
}
