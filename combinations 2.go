package sprint

import "fmt"

func Combinations() string {

	var result string

	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				num := fmt.Sprintf("%d%d%d", i, j, k)
				if result == "" {
					result += num
				} else {
					result += ", " + num
				}
			}
		}
	}
	return result
}
