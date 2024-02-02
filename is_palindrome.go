package sprint

func IsPalindrome(s string) bool {

	var cleanedStr []rune
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			char += 'a' - 'A'
			cleanedStr = append(cleanedStr, char)
		}
	}

	for i, j := 0, len(cleanedStr)-1; i < j; i, j = i+1, j-1 {
		if cleanedStr[i] != cleanedStr[j] {
			return false
		}
	}
	return true
}
