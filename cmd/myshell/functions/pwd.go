package functions

import (
	"fmt"
	"os"
)

func HandlePwd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}