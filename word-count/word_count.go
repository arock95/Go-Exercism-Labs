package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a custom type to represent a map from string -> int
type Frequency map[string]int

// WordCount takes a phrase and returns the word count as a Frequency
func WordCount(phrase string) Frequency {
	wordCount := Frequency{}

	pattern := `[A-Za-z0-9]+('[a-zA-Z]+)*`
	r := regexp.MustCompile(pattern)
	words := r.FindAllString(phrase, -1)

	for _, v := range words {
		wordCount[strings.ToLower(v)] ++
	}

	return wordCount
}
