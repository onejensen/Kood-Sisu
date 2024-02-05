package sprint

import "fmt"

func StrCompress(input string) string {

	var result string

	for i := 0; i < len(input); {
		count := 0
		for j := i; j < len(input) && input[j] == input[i]; j++ {
			count++

		}
		if count > 1 {
			result += fmt.Sprintf("%0d", count)
		}

		result += string(input[i])
		if count == 0 {
			count = 1
		}
		i += count

	}
	return result
}
