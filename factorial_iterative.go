package sprint

func FactorialIterative(n int) int {

	result := 1

	if n < 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		result *= i

	}
	return result
}
