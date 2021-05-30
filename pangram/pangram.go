package pangram

import "strings"

var (
	letters = []string{"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w",
		"x", "y", "z"}
)

// IsPangram determines whether a phrase is a pangram
func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for _, letter := range letters {
		if !strings.Contains(sentence, string(letter)) {
			return false
		}
	}
	return true
}
