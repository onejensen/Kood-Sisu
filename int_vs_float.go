package sprint

func IntVsFloat(i int, f float32) string {

	ii := float32(i)

	if ii > f {
		return "Integer"
	} else if ii < f {
		return "Float"
	} else {
		return "Same"
	}

}
