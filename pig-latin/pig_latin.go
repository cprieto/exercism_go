package piglatin

import (
	"bufio"
	"regexp"
	"strings"
)

var vowels = regexp.MustCompile(`(?mi)^([aeiouy]|xr)`)
var digraphs = regexp.MustCompile(`(?mi)^(ch|qu|th|sm|st|gl|tr|fl|rh)`)
var trigraphs = regexp.MustCompile(`(?mi)^(squ|thr|sch|str)`)
var consonants = regexp.MustCompile(`(?mi)^y[aeiou]`)

func translate(input string) (result string) {
	word := []rune(input)
	switch {
	case trigraphs.MatchString(input):
		result = string(append(word[3:], word[:3]...))
	case digraphs.MatchString(input):
		result = string(append(word[2:], word[:2]...))
	case consonants.MatchString(input):
		result = string(append(word[1:], word[0]))
	case vowels.MatchString(input):
		result = input
	default:
		result = string(append(word[1:], word[0]))
	}

	result += "ay"
	return
}

// Sentence translates a sentence using pig latin
func Sentence(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	output := make([]string, 0)
	for scanner.Scan() {
		output = append(output, translate(scanner.Text()))
	}
	return strings.Join(output, " ")
}
