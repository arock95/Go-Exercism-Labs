// Package isogram provides function to test whether a string is an isogram
package isogram

import (
	"unicode"
)

// IsIsogram tests whether a string is an isogram
func IsIsogram(test string) bool {
	dict := make(map[rune]int)

	for _, c := range test {
		if unicode.IsLetter(c) {
			_, ok := dict[unicode.ToUpper(c)]

			if ok {
				return false
			}
			dict[unicode.ToUpper(c)] = 1
		}
	}
	return true
}

// func prepWord(word string) string {
// 	return strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(word, "-", ""), " ", ""))
// }

// // IsIsogram tests whether a string is an isogram
// func IsIsogram(test string) bool {
// 	test = prepWord(test)
// 	for i, v := range test {
// 		if strings.ContainsRune(test[i+1:], v) {
// 			return false
// 		}
// 	}
// 	return true
// }
