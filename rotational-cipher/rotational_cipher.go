// Package rotationalcipher exercism cipher stuff
package rotationalcipher

import (
	"strings"
	"unicode"
)

func mod(a, b rune) rune {
	return (a%b + b) % b
}

func rot(char rune, n rune, from rune) rune {
	idx := char - from
	return mod((idx+n), 26) + from
}

// RotationalCipher encrypts a string using a simple ROT cipher
func RotationalCipher(input string, shift int) string {
	var builder strings.Builder
	n := rune(shift)
	for _, char := range input {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				char = rot(char, n, 'A')
			} else {
				char = rot(char, n, 'a')
			}
		}
		builder.WriteRune(char)
	}

	return builder.String()
}
