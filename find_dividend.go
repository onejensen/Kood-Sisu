package sprint

/*
Write a Go function that takes three integers as input: from, to,
and divisor. The function should loop through the numbers from from
(inclusive) to to (exclusive) and return the first number in that
range that is divisible by the divisor. If there is no such number,
the function should return -1.
*/
func FindDividend(from, to, divisor int) int {

	for i := from; i < to; i++ {
		if i%divisor == 0 {
			return i
		}
	}
	return -1
}
