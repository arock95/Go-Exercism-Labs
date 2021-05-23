// Package strand transcribes dna to rna
package strand

import "strings"

var (
	dnaToRna = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
)

// ToRNA takes a dna string and retuns the rna equivalent
func ToRNA(dna string) string {
	var sb strings.Builder
	for _, letter := range dna {
		sb.WriteRune(dnaToRna[letter])
	}
	return sb.String()
}
