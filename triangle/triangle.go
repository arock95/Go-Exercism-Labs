package triangle

import (
	"math"
)

// Kind represents the type of triangle
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = -1
	// Equ is an equilateral triangle
	Equ Kind = 3
	// Iso is an isosoles triangle
	Iso Kind = 2
	// Sca is a scalene triangle
	Sca Kind = 0
)

// isValidNumber checks a slice of numbers to see if they're negative, -Infinity, +Infinity, or NaN
func isValidNumber(nums []float64) bool {
	for _, v := range nums {
		if math.IsInf(v, 1) || math.IsInf(v, -1) || math.IsNaN(v) || v <= 0 {
			return false
		}
	}
	return true
}

// KindFromSides returns the type of triangle give 3 side values
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	switch {
	case a+b < c || a+c < b || b+c < a:
		k = NaT
	case !isValidNumber([]float64{a, b, c}):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	case a != b && b != c && a != c:
		k = Sca
	}

	return k
}
