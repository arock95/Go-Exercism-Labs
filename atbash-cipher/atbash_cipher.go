package atbash

import (
	"strings"
	"unicode"
)

func formatCipher(cipher string) string {
	cipher = strings.ToLower(cipher)
	cipher = strings.ReplaceAll(cipher, ".", "")
	cipher = strings.ReplaceAll(cipher, ",", "")
	cipher = strings.ReplaceAll(cipher, " ", "")
	return cipher
}

// Atbash returns the atbash cipher result of a string
func Atbash(cipher string) string {
	output := strings.Builder{}

	for i, v := range formatCipher(cipher) {
		if i%5 == 0 && i != 0 {
			output.WriteRune(' ')
		}

		if unicode.IsLetter(v) {
			newChar := v + 25
			if newChar > 'z' {
				newChar = rune('z' - int(newChar-'z'))
			}
			output.WriteRune(newChar)
		}

		if unicode.IsNumber(v) {
			output.WriteRune(v)
		}
	}
	return output.String()

}
