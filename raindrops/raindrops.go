package raindrops

import "fmt"

// Convert takes a number and translates it into rain sounds, if possible
func Convert(test int) string {
	ret := ""
	if test%3 == 0 {
		ret += "Pling"
	}
	if test%5 == 0 {
		ret += "Plang"
	}
	if test%7 == 0 {
		ret += "Plong"
	}
	if len(ret) == 0 {
		return fmt.Sprintf("%d", test)
	}

	return ret
}
