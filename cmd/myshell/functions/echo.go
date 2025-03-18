package functions

import (
	"fmt"
	"strings"
)

func HandleEcho(text []string) {
	reformatted := strings.Join(text, " ")
	reformatted = strings.ReplaceAll(reformatted, "'", "")
	fmt.Println(reformatted)
}
