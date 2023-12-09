package ejercicio2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Esta función testea que si se obtenaga el promedio de las notas correctamente
func TestPromedioNotas(t *testing.T) {

	//Arrange - Se define el escenario de la prueba, donde tenemos un arreglo de notas positivas
	notas := []float64{5, 5, 5, 5, 4}
	resultadoEsperado := 4.8

	//Act - Se realiza la prueba
	resultadoObtenido := promedioNotas(notas...)

	//Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido, "El promedio no es el esperado, debería ser 5")
}
