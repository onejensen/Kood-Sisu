package sprint

func Accumulate(n int) int {

	var total int

	if n < 0 {
		return 0
	}
	for i := 0; i <= n; i++ {
		total += i
	}
	return total

}
