package sprint

func StrLength(s string) []int {

	bytes := len(s)
	char := 0

	for range s {
		char++
	}
	return []int{char, bytes}
}
