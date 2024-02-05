package sprint

func GetLastRune(s string) rune {

	runes := []rune(s)

	if len(runes) > 0 {
		return runes[len(runes)-1]
	}
	return -1
}
