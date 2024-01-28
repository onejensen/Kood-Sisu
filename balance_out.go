package sprint

func BalanceOut(arr []bool) []bool {

	var trues int
	//Contar cuantos true hay en el array
	for _, item := range arr {
		if item == true {
			trues++
		}
	}
	//Contar cuantos false hay en el array
	falses := len(arr) - trues
	if trues > falses {
		equals := trues - falses
		//Loop para añadirlos al array
		for i := 0; i < equals; i++ {
			arr = append(arr, false)
		}
		//Contar cuantos trues hay en el array
	} else if trues < falses {
		equals := falses - trues
		//Loop para añadirlos al array
		for i := 0; i < equals; i++ {
			arr = append(arr, true)
		}
	}
	return arr
}
