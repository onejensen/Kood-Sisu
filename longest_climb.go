package sprint

func LongestClimb(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var current []int
	var longest []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] > arr[i] {
			current = append(current, arr[i])
			if i == len(arr)-2 {
				current = append(current, arr[i+1])
			}
		} else {
			current = append(current, arr[i])
			if len(current) > len(longest) {
				longest = append([]int{}, current...)
			}
			current = nil
		}
	}
	return longest
}
