package sprint

func NRune(s string, i int) rune {

	runes := []rune(s)

	if i > len(s) {
		return 0
	}
	if len(runes) > 0 {
		return runes[i]
	}
	return runes[i]

}
