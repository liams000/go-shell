package functions

import (
	"fmt"
	"strings"
)

func HandleEcho(text []string) {
	reformatted := strings.Join(text, " ")
	fmt.Println(reformatted)
}