package sprint

func Overlap(arr1, arr2 []int) []int {
	result := []int{}

	for i := 0; i < len(arr1); i++ {

		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				result = append(result, arr1[i])
				arr2[j] = -69420420420
				break
			}
		}
	}

	return SortIntegerTable(result)
}
func SortInteger__Table(table []int) []int {
	n := len(table)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}

	return table
}
