package sprint

func ToCapitalCase(s string) string {
	result := ""
	for i, letter := range s {
		switch true {
		// Primera letra minuscula
		case (i == 0 && isRuneLower(letter)):
			result += string(letter - ('a' - 'A'))
		// Primera letra NO es minuscula, el caracter anterior no es alfanumerico
		case (i > 0 && !(isAlphaNumeric(rune(s[i-1]))) && isRuneLower(letter)):
			result += string(letter - ('a' - 'A'))
		// La primera letra NO es mayuscula
		case i > 0 && isRuneUpper(letter) && isAlphaNumeric(rune(s[i-1])):
			result += string(letter - ('A' - 'a'))
		// Cualquier otra cosa
		default:
			result += string(letter)
		}
	}
	return result
}

func isRuneLower(char rune) bool {

	if char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

func isRuneNumber(char rune) bool {

	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isRuneUpper(char rune) bool {

	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}

func isAlphaNumeric(r rune) bool {

	return isRuneLower(r) || isRuneUpper(r) || isRuneNumber(r)
}
