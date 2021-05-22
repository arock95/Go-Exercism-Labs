// Package twofer provides a function to return two-for-one string
package twofer

import "fmt"

// ShareWith returns a two-for-one string with a user's name if provided
// Or a generic two-for-one string if not
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
