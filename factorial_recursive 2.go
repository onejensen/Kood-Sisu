package sprint

func FactorialRecursive(n int) int {

	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}
