package cipher

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type CeasarImp struct {
}

//97 - 122
func (c CeasarImp) Encode(pt string) string {
	lettersOnly := strings.Builder{}
	for _, l := range pt {
		if unicode.IsLetter(l) {
			if newLetter := unicode.ToLower(l) + 3; newLetter > 122 {
				lettersOnly.WriteRune(newLetter - 26)
			} else {
				lettersOnly.WriteRune(newLetter)
			}

		}
	}
	return lettersOnly.String()
}

func (c CeasarImp) Decode(ct string) string {
	lettersOnly := strings.Builder{}

	for _, c := range ct {
		if newLetter := unicode.ToLower(c) - 3; newLetter < 97 {
			lettersOnly.WriteByte(byte(unicode.ToLower(c) + 23))
		} else {
			lettersOnly.WriteRune(newLetter)
		}
	}
	return lettersOnly.String()
}

type ShiftImp struct {
	shift int
}

func (s ShiftImp) Encode(pt string) string {
	lettersOnly := strings.Builder{}
	for _, l := range pt {
		if unicode.IsLetter(l) {
			if newLetter := unicode.ToLower(l) + rune(s.shift); newLetter > 122 {
				lettersOnly.WriteRune(newLetter - 26)
			} else if newLetter < 97 {
				lettersOnly.WriteByte(byte(unicode.ToLower(l) + 23))
			} else {
				lettersOnly.WriteRune(newLetter)
			}
		}
	}
	return lettersOnly.String()
}

func (s ShiftImp) Decode(ct string) string {
	lettersOnly := strings.Builder{}

	for _, c := range ct {
		if newLetter := unicode.ToLower(c) - rune(s.shift); newLetter < 97 {
			lettersOnly.WriteByte(byte(unicode.ToLower(c) + 23))
		} else if newLetter > 122 {
			lettersOnly.WriteRune(newLetter - 26)
		} else {
			lettersOnly.WriteRune(newLetter)
		}
	}
	return lettersOnly.String()
}
