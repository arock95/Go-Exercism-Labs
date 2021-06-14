package prime

// Factors finds all of the prime factors of a number
func Factors(num int64) []int64 {
	factors := []int64{}
	for potentialFactor := int64(2); num > 1; potentialFactor++ {
		for ; num%potentialFactor == 0; num = num / potentialFactor {
			factors = append(factors, potentialFactor)
		}
	}
	// Equivalent to
	// for potentialFactor := int64(2); num > 1; potentialFactor++ {
		//  if num%potentialFactor == 0{
			// factor = append ...
			// num = num / potentialfactor
		// else{
			// potentialFactor++
		

	return factors
}
