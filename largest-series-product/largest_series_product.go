// Package lsproduct has functions to calculate some numbers and stuff
package lsproduct

import (
	"errors"
	"regexp"
)

var onlyNumbers *regexp.Regexp = regexp.MustCompile(`(?m)^[[:digit:]]+$`)

// LargestSeriesProduct returns what is the greater of a series of numbers in a sequence
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return 0, errors.New("span cannot be bigger than the digits string")
	}

	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	if !onlyNumbers.MatchString(digits) && digits != "" {
		return 0, errors.New("digits must be only numbers")
	}

	greater, n := 0, len(digits)
	for idx := 0; idx <= n-span; idx++ {
		current := 1
		for _, val := range digits[idx : idx+span] {
			current *= int(val - '0')
		}
		if current > greater {
			greater = current
		}
	}

	return greater, nil
}
