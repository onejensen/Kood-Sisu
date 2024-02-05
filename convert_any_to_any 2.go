package sprint

import "math"

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	decimalValue := ParseBase(nbr, baseFrom)
	result := ConvertFromDecimal(decimalValue, baseTo)
	return result
}

func ParseBase(s string, base string) int {
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

func ConvertFromDecimal(decimalValue int, baseTo string) string {
	if decimalValue == 0 {
		return string(baseTo[0])
	}
	result := ""
	baseLength := len(baseTo)

	for decimalValue > 0 {
		remainder := decimalValue % baseLength
		result = string(baseTo[remainder]) + result
		decimalValue /= baseLength
	}
	return result
}

func isValidBase2(base string) bool {
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

func uniqueCharacters2(s string) string {
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
