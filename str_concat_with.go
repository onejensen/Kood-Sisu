package sprint

func StrConcatWith(strs []string, sep string) string {

	var toConcat string

	for p, i := range strs {
		if p == 0 {
			toConcat += i
		} else {
			toConcat += sep + i
		}
	}
	return toConcat
}
