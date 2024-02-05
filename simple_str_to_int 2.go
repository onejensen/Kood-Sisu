package sprint

func SimpleStrToInt(s string) int {

	result := 0

	for i := len(s); i > 0; i-- {
		char := s[0]
		if char == '0' {
			s = s[1:]
		} else {
			break
		}
	}

	for i := len(s); i > 0; i-- {
		num := s[0] - '0'
		result = result*10 + int(num)
		s = s[1:]
	}
	return result
}
