package cipher

import (
	"strings"
	"unicode"
)

// Cipher is generic interface inherited by all implementations
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// CeasarImp is the simple caesar implementation of Cipher
type CeasarImp struct {
}

// Encode ...
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

// Decode ...
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

// ShiftImp is the more customizable caesar implementation of Cipher
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

type VigenereImp struct {
	key string
}

func (v VigenereImp) Encode(pt string) string {
	keyLen := len(v.key)
	cipherText := strings.Builder{}
	counter := 0
	for _, l := range pt {
		if unicode.IsLetter(l) {
			shiftAmt := v.key[counter%keyLen] - 97
			newLetter := unicode.ToLower(l) + rune(shiftAmt)

			if newLetter > 122 {
				newLetter -= 26
			}
			cipherText.WriteRune(newLetter)
			counter++
		}
	}
	return cipherText.String()
}

func (v VigenereImp) Decode(ct string) string {
	keyLen := len(v.key)
	plainText := strings.Builder{}
	for i, l := range ct {
		if unicode.IsLetter(l) {
			keyIndex := i % keyLen
			shiftAmt := v.key[keyIndex] - 97
			newLetter := unicode.ToLower(l) - rune(shiftAmt)
			if newLetter < 97 {
				newLetter += 26
			}
			plainText.WriteRune(newLetter)
		}
	}

	return plainText.String()
}
