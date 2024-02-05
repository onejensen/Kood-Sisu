package sprint

func ToUpperCase(s string) string {

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = runes[i] - ('a' - 'A')
		}
	}
	resultado := string(runes)
	return resultado
}
