package sprint

func BetweenLimits(from, to rune) string {

	result := ""

	if from > to {
		from, to = to, from
	}
	for i := int(from) + 1; i < int(to); i++ {
		result += string(i)
	}
	return result
}
