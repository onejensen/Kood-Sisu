package sprint

func Sqrt(n int) int {
	return guessSqrt(0, n)
}

func guessSqrt(guess, n int) int {
	diff := n - (guess * guess)
	if diff < 0 {
		return 0
	} else if diff == 0 {
		return guess
	}
	return guessSqrt(guess+1, n)
}
