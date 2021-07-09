// Package raindrops is the answer to an Exercism question
package raindrops

import (
	"strconv"
)

// Convert generates the sound of the rain in an empty can
// depending on the number
func Convert(input int) string {
	var output string

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		output = strconv.Itoa(input)
	}

	return output
}
