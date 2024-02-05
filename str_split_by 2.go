package sprint

func StrSplitBy(s, sep string) []string {

	start := 0
	result := []string{}

	if s == "" {
		return result
	}

	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
		}
	}
	result = append(result, s[start:])
	return result
}
