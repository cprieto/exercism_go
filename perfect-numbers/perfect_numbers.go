package perfect

import "errors"

func factors(n int64) (result []int64) {
	if n == 1 {
		return
	}

	result = append(result, 1)
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return
}

func sum(elements []int64) (result int64) {
	for _, v := range elements {
		result += v
	}
	return
}

// Classification is an aliquot type
type Classification int

const (
	// ClassificationAbundant Aliquot Abundant
	ClassificationAbundant Classification = iota
	// ClassificationDeficient Aliquot Deficient
	ClassificationDeficient
	// ClassificationPerfect Aliquot Perfect
	ClassificationPerfect
)

// ErrOnlyPositive Classify only works with positive numbers
var ErrOnlyPositive = errors.New("Only positive numbers are supported")

// Classify classifies a number based in its aliquot properties
func Classify(input int64) (Classification, error) {
	if input < 1 {
		return ClassificationPerfect, ErrOnlyPositive
	}

	aliquot := sum(factors(input))
	switch {
	case aliquot == input:
		return ClassificationPerfect, nil
	case aliquot > input:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
