package sprint

func AlphabetMastery(n int) string {

	var x string
	for i := int('a'); i < int('a')+n; i++ {
		x = x + string(i)
	}
	return x
}
