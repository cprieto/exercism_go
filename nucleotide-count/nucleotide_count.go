// Package dna has functions to process DNA sequences
package dna

import (
	"fmt"
)

// Histogram is the aminoacids in a DNA strand
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts returns the account of aminoacids in a DNA sequence
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{
		'A': 0,
		'C': 0,
		'T': 0,
		'G': 0,
	}

	for _, e := range d {
		if _, ok := h[e]; !ok {
			return Histogram{}, fmt.Errorf("invalid aminoacid: %v", e)
		}
		h[e]++
	}

	return h, nil
}
