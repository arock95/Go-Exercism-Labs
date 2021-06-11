package sieve

func contains(num int, composites []int) bool {
	for _, v := range composites {
		if num == v {
			return true
		}
	}
	return false
}

// Sieve implements the Sieve of Eratosthenes to find primes
func Sieve(limit int) []int {
	composites := []int{}
	primes := []int{}

	for i := 2; i <= limit; i++ {
		if !contains(i, composites) {
			primes = append(primes, i)
			for multiple := i * 2; multiple <= limit; multiple += i {
				composites = append(composites, multiple)
			}
		}
	}
	return primes
}
