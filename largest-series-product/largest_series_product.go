package lsproduct

import (
	"errors"
	"unicode"
)

func multiply(nums []int) int {
	total := 1
	for _, v := range nums {
		total *= v
	}
	return total
}

// LargestSeriesProduct takes a string of digits, calculate the 
// largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("bad span")
	}

	nums := []int{}
	for _, v := range digits {
		if !unicode.IsNumber(v) {
			return 0, errors.New("bad number")
		}
		nums = append(nums, int(v-'0'))
	}

	largest := 0
	for i := 0; i <= len(digits)-span; i++ {
		if test := multiply(nums[i : i+span]); test > largest {
			largest = test
		}
	}
	return largest, nil
}
