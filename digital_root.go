package sprint

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	total := 0
	for n > 0 {
		total += n % 10

		n /= 10
	}
	return DigitalRoot(total)
}
