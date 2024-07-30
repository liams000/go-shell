package functions

import (
  "os"
  "strconv"
  "fmt"
  "strings"
)

func HandleExit(text []string) {
    if len(text) == 0 {
      os.Exit(0)
    }
    exitCode, err := strconv.Atoi(text[0])
    if err != nil {
        os.Exit(0)
    }
    os.Exit(exitCode)
}

func HandleEcho(text []string) {
    reformatted := strings.Join(text," ")
    fmt.Println(reformatted)
}

func HandleType(args []string) {
  if len(args) == 0 {
    fmt.Println("invalid args for type")
    return
  }

  builtin := map[string]bool{"echo": true, "exit": true, "type": true}

  if _, exists := builtin[args[0]]; exists {
    fmt.Printf("%s is a shell builtin\n", args[0])
  } else {
    fmt.Printf("%s: not found\n", args[0])
  }
}
