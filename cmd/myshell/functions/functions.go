package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func HandlePwd() {
  dir, _ := os.Getwd()
  fmt.Println(dir)
}

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

  builtin := map[string]bool{"echo": true, "exit": true, "type": true, "pwd": true}

  if _, exists := builtin[args[0]]; exists {
    fmt.Printf("%s is a shell builtin\n", args[0])
  } else {
    env := os.Getenv("PATH")
    paths := strings.Split(env, ":")
    for _, path := range paths {
      exec := path + "/" + args[0]
      if _, err := os.Stat(exec); err == nil {
        fmt.Fprintf(os.Stdout, "%v is %v\n",args[0], exec)
        return
      }
    }
    fmt.Printf("%s: not found\n", args[0])
  }
}
