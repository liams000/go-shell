package utils

import "strings"

func isEnclosedQuotes(s string) bool {
	return len(s) >= 2 && s[0] == '\'' && s[len(s)-1] == '\''
}

func SplitInput(text string) (string, []string) {
	stringArray := strings.Split(text, " ")
	return stringArray[0], stringArray[1:]
}
