package sprint

func ShiftBy(r rune, step int) rune {

	var rem int
	var let int
	rem = step % 26
	let = int(r)
	let = let + rem //shifted letter
	if 'a' <= let && let <= 'z' {
		return rune(let)
	} else {
		return rune('a' + let - 26) //
	}

}
