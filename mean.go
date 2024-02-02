package sprint

func Mean(a, b, c float32) float32 {
	return (a + b + c) / 3
}

func Mean_() {
	num1, num2, num3 := 1.15, -2.0, 8.35

	result := Mean(float32(num1), float32(num2), float32(num3))

	println("Mean(a, b, c)\n", num1, num2, num3)
	println(">>\n", result)
}
