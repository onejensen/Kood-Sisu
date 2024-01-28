package sprint

func SplitWhitespaces(s string) []string {

	word := ""
	words := []string{}

	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			words = append(words, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	words = append(words, word)
	return words
}
