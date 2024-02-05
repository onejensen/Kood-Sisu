package sprint

func CountDivisible(from, to, step, divisor int) int {
	/*The function should check for exceptions and return 0 if any of the following
	conditions are met:step is negative or zero. divisor is zero. Otherwise, it should
	loop through the range of numbers from from (inclusive) to to (exclusive), checking every
	step-th element, and return the count of elements among these that are divisible
	by the divisor.*/
	if step <= 0 || divisor == 0 {
		return 0
	}
	var count int
	for i := from; i < to; i += step {
		if i%divisor == 0 {
			count++
		}
	}

	return count
}
