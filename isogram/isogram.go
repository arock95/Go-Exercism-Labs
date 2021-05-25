// Package isogram provides function to test whether a string is an isogram
package isogram

import "strings"

// prepWord removes hyphens and spaces from the word or phrase, and capitalizes it
func prepWord(word string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(word, "-", ""), " ", ""))
}

// IsIsogram tests whether a string is an isogram
func IsIsogram(test string) bool {
	test = prepWord(test)
	for i, v := range test {
		if strings.ContainsRune(test[i+1:], v) {
			return false
		}
	}
	return true
}
