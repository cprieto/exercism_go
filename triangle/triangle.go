// Package triangle is the answer for the Exercism question
package triangle

import "math"

// Kind represents a type of triangle
type Kind int

// Represents a type of triangle
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides returns a type of triangle depending on the sides
func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c),
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT
	case a <= 0 || b <= 0 || c <= 0,
		a+b < c || b+c < a || a+c < b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
