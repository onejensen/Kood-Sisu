package sprint

func AlphaNumber(n int) string {

	if n == 0 {
		return "a"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	result := ""
	for n > 0 {
		digit := n % 10
		letter := rune('a' + digit)
		result = string(letter) + result
		n /= 10
	}
	return sign + result
}
