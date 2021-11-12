package romannumerals

import (
	"fmt"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3000 {
		return "", fmt.Errorf("invalid number: %d", input)
	}

	thousands, hundreds, tens, ones := breakdownInput(input)

	var romanOutput strings.Builder

	romanOutput.WriteString(strings.Repeat("M", thousands))
	romanOutput.WriteString(convertHundreds(hundreds))
	romanOutput.WriteString(convertTens(tens))
	romanOutput.WriteString(convertOnes(ones))

	return romanOutput.String(), nil
}

func breakdownInput(input int) (int, int, int, int) {

	thousands := input / 1000

	minusThousands := input - (thousands * 1000)
	hundreds := (minusThousands) / 100

	minusHundreds := minusThousands - (hundreds * 100)
	tens := (minusHundreds) / 10

	ones := minusHundreds - (tens * 10)

	return thousands, hundreds, tens, ones
}

func convertHundreds(hundreds int) string {
	switch hundreds {
	case 1, 2, 3:
		return (strings.Repeat("C", hundreds))
	case 4:
		return ("CD")
	case 5, 6, 7, 8:
		return ("D" + strings.Repeat("C", hundreds-5))
	case 9:
		return ("CM")
	}
	return ""
}

func convertTens(tens int) string {
	switch tens {
	case 1, 2, 3:
		return (strings.Repeat("X", tens))
	case 4:
		return ("XL")
	case 5, 6, 7, 8:
		return ("L" + strings.Repeat("X", tens-5))
	case 9:
		return ("XC")
	}
	return ""
}

func convertOnes(ones int) string {
	switch ones {
	case 1, 2, 3:
		return (strings.Repeat("I", ones))
	case 4:
		return ("IV")
	case 5, 6, 7, 8:
		return ("V" + strings.Repeat("I", ones-5))
	case 9:
		return ("IX")
	}
	return ""
}
