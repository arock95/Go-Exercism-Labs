package protein

import (
	"regexp"
)

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

const (
	ErrInvalidBase = CustomError("invalid base")
	ErrStop        = CustomError("stop received")
)

var (
	codonProtein = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine", "UUC": "Phenylalanine",
		"UUA": "Leucine", "UUG": "Leucine",
		"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
		"UAU": "Tyrosine", "UAC": "Tyrosine",
		"UGU": "Cysteine", "UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
	}
)

func FromCodon(input string) (string, error) {
	protein := codonProtein[input]
	if protein == "STOP" {
		return "", ErrStop
	}
	if protein == "" {
		return "", ErrInvalidBase
	}
	return protein, nil
}

func FromRNA(input string) ([]string, error) {
	proteins := []string{}

	threeLetters := regexp.MustCompile("[A-Z]{3}")
	RNA := threeLetters.FindAllString(input, -1)

	for _, v := range RNA {
		v, err := FromCodon(v)
		switch {
		case err == ErrInvalidBase:
			return proteins, err
		case err == ErrStop:
			return proteins, nil
		case err == nil:
			proteins = append(proteins, v)
		}
	}
	return proteins, nil
}
