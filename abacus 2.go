package sprint

func Abacus(a, b int) int {
	return a / b
}

func main() {
	num1 := 8
	num2 := 3

	result := Abacus(num1, num2)

	println("Abacus(", num1, ", ", num2, ")")
	println(">>", result)
}
