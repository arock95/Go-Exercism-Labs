package collatzconjecture

import (
	"errors"
)

func processEven(n int) int {
	return n / 2
}

func processOdd(n int) int {
	return 3*n + 1
}

// CollatzConjecture determines number of steps using collatz conjecture to get to 1
func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return 0, errors.New("invalid number")
	}

	for n > 1 {
		if n%2 == 0 {
			n = processEven(n)
		} else {
			n = processOdd(n)
		}
		steps++
	}
	return steps, nil
}
