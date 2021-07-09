// Package bob has functions to interact with a whimsical teenager
package bob

import (
	"regexp"
	"strings"
)

var yellQuestion *regexp.Regexp = regexp.MustCompile(`(?m)^[A-Z! .,'-]+\?$`)
var question *regexp.Regexp = regexp.MustCompile(`(?m).\?$`)
var yell *regexp.Regexp = regexp.MustCompile(`(?m)^[A-Z0-9 ,!%\*@#\$(*^]*$`)
var onlyNumbers *regexp.Regexp = regexp.MustCompile(`(?m)^[0-9 ,]+$`)
var nothing *regexp.Regexp = regexp.MustCompile(`(?m)^[[:blank:]]+$`)

// Hey returns what bob would say
func Hey(remark string) string {
	remark = strings.TrimSpace(strings.ReplaceAll(remark, "\n", " "))
	switch {
	case nothing.MatchString(remark) || len(remark) == 0:
		return "Fine. Be that way!"
	case yellQuestion.MatchString(remark):
		return "Calm down, I know what I'm doing!"
	case question.MatchString(remark):
		return "Sure."
	case yell.MatchString(remark) && !onlyNumbers.MatchString(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
