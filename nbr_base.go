package sprint

func NbrBase(n int, base string) string {

	if !isValidBase1(base) {
		return "NV"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	baseLength := len(base)
	result := ""

	for n > 0 {
		remainder := n % baseLength
		result = string(base[remainder]) + result
		n /= baseLength
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func isValidBase1(base string) bool {
	if len(base) < 2 {
		return false
	}

	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
	}

	return len(base) == len(uniqueCharacters1(base))
}

func uniqueCharacters1(s string) string {
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
