package sprint

func StrToInt(s string) int {
	//check first to symbols
	result := 0
	neg := 1
	if s == "" {
		return result
	}
	if (s[0] == '+' || s[0] == '-') && (s[1] == '+' || s[1] == '-') {
		return 0
	}

	//check if it is negative and cut + or -
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = -1
		s = s[1:]
	}

	for _, number := range s {
		if number < '0' || number > '9' {
			return 0
		}
	}

	//Cut the zeros from the beginning
	for i := len(s); i > 0; i-- {
		char := s[0]
		if char == '0' {
			s = s[1:]
		} else {
			break
		}
	}
	//Converting slices one by one
	for i := len(s); i > 0; i-- {
		num := s[0] - '0'
		result = result*10 + int(num)
		s = s[1:]
	}

	return result * neg
}
