package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a specific square
func Square(input int) (uint64, error) {
	if input > 64 || input < 1 {
		return 0, errors.New("invalid square")
	}
	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	total, _ := Square(64)
	return total*2 - 1
}
