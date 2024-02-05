package sprint

func NthFibonacci(index int) int {

	if index < 0 {
		return -1
	}

	if index <= 1 {
		return index
	}

	return NthFibonacci(index-1) + NthFibonacci(index-2)

}
