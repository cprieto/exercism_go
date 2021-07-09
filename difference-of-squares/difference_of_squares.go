// Package diffsquares has mathematic functions for squares
package diffsquares

import "math"

// SquareOfSum returns the square of all the numbers from 1 to n
func SquareOfSum(upTo int) int {
	result := 0.0
	for i := 1; i <= upTo; i++ {
		result += float64(i)
	}
	return int(math.Pow(result, 2))
}

// SumOfSquares returns the sum of all the squares of the numbers from 1 to n
func SumOfSquares(upTo int) int {
	result := 0.0
	for i := 0; i <= upTo; i++ {
		result += math.Pow(float64(i), 2)
	}

	return int(result)
}

// Difference returns the difference between the square of sum and sum of squares
func Difference(upTo int) int {
	return SquareOfSum(upTo) - SumOfSquares(upTo)
}
