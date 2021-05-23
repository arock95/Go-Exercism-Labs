// Package darts library provides a function to calculate a dart throw score
package darts

import "math"

func distanceFromCenter(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

// Score returns the score of a dart throw
func Score(x, y float64) int {
	switch {
	case distanceFromCenter(x, y) > 10:
		return 0
	case distanceFromCenter(x, y) > 5:
		return 1
	case distanceFromCenter(x, y) > 1:
		return 5
	default:
		return 10
	}
}
