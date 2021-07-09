// Package summultiples does stuff with multiples
package summultiples

func any(elems []int, fn func(int) bool) bool {
	for _, v := range elems {
		if fn(v) {
			return true
		}
	}
	return false
}

// SumMultiples sum the multiples of divisors up to (not inclusive) limit
func SumMultiples(limit int, divisors ...int) (result int) {
	sum := 0
	for i := 1; i < limit; i++ {
		if any(divisors, func(x int) bool { return x != 0 && i%x == 0 }) {
			sum += i
		}
	}
	return sum
}
