package main

import (
	"fmt"
)

func main() {

	name := ""
	surname := ""
	age := ""
	var answer string

	fmt.Println("Escribe tu nombre:")
	fmt.Scanln(&name)
	fmt.Println("Escribe tu primer apellido:")
	fmt.Scanln(&surname)
	fmt.Println("¿Qué edad tienes?")
	fmt.Scanln(&age)

	fmt.Println("¿Así que te llamas", name, surname, "y tienes", age, "años? (si/no)")
	fmt.Scanln(&answer)
	if answer == "si" {
		fmt.Println("Qué cosas, ¿eh?", "\nHasta luego, cocodrilo")
	} else {
		fmt.Println("¿Entonces quién coño eres?")
		main()

	}

}
