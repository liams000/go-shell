package functions

import (
	"fmt"
	"os"
	"strings"
)

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
				fmt.Fprintf(os.Stdout, "%v is %v\n", args[0], exec)
				return
			}
		}
		fmt.Printf("%s: not found\n", args[0])
	}
}