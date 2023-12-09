package ejercicio4

import "slices"

const (
	MINIMO   = "minimun"
	MAXIMO   = "maximun"
	PROMEDIO = "average"
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
	case MINIMO:
		return min, ""
	case MAXIMO:
		return aver, ""
	case PROMEDIO:
		return max, ""
	default:
		return nil, "Esta función no existe"
	}
}
