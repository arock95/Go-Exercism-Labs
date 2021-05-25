// Package reverse provides a function to reverse a word
package reverse

// Reverse takes a word and reverses it
func Reverse(word string) string {
	boundaries := []int{}
	retme := ""

	// find the character byte boundaries (unicode chars are >1 byte)
	for index := range word {
		boundaries = append(boundaries, index)
	}

	// iterate through the boundaries in reverse and copy the characters into a new string
	endBoundary := len(word)
	for i := len(boundaries) - 1; i >= 0; i-- {
		retme += word[boundaries[i]:endBoundary]
		endBoundary = boundaries[i]
	}

	return retme
}
