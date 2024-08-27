package functions

import (
	"os"
	"strconv"
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