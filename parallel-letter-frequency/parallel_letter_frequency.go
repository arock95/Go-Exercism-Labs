package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// unnecessary just use anonymous function
// func channelFrequency(s string, c chan FreqMap) {
// 	m := FreqMap{}
// 	for _, r := range s {
// 		m[r]++
// 	}
// 	c <- m
// }

func addMaps(master FreqMap, toAdd FreqMap) FreqMap {
	for v := range toAdd {
		master[v] += toAdd[v]
	}
	return master
}

// ConcurrentFrequency calculates the combined frequency maps
// of various strings and combines them
func ConcurrentFrequency(quotes []string) FreqMap {
	c := make(chan FreqMap, len(quotes))
	m := FreqMap{}

	for _, x := range quotes {
		//go channelFrequency(x, c)
		go func(s string, c chan FreqMap) {
			m := FreqMap{}
			for _, r := range s {
				m[r]++
			}
			c <- m
		}(x, c)
	}

	// can just use for range quotes{...}
	// for i := 0; i < len(quotes); i++ {
	// 	m = addMaps(m, <-c)
	// }
	for range quotes {
		m = addMaps(m, <-c)
	}

	return m
}
