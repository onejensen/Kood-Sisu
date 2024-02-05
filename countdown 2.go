package sprint

func Countdown(n int) string {
	var result string

	for i := n; i > 0; i = i - 2 {
		result += string('0'+i) + ", "
	}
	return result + "0!"
}
