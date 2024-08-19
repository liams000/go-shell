package utils

import "strings"

func CleanInput(text string) string {
  output := strings.TrimSpace(text)
  // output = strings.ToLower(output)
  return output
}
