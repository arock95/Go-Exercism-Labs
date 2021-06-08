package cipher

import "math"


func NewCaesar() Cipher {
	return CeasarImp{}
}

func NewShift (shift int) Cipher {
	if math.Abs(float64(shift)) >= 26 || shift == 0 {
		return nil
	}
	return ShiftImp{
		shift: shift,
	}
}