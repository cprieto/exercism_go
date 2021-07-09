package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

var vignereValid *regexp.Regexp = regexp.MustCompile(`(?m)[^a-z]|^(a)+$`)

// Cipher returns different cipher methods
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type caesar struct {
	n rune
}

type vignere struct {
	key []rune
}

// NOTE: real modulo operator, not remainder
func mod(a, b rune) rune {
	return (a%b + b) % b
}

func shift(char rune, n rune) rune {
	idx := char - 'a'
	return mod((idx+n), 26) + 'a'
}

func rotate(input string, n rune) string {
	if input == " " || input == "" {
		return input
	}

	var builder strings.Builder
	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			builder.WriteRune(shift(char, n))
		}
	}
	return builder.String()
}

// Encode a string using the default shift method
func (c caesar) Encode(input string) string {
	return rotate(input, c.n)
}

// Decode a string encoded with the shift method
func (c caesar) Decode(input string) string {
	return rotate(input, -c.n)
}

// Encode a string using the vignere method
func (v vignere) Encode(input string) string {
	var builder strings.Builder
	length := len(v.key)
	counter := 0
	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			idx := counter % length
			n := v.key[idx] - 'a'
			builder.WriteRune(shift(char, n))
			counter++
		}
	}
	return builder.String()
}

// Decode a string encoded with the vignere method
func (v vignere) Decode(input string) string {
	var builder strings.Builder
	length := len(v.key)
	counter := 0
	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			idx := counter % length
			n := 'a' - v.key[idx]
			builder.WriteRune(shift(char, n))
			counter++
		}
	}
	return builder.String()
}

// NewCaesar returns a rotation cipher
// using a rotation of 3
func NewCaesar() Cipher {
	return caesar{3}
}

// NewShift returns a rotation cipher
// using n rotation places
func NewShift(n int) Cipher {
	if n > 25 || n < -25 || n == 0 {
		return nil
	}
	return caesar{rune(n)}
}

// NewVigenere returns a rotation cipher
// using the vigenere method with a given key
func NewVigenere(key string) Cipher {
	if key == "" || vignereValid.MatchString(key) {
		return nil
	}

	return vignere{[]rune(key)}
}
