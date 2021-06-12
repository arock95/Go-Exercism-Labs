package isbn

import (
	"unicode"
)

func isValidFormat(isbn string) ([]int, bool) {
	noSpecials := []int{}

	for i, v := range isbn {
		if unicode.IsLetter(v) {
			if i != len(isbn)-1 || v != 'X' {
				return noSpecials, false
			}
			noSpecials = append(noSpecials, 10)
		}
		if unicode.IsNumber(v) {
			noSpecials = append(noSpecials, int(v-'0'))
		}
	}

	if len(noSpecials) != 10 {
		return noSpecials, false
	}
	return noSpecials, true
}

// IsValidISBN takes a string and determines if it is a valid ISBN 10
func IsValidISBN(isbn string) bool {
	validFormat, isValid := isValidFormat(isbn)
	if !isValid {
		return false
	}

	total := 0
	for i, v := range validFormat {
		total += v * (10 - i)
	}
	return total%11 == 0
}
