package sprint

func ReverseAlphabet(step int) string {

	n := step

	if n <= 0 {
		n = 1
	}

	var result string

	for i := 'z'; i >= 'a'; i-- {
		result += string(i)
		i -= rune(n - 1)
	}

	return result
}
