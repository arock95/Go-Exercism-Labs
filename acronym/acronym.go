// Package acronym provides an abbreviate function to abbreviate a phrase
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes a string and generates an abbreviation
func Abbreviate(s string) string {
	abbrev := strings.Builder{}
	
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")

	words := strings.Fields(s)
	for _, v := range words {
		abbrev.WriteRune(unicode.ToUpper(rune(v[0])))
	}

	return abbrev.String()
}
