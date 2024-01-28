package sprint

func SortIntegerTable(table []int) []int {
	//n es el lenght de table
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
