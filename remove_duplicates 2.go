package sprint

func RemoveDuplicates(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	seen := make(map[int]bool)
	var result []int

	for _, num := range arr {
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	return result
}
