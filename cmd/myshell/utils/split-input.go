package utils

import "strings"

func SplitInput(text string) (string, []string) {
  stringArray := strings.Split(text, " ")
  return stringArray[0], stringArray[1:]
}
