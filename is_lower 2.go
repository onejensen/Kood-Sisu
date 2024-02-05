package sprint

func IsLower(s string) bool {

	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
