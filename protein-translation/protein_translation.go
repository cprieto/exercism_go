// Package protein contains functions and types for coding proteins
package protein

import (
	"errors"
	"regexp"
)

var validAminoacid = regexp.MustCompile(`(?mi)^[AUCG]{3}$`)
var slicerRegex = regexp.MustCompile(`(?mi)[AUCG]{1,3}`)

var (
	// ErrStop is raised when we finished the protein coding
	ErrStop = errors.New("stop protein coding")
	// ErrInvalidBase is raised when the sequence contains a unknown base
	ErrInvalidBase = errors.New("invalid base")
)

// FromCodon translate a codon RNA sequence into a protein
func FromCodon(input string) (string, error) {
	if !validAminoacid.MatchString(input) {
		return "", ErrInvalidBase
	}

	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA decodes a sequence of aminoacids into a sequence of proteins
func FromRNA(input string) ([]string, error) {
	if len(input) <= 3 {
		return nil, ErrInvalidBase
	}

	output := make([]string, 0)
	codons := slicerRegex.FindAllString(input, -1)
	for _, codon := range codons {
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return output, nil
			}
			return output, err
		}

		output = append(output, protein)
	}
	return output, nil
}
