package sprint

func ToRoman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid input"
	}

	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for i := 0; i < len(romanSymbols); i++ {
		count := num / romanValues[i]

		for j := 0; j < count; j++ {
			result += romanSymbols[i]
		}
		num -= count * romanValues[i]
	}
	return result
}
