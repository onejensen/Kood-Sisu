package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	if from > to {
		from, to = to, from
	}

	switch {
	case from == to:
		return arr
	case from < 0 && to >= len(arr):
		return []float64{}
	case from < 0 && to < 0:
		return arr
	case from >= len(arr) && to >= len(arr):
		return arr
	case from < len(arr) && to >= len(arr):
		return arr[:from]
	case from < 0 && to < len(arr):
		return arr[to:]
	default:
		return append(arr[:from], arr[to:]...)
	}
}
