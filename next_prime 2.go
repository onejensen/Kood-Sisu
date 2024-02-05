package sprint

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}

	if n%2 == 0 {
		n++
	}

	for !isPrime(n) {
		n += 2
	}

	return n
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	sqrt := int_Sqrt(number)

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func int_Sqrt(x int) int {
	guess := x / 2
	for guess*guess > x {
		guess = (guess + x/guess) / 2
	}
	return guess
}
