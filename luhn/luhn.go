package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// formatCheck determines whether the initial string given to Valid is valid number
// and returns the string with no spaces if valid
func formatCheck(num string) (bool, string) {
	num = strings.ReplaceAll(num, " ", "")
	if len(num) <= 1 {
		return false, ""
	}

	r := regexp.MustCompile("^[0-9]+$")
	if match := r.MatchString(num); !match {
		return false, ""
	}

	return true, num
}

// doubleEveryOther takes a string of numbers and returns a []int with every
// other number from the right doubled per the Luhn rules
func doubleEveryOther(num string) []int {
	numLen := len(num)
	doubledEveryOther := []int{}

	for i := numLen - 1; i >= 0; i-- {
		intConv, _ := strconv.Atoi(string(num[i]))
		
		if (numLen-i)%2 == 0 {
			intConv *= 2
			if intConv > 9 {
				intConv -= 9
			}
		}
		doubledEveryOther = append(doubledEveryOther, intConv)
	}
	return doubledEveryOther
}

// Valid takes a string and returns whether it is a valid Luhn number
func Valid(num string) bool {
	valid, formattedNum := formatCheck(num)
	if !valid {
		return false
	}

	doubledEveryOther := doubleEveryOther(formattedNum)

	sum := 0
	for _, v := range doubledEveryOther {
		sum += v
	}

	return sum%10 == 0
}
