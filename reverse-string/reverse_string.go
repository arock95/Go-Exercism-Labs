// Package reverse provides a function to reverse a word
package reverse

// Reverse takes a word and reverses it
func Reverse(word string) string {
	retme := ""

	for _, v := range word {
		retme = string(v) + retme
	}

	return retme
}
