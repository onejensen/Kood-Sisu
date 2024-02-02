package sprint

func FromRoman(s string) int {
	result := 0
	roman_nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last_num := 0

	for _, rom_num := range s {
		num := roman_nums[rom_num]
		if num > last_num {

			num -= last_num * 2
		}
		result += num
		last_num = num
	}
	return result
}
