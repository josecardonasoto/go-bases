package ejercicio1

func impuestoSalario(salario float64) (impuesto float64) {
	if salario > 150_000 {
		impuesto = salario * 0.27
	} else if salario > 50_000 {
		impuesto = salario * 0.17
	}
	return
}
