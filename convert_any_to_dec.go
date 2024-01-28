package sprint

import "math"

func ConvertAnyToDec(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}

	result := 0
	baseLength := len(base)

	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		result += baseMap[char] * int(math.Pow(float64(baseLength), float64(len(s)-1-i)))
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
	}

	return len(base) == len(uniqueCharacters(base))
}

func uniqueCharacters(s string) string {
	result := ""
	seen := map[rune]bool{}

	for _, char := range s {
		if !seen[char] {
			result += string(char)
			seen[char] = true
		}
	}

	return result
}
