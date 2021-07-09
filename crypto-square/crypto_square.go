package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var validLetters *regexp.Regexp = regexp.MustCompile(`(?mi)[^a-z0-9]`)

func Encode(input string) string {
	input = validLetters.ReplaceAllString(strings.ToLower(input), "")
	length := len(input)

	num := math.Sqrt(float64(length))
	cols, rows := int(math.Ceil(num)), int(math.Floor(num))
	if rows*cols < length {
		rows = cols
	}

	if length < cols*rows {
		input += strings.Repeat(" ", cols*rows-length)
	}

	groups := make([]string, 0)
	for r := 0; r < rows; r++ {
		content := ""
		for c := 0; c < cols; c++ {
			idx := c*rows + r
			content += string(input[idx])
		}
		groups = append(groups, content)
	}

	return strings.Join(groups, " ")
}
