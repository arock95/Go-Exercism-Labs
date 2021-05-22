package etl

import "strings"

// Transform puts the scrabble map into the new letter:value format
func Transform(original map[int][]string) map[string]int {
	scrabble := map[string]int{}

	for key, slice := range original {
		for _, value := range slice {
			scrabble[strings.ToLower(value)] = key
		}
	}
	return scrabble
}
