// Package pangram has functions to deal with pangrams
package pangram

import (
	"strings"
	"unicode"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

// IsPangram returns if a word is a pangram or not
func IsPangram(input string) bool {
	current := make(map[rune]bool)

	for _, letter := range strings.ToLower(input) {
		if _, ok := current[letter]; !ok && unicode.IsLetter(letter) {
			current[letter] = true
		}
	}

	for _, v := range abc {
		if _, ok := current[v]; !ok {
			return false
		}
	}

	return true
}
