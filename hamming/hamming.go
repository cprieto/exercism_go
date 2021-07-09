// Package hamming implements the Exercism question
package hamming

import "errors"

// Distance calculates the hamming distance between
// two gene sequences, both sequences needs to have
// the same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("gene sequences need to be equal length")
	}
	distance := 0
	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			distance++
		}
	}
	return distance, nil
}
