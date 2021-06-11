package prime

import (
	"math"
)

func isPrime(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

// Nth returns the 'nth' prime number
func Nth(nth int) (int, bool) {
	if nth == 0 {
		return 0, false
	}
	counter := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			counter++
			if counter == nth {
				return i, true
			}
		}
	}

}
