package sprint

func BalancedParentheses(input string) bool {

	parenthesis := 0
	squared := 0
	fancyones := 0

	for _, char := range input {
		if char == '(' || char == ')' {
			parenthesis++
		} else if char == '[' || char == ']' {
			squared++
		} else if char == '{' || char == '}' {
			fancyones++
		}
	}
	if parenthesis%2 == 0 && squared%2 == 0 && fancyones%2 == 0 {
		return true
	}
	return false
}
