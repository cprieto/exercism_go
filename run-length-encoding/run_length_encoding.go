package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func writeFreq(letter rune, freq int) string {
	var builder strings.Builder
	if freq > 1 {
		builder.WriteString(strconv.Itoa(freq))
	}
	builder.WriteRune(letter)
	return builder.String()
}

// RunLengthEncode uses RLE to encode a string
func RunLengthEncode(input string) string {
	var result strings.Builder
	var current rune
	var counter = 1
	for _, v := range input {
		switch {
		case current == 0:
			current = v
			continue
		case current == v:
			counter++
			continue
		}

		result.WriteString(writeFreq(current, counter))
		counter = 1
		current = v
	}

	// Last member is not written
	if current != 0 {
		result.WriteString(writeFreq(current, counter))
	}

	return result.String()
}

// RunLengthDecode uses RLE to decode a string
func RunLengthDecode(input string) string {
	var result strings.Builder
	var currentDigit strings.Builder

	for _, v := range input {
		if unicode.IsLetter(v) || unicode.IsSpace(v) {
			num, err := strconv.Atoi(currentDigit.String())
			if err != nil {
				num = 1
			}
			for i := 0; i < num; i++ {
				result.WriteRune(v)
			}
			currentDigit.Reset()
		}

		if unicode.IsNumber(v) {
			currentDigit.WriteRune(v)
		}
	}

	return result.String()
}