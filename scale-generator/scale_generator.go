package scale

import (
	"fmt"
	"strings"
)

func Scale(tonic, interval string) []string {
	// figure out based on tonic if need to use sharp or flat scale
	tonicScale := getScale(tonic)

	tonic = strings.Title(tonic)
	// currentNoteInScale is numeric value of where note lies in the scale
	currentNoteInScale, err := findNoteInScale(tonic, tonicScale)
	if err != nil {
		panic(fmt.Sprintf("invalid tonic %s", tonic))
	}

	if interval == "" {
		interval = "mmmmmmmmmmm"
	}

	// process intervals
	returnScale := []string{tonic}

	for _, r := range interval {
		switch r {
		case 'm':
			newNote := getNextNote(currentNoteInScale, 1, tonicScale)
			if newNote != tonic {
				returnScale = append(returnScale, newNote)
			}
		case 'M':
			newNote := getNextNote(currentNoteInScale, 2, tonicScale)
			if newNote != tonic {
				returnScale = append(returnScale, newNote)
			}
		case 'A':
			newNote := getNextNote(currentNoteInScale, 3, tonicScale)
			if newNote != tonic {
				returnScale = append(returnScale, newNote)
			}
		}
		// switch currentnoteinscale to be the one we just appended
		currentNoteInScale, err = findNoteInScale(returnScale[len(returnScale)-1], tonicScale)
		if err != nil {
			panic(fmt.Sprintf("invalid note %s", tonic))
		}
	}
	fmt.Println(returnScale)
	return returnScale
}

func getScale(tonic string) []string {
	useSharps := []string{"G", "C", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "a"}

	for _, v := range useSharps {
		if tonic == v {
			return []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
		}
	}
	return []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
}

func getNextNote(currentNoteInScale, interval int, scale []string) string {
	nextNote := currentNoteInScale + interval
	if nextNote > len(scale)-1 {
		nextNote -= len(scale)
	}
	return scale[nextNote]
}

func findNoteInScale(note string, scale []string) (int, error) {
	for i, v := range scale {
		if note == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("note not found: %s", note)
}
