// Package darts has functions to handle the game of darts
package darts

import "math"

// Score returns the score of a dart position
func Score(x, y float64) int {
	dist := math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))
	if dist > 10 {
		return 0
	} else if dist > 5 {
		return 1
	} else if dist > 1 {
		return 5
	} else {
		return 10
	}
}
