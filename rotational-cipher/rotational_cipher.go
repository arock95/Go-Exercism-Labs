package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher performs a Caesar cipher rotation on a string
func RotationalCipher(plaintext string, rot int) string {
	cipher := strings.Builder{}

	for _, r := range plaintext {
		if unicode.IsLetter(r) {
			newChar := rune(rot%26 + int(r))

			if unicode.IsLower(r) {
				if newChar > 'z' {
					newChar = rune(int(newChar - 'z' + 'a' - 1))
				}
				r = newChar
			}
			if unicode.IsUpper(r) {
				if newChar > 'Z' {
					newChar = rune(int(newChar - 'Z' + 'A' - 1))
				}
				r = newChar
			}
		}
		cipher.WriteRune(r)
	}
	return cipher.String()
}
