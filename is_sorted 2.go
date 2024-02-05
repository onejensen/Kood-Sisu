package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {

	direction := f(arr[0], arr[1])

	for i := 0; i < len(arr)-1; i++ {
		if f(arr[i], arr[i+1]) != direction && f(arr[i], arr[i+1]) != 0 {
			return false
		}
	}
	return true
}
func StrCompare(a, b string) int {

	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
