package functions

import (
	"fmt"
	"os"
)

func HandleCd(args []string) {
  command := args[0]
  if args[0] == "~" {
    command = os.Getenv("HOME")
  }
  if err := os.Chdir(command); err != nil {
    fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", command)
  }
}