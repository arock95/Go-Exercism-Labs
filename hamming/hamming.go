// Package hamming provides a function to calculate the Hamming Distance of 2 strings
package hamming

import "fmt"

// Distance calculates the Hamming distance of 2 strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("two values %s and %s are not of the same length", a, b)
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
