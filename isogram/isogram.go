// Package isogram handles
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns if a word does not have repeated characters
func IsIsogram(input string) bool {
	found := make(map[rune]bool)
	for _, letter := range strings.ToUpper(input) {
		if unicode.IsLetter(letter) {
			if _, ok := found[letter]; ok {
				return false
			}
			found[letter] = true
		}
	}

	return true
}
