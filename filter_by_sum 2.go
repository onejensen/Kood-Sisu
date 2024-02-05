package sprint

func FilterBySum(arr [][]int, limit int) [][]int {

	result := [][]int{}
	//Loop para averiguar que hay dentro de los arrays
	for _, x := range arr {
		sum := 0
		for _, y := range x {
			sum += y
		}
		//Si 'sum' es igual o mayor
		if sum >= limit {
			result = append(result, x)
		}
	}
	return result
}
