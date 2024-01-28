package main

func NoLetras(letra string) bool {

	for _, char := range letra {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			return false
		}
	}
	return true
}
