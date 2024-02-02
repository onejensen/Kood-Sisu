package sprint

func LCM(a, b int) int {

	num1, num2 := a, b

	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return (a * b) / num1
}
