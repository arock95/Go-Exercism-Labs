package diffsquares

// SquareOfSum adds integers from 1 to the given number and returns the square
func SquareOfSum(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return total * total
}

// SumOfSquares squares each number from 1 to the given number and returns the sum
func SumOfSquares(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i * i
	}
	return total
}

// Difference finds the difference between sumofsquare and squareofsum of a particular ranges of numbers
func Difference(num int) int {
	difference := SumOfSquares(num) - SquareOfSum(num)
	if difference < 0 {
		return difference * -1
	}
	return difference
}
