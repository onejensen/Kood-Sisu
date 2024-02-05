package sprint

func GetFirstRune(s string) rune {

	runes := []rune(s)

	if len(runes) > 0 {
		return runes[0]
	}
	return 0

}
