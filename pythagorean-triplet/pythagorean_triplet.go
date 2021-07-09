// Package pythagorean returns some Pythagorean algorithms
package pythagorean

import "math"

type Triplet [3]int

// Range return the Pythagorean triplets between start and end
func Range(start, end int) (result []Triplet) {
	for i := start; i <= end; i++ {
		for j := i + 1; j <= end; j++ {
			for k := j + 1; k <= end; k++ {
				a := math.Pow(float64(i), 2)
				b := math.Pow(float64(j), 2)
				c := math.Pow(float64(k), 2)
				if a+b == c {
					result = append(result, Triplet{int(i), (j), k})
				}
			}
		}
	}
	return
}

// Sum it returns the Pythagorean triplets summing sum
func Sum(sum int) (result []Triplet) {
	for i := 1; i < sum; i++ {
		for j := i + 1; j < sum; j++ {
			for k := j + 1; k < sum; k++ {
				a := math.Pow(float64(i), 2)
				b := math.Pow(float64(j), 2)
				c := math.Pow(float64(k), 2)
				if a+b == c && i+j+k == sum {
					result = append(result, Triplet{int(i), (j), k})
				}
			}
		}
	}
	return
}
