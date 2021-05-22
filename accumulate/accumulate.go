package accumulate

// Accumulate takes a list of strings and performs a given convert action on them
func Accumulate(words []string, converter func(string) string) []string {
	var retList []string
	for _, word := range words {
		retList = append(retList, converter(word))
	}
	return retList
}
