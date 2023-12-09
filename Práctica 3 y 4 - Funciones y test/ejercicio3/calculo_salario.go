package ejercicio3

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
