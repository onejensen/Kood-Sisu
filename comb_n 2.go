package sprint

func CombN(n int) []string {
	var result []string
	var currentCombination []rune
	var generate func(int)

	generate = func(currentPosition int) {
		if currentPosition == n {
			result = append(result, string(currentCombination))
			return
		}

		startDigit := '0'
		if currentPosition > 0 {
			startDigit = currentCombination[currentPosition-1] + 1
		}

		for digit := startDigit; digit <= '9'; digit++ {
			currentCombination = append(currentCombination, digit)
			generate(currentPosition + 1)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
	generate(0)
	return result
}
