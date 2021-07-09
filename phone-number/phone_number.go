package phonenumber

import "regexp"

var cleaner *regexp.Regexp = regexp.MustCompile(``)

func Number(input string) (string, error) {
	input = cleaner.ReplaceAllString(input, "")
}