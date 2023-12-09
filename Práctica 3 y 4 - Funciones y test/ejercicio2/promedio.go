package ejercicio2

func promedioNotas(notas ...float64) (promedio float64) {
	suma := 0.0
	for _, nota := range notas {
		if nota > 0 {
			suma += nota
		}
	}
	promedio = suma / float64(len(notas))
	return
}
