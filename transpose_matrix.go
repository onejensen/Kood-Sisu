package sprint

func TransposeMatrix(matrix [][]int) [][]int {

	row := len(matrix)
	col := len(matrix[0])

	result := make([][]int, col)
	for i := range result {
		result[i] = make([]int, row)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result

}
