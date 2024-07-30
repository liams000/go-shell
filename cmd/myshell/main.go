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

func splitInput(text string) []string {
  stringArray := strings.Split(text, " ")
  return stringArray
}

func handleCommand(text []string) {
  if strings.EqualFold("exit", text[0]) {
    os.Exit(0)
  } else if strings.EqualFold("echo", text[0]){
    text = append(text[:0],text[0+1:]...)
    reformatted := strings.Join(text," ")
    fmt.Println(reformatted)
  } else {
    printUnknown(text[0])
  }
}

func main() {
  reader := bufio.NewScanner(os.Stdin)
  printPrompt()
  for reader.Scan() {
    text := cleanInput(reader.Text())
    splitText := splitInput(text)
    handleCommand(splitText)
    printPrompt()
  }
  fmt.Println()
}
