package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
  reader := bufio.NewScanner(os.Stdin)
  printPrompt()
  for reader.Scan() {
    text := cleanInput(reader.Text())
    printUnknown(text)
    printPrompt()
  }
  fmt.Println()
}
