package sprint

import "fmt"

func Pairs() string {

	var result string

	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			result = result + fmt.Sprintf("%02d %02d, ", i, j)
		}
	}
	result = result[:len(result)-2]
	return result
}
