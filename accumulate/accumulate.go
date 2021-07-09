// Package accumulate is my answer to an Exercism question
package accumulate

// Accumulate maps an input string slice into another string slice
// using the operation function
func Accumulate(input []string, operation func(string) string) []string {
	output := make([]string, 0)

	for _, elem := range input {
		output = append(output, operation(elem))
	}

	return output
}
