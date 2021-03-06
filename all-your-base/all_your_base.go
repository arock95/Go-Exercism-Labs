package allyourbase

import (
	"errors"
)

func intPow(base, pow int) int {
	total := 1
	for i := 0; i < pow; i++ {
		total *= base
	}
	return total
}

func toBaseTen(inputBase int, inputDigits []int) int {
	total := 0
	numInputDigits := len(inputDigits) - 1
	for i := numInputDigits; i >= 0; i-- {
		total += inputDigits[numInputDigits-i] * intPow(inputBase, i)
	}
	return total
}

func fromBaseTenToOutputBase(outputBase, baseTenValue int) []int {
	outputDigits := []int{}
	divisionResult := baseTenValue

	for divisionResult != 0 {
		outputDigits = append([]int{divisionResult%outputBase}, outputDigits... )
		divisionResult = divisionResult / outputBase
	}
	return outputDigits
}

// ConvertToBase takes an integer slice and in/output bases and converts values to output base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int(nil), errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int(nil), errors.New("output base must be >= 2")
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return []int(nil), errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	bTen := toBaseTen(inputBase, inputDigits)
	ret := fromBaseTenToOutputBase(outputBase, bTen)

	if len(ret) == 0 {
		return []int{0}, nil
	}

	return ret, nil
}
