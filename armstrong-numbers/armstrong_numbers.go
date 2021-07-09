package armstrong

import (
	"math"
	"strconv"
)

// IsNumber returns if a given number is an Armstrong number
func IsNumber(input int) bool {
	asStr := strconv.Itoa(input)
	acc, length := 0, float64(len(asStr))
	for _, v := range asStr {
		acc += int(math.Pow(float64(v - 48), length))
	}
	return acc == input
}