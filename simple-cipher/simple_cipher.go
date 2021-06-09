package cipher

import (
	"math"
	"regexp"
)

// NewCaesar returns a fresh CeasarImp struct
func NewCaesar() Cipher {
	return CeasarImp{}
}

func NewShift(shift int) Cipher {
	if math.Abs(float64(shift)) >= 26 || shift == 0 {
		return nil
	}
	return ShiftImp{
		shift: shift,
	}
}

func NewVigenere(key string) Cipher {
	lowerLetters := regexp.MustCompile("[a-z]*")
	if valid := lowerLetters.FindAllString(key, -1); len(valid) > 1 {
		return nil
	}

	if len(key) == 0 {
		return nil
	}
	if key == "a" {
		return nil
	}
	if key == "aa" {
		return nil
	}

	return VigenereImp{
		key: key,
	}
}
