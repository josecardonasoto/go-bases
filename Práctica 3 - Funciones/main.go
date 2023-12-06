package main

import (
	"fmt"
	"slices"
)

// ---- EJERCICIO 1 ----

func impuestoSalario(salario float64) (impuesto float64) {
	if salario > 150_000 {
		impuesto = salario * 0.27
	} else if salario > 50_000 {
		impuesto = salario * 0.17
	}
	return
}

// ---- EJERCICIO 2 ----

func promedioNotas(notas ...float64) (promedio float64) {
	suma := 0.0
	for _, nota := range notas {
		suma += nota
	}
	promedio = suma / float64(len(notas))
	return
}

// ---- EJERCICIO3 ----

func calculoSalario(minutos float64, categoria string) (salario float64) {
	var porcentaje float64
	switch categoria {
	case "C":
		salario = 1_000
	case "B":
		salario = 1_500
		porcentaje = 0.2
	case "A":
		salario = 3_000
		porcentaje = 0.5
	default:
		salario = 0
	}
	salario *= minutos / 60
	salario += salario * porcentaje
	return
}

// ---- EJERCICIO4 ----
const (
	minimo   = "minimun"
	maximo   = "maximun"
	promedio = "average"
)

func min(valores ...float64) (minimo float64) {
	minimo = slices.Min(valores)
	return
}

func max(valores ...float64) (maximo float64) {
	maximo = slices.Max(valores)
	return
}

func aver(valores ...float64) (promedio float64) {
	suma := 0.0
	for _, valor := range valores {
		suma += valor
	}
	promedio = suma / float64(len(valores))
	return
}
func operacion(funcion string) (func(...float64) float64, string) {
	switch funcion {
	case minimo:
		return min, ""
	case maximo:
		return aver, ""
	case promedio:
		return max, ""
	default:
		return nil, "Esta función no existe"
	}
}

func main() {
	// ---- EJERCICIO 1 ----
	var salario float64 = 100_000
	fmt.Println("\n---- EJERCICIO1 ----\n"+
		"El impuesto por el salario es de", impuestoSalario(salario), "$\n"+
		"El monto a pagar es de", salario-impuestoSalario(salario), "\n\n")

	// ---- EJERCICIO2 ----
	var notas = []float64{5, 4.6, 4.7, 4.9, 4.0}
	fmt.Println("---- EJERCICIO2 ----\n"+
		"El promedio para este semestre es de", promedioNotas(notas...), "\n\n")

	// ---- EJERCICIO3 ----
	var (
		minutos   float64 = 120
		categoria string  = "A"
	)
	fmt.Println("---- EJERCICIO3 ----\n"+
		"El salario de a pagar a este empleado es de", calculoSalario(minutos, categoria), "\n\n")

	// ---- EJERCICIO4 ----
	slice := []float64{1, 6, 4, 3, 5, 1}
	funcion, error := operacion("minimun")
	if funcion != nil {
		fmt.Println("---- EJERCICIO4 ----\n"+
			"El resultado de la función es", funcion(slice...), "\n\n")
	} else {
		fmt.Println("---- EJERCICIO4 ----\n" +
			error + "\n\n")
	}

}
