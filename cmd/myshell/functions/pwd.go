package functions

import (
	"fmt"
	"os"
)

func HandlePwd(args []string) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}
