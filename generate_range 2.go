package sprint

func GenerateRange(min, max int) []int {

	var result []int

	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
