package main

import "fmt"

func main() {
	//EJERCICIO 1
	fmt.Println("------- EJERCICIO 1 -------")
	palabra := "Manzana"
	fmt.Println("La palabra", palabra, "tiene", len(palabra), "letras\n")
	for i, leter := range palabra {
		fmt.Println("La letra número", i+1, "es", string(leter))
	}

	fmt.Println("")
	fmt.Println("")

	//EJERCICIO 2
	fmt.Println("------- EJERCICIO 2 -------")
	var (
		edad       = 22
		empleado   = true
		antiguedad = 1 / 12
		sueldo     = 300_000
	)

	switch {
	case edad <= 22:
		fmt.Println("Debes tener mas de 22 años para poder acceder al crédito")
	case empleado != true:
		fmt.Println("Debes estar empleado para poder acceder al crédito")
	case antiguedad < 1:
		fmt.Println("Debes tener mas de un año de antigüedad para acceder al crédito")
	case sueldo < 100_000:
		fmt.Println("Tienes acceso al crédito, no se te va a cobrar interés")
	default:
		fmt.Println("Tienes acceso al crédito, pero se te cobrará intereses")

	}

	fmt.Println("")
	fmt.Println("")

	//EJERCICIO 3
	fmt.Println("------- EJERCICIO 3 -------")

	var mes = 12

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	default:
		fmt.Println("Diciembre")
	}

	fmt.Println("")
	fmt.Println("")

	//EJERCICIO 4
	fmt.Println("------- EJERCICIO 4 -------")

	var employees = map[string]int{"Benja": 20, "Nata": 26, "Brenda": 19, "Darío": 44}
	mayores := 0
	for _, value := range employees {
		if value > 21 {
			mayores++
		}
	}
	fmt.Println("Mayores de 21:", mayores)
	fmt.Println("Edad de Benja", employees["Benja"])
	employees["Fede"] = 25
	delete(employees, "Brenda")
	fmt.Println(employees)

}
