package anagram

import (
	"reflect"
	"strings"
)

func extract(entry string) map[rune]int {
	result := make(map[rune]int)
	for _, v := range entry {
		result[v]++
	}
	return result
}

// Detect select the anagrams from a list of candidates
func Detect(entry string, options []string) (result []string) {
	entry = strings.ToLower(entry)
	letters := extract(entry)
	for idx, candidate := range options {
		candidate = strings.ToLower(candidate)
		if candidate != entry && reflect.DeepEqual(letters, extract(candidate)) {
			result = append(result, options[idx])
		}
	}
	return
}
