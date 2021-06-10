package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func validateNum(num string) bool {
	if len(num) > 11 || len(num) < 10 {
		return false
	}
	if len(num) == 11 {
		if num[0] != '1' {
			return false
		}
		if num[1] <= '1' {
			return false
		}
		if num[4] <= '1' {
			return false
		}
	} else {
		if num[0] == '0' || num[0] == '1' {
			return false
		}
		if num[3] == '0' || num[3] == '1' {
			return false
		}
	}
	return true
}

// Number returns the 10 digit phone number
func Number(num string) (string, error) {
	numTest := regexp.MustCompile(`\D+`)
	validNum := numTest.ReplaceAllString(num, "")
	if !validateNum(validNum) {
		return "", errors.New("invalid number")
	}
	if len(validNum) == 11 {
		validNum = validNum[1:]
	}
	return validNum, nil
}

// AreaCode returns just the area code of a number
func AreaCode(num string) (string, error) {
	fullNum, err := Number(num)
	if err != nil {
		return "", err
	}
	return fullNum[0:3], err
}

// Format processes a phone number and formats it nicely
func Format(num string) (string, error) {
	fullNum, err := Number(num)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", fullNum[0:3], fullNum[3:6], fullNum[6:]), nil
}
