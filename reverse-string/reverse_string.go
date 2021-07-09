// Package reverse has functions for reversing strings
package reverse

import "strings"

// Reverse a given string
func Reverse(input string) string {
	if len(input) == 0 {
		return ""
	}

	inRunes := []rune(input)
	var builder strings.Builder
	for idx := len(inRunes) - 1; idx >= 0; idx-- {
		builder.WriteRune(inRunes[idx])
	}

	return builder.String()
}
