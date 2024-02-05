package sprint

func BalancedParentheses(input string) bool {

	parenthesis := 0
	squared := 0
	fancyones := 0
	parenthesisB := 0
	squaredB := 0
	fancyonesB := 0

	for _, char := range input {
		if char == '(' {
			parenthesis++
		} else if char == ')' {
			parenthesisB++
		} else if char == '[' {
			squared++
		} else if char == ']' {
			squaredB++
		} else if char == '{' {
			fancyones++
		} else if char == '}' {
			fancyonesB++
		}
	}

	if parenthesis == parenthesisB && squared == squaredB && fancyones == fancyonesB {
		return true
	}
	return false
}
