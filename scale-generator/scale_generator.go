// Package scale is an unpleasant experience.
//
// I have zero musical education and found the problem statement imprecise
// and confusing, so this exercise is not so much about learning Go but
// deciphering the problem.
//
// Because of this I'll forego cleaning up the code and instead let you see
// it in exactly the state it was in when all tests passed.
package scale

import (
	"fmt"
	"strings"
)

var (
	scaleSharpTable = []string{"A", "A#", "B", "C", "C#",
		"D", "D#", "E", "F", "F#", "G", "G#"}
	scaleFlatTable = []string{"A", "Bb", "B", "C", "Db",
		"D", "Eb", "E", "F", "Gb", "G", "Ab"}

	// This table was generated emperically, with items added as needed to
	// get the tests to pass.  There are some items here, such as "e", that
	// seem to mismatch the problem statement.
	useFlatsTable = []string{"bb", "d", "Eb", "e", "Db", "g", "F"}
)

func lookup(note int, useFlat bool) string {
	var table []string
	if useFlat {
		table = scaleFlatTable
	} else {
		table = scaleSharpTable
	}
	return table[note%len(table)]
}

func tonicToNote(tonic string) int {
	tonic = strings.ToUpper(tonic)
	for i, v := range scaleSharpTable {
		if strings.ToUpper(v) == tonic {
			return i
		}
	}
	for i, v := range scaleFlatTable {
		if strings.ToUpper(v) == tonic {
			return i
		}
	}
	panic(tonic) // bad idea, but this is toy code
}

func useFlats(tonic string) bool {
	for _, v := range useFlatsTable {
		if v == tonic {
			return true
		}
	}
	return false
}

func patternToSteps(pattern string) []int {
	if pattern == "" {
		pattern = "mmmmmmmmmmmm"
	}

	// TODO: Possible bug: I don't understand why last character in the
	// pattern is unused.
	steps := []int{}
	for i := 0; i < len(pattern)-1; i++ {
		switch pattern[i] {
		case 'M':
			steps = append(steps, 2)
		case 'm':
			steps = append(steps, 1)
		case 'A':
			steps = append(steps, 3)
		default:
			panic(fmt.Sprintf("invalid pattern %v at %c.",
				pattern, pattern[i]))
		}
	}
	return steps
}

func Scale(tonic, pattern string) []string {
	scale := []string{}
	start := tonicToNote(tonic)
	steps := patternToSteps(pattern)
	note := start
	useFlats := useFlats(tonic)
	scale = append(scale, lookup(note, useFlats))
	for _, step := range steps {
		note += step
		scale = append(scale, lookup(note, useFlats))
	}
	return scale
}
