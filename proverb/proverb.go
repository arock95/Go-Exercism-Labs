// Package proverb generates a proverb
package proverb

import "fmt"

// Proverb generates a proverb given a slice of strings
func Proverb(rhyme []string) []string {
	proverb := []string{}

	length := len(rhyme) - 1

	for i, v := range rhyme {
		if i == length {
			proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			return proverb
		}
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", v, rhyme[i+1]))
	}

	return proverb
}
