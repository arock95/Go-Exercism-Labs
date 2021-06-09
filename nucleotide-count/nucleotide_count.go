package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

var (
	validNucleotide = []rune{'A', 'C', 'G', 'T'}
)

// Counts counts the number of occurrences of each nucleotide
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{}
	for _, v := range validNucleotide {
		h[v] = 0
	}

	for _, v := range d {
		if _, ok := h[v]; ok {
			h[v]++
		} else {
			return nil, errors.New("invalid nucleotide")
		}
	}
	return h, nil
}
