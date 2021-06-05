package grains

import (
	"errors"
)

const (
	TotalSquares = 64
)

// Square returns the number of grains on a specific square
func Square(input int) (uint64, error) {
	if input > 64 || input < 1 {
		return 0, errors.New("invalid square")
	}
	return 1 << (input - 1), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	return  1 << (TotalSquares) - 1
}