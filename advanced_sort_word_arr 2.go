package sprint

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {

	for h := 0; h < len(a); h++ {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) == 1 {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}

/*func StrCompare(a, b string) int {

	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
*/
