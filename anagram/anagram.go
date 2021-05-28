// Package anagram includes function to detect anagrams
package anagram

import (
	"reflect"
	"strings"
)

// howMany builds a rune:#occurence map for each word
func howMany(word string) map[rune]int {
	stats := map[rune]int{}
	for _, v := range strings.ToLower(word) {
		_, err := stats[v]
		if !err {
			stats[v] = 0
		}
		stats[v]++
	}
	return stats
}

// Detect returns a list of anagrams
func Detect(word string, candidates []string) []string {
	wordStats := howMany(word)
	anagrams := []string{}

	for _, c := range candidates {
		if !strings.EqualFold(word, c) {
			cStats := howMany(c)
			if reflect.DeepEqual(wordStats, cStats) {
				anagrams = append(anagrams, c)
			}
		}
	}

	return anagrams
}
