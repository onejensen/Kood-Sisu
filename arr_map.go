package sprint

func ArrMap(f func(int) bool, a []int) []bool {

	var result []bool

	for _, sli := range a {
		result = append(result, f(sli))
	}
	return result
}

func IsNegative(n int) bool {
	//Si el valor es mayor o igual a 0, retorna false
	if n >= 0 {
		return false
	} else { //De lo contrario retorna true
		return true
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	sqrt := intSqrt(n)

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func intSqrt(x int) int {
	guess := x / 2
	for guess*guess > x {
		guess = (guess + x/guess) / 2
	}
	return guess
}
