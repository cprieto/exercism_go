package wordcount

import (
	"bufio"
	"regexp"
	"strings"
)

type Frequency map[string]int

var contractions = []string{
	"that's", "it's", "i'm", "you're", "they're", "we're", "she's", "he's", "won't", "don't", "wouldn't",
}

var re = regexp.MustCompile(`(?mi)[^a-z0-9]`)

func isContraction(word string) bool {
	for _, v := range contractions {
		if v == word {
			return true
		}
	}
	return false
}

func WordCount(input string) Frequency {
	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(input)))
	scanner.Split(bufio.ScanWords)

	result := make(map[string]int)
	for scanner.Scan() {
		text := scanner.Text()
		if !isContraction(text) {
			text = re.ReplaceAllString(text, "")
		}
		result[text]++
	}
	return result
}
