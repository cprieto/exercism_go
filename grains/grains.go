package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains in a given square in a chess board
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Chess board has between 1 and 64 squares, inclusive")
	}

	return uint64(math.Pow(2, float64(n)-1)), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		n, _ := Square(i)
		total += n
	}
	return total
}
