package diffsquares

// SquareOfSum adds integers from 1 to the given number and returns the square
func SquareOfSum(num int) int {
	sum := (num * (num + 1)) / 2
	return sum * sum
}

// SumOfSquares squares each number from 1 to the given number and returns the sum
func SumOfSquares(num int) int {
	return (num * (num + 1) * (2*num + 1)) / 6
}

// Difference finds the difference between sumofsquare and squareofsum of a particular ranges of numbers
func Difference(num int) int {
	difference := SumOfSquares(num) - SquareOfSum(num)
	if difference < 0 {
		return difference * -1
	}
	return difference
}
