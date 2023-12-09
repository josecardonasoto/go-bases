package ejercicio1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Este test comprueba que se esté aplicando el impuesto correcto

func TestImpuestoSalario(t *testing.T) {

	//En este subtest se comprueba el impuesto cuando el salario es menor a 50_000
	t.Run("Menor a 50_000", func(t *testing.T) {
		//Arrange - Se define el escenario de la prueba, donde tenemos un sueldo menor a 50_000, en
		//este caso 30_000
		salario := 30_000.0
		resultadoEsperado := 0.0

		//Act - Se realiza la prueba
		resultadoObtenido := impuestoSalario(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "El impuesto no es el esperado, no debería haber impuesto")
	})

	//En este subtest se comprueba el impuesto cuando el salario es mayor a 50_000 y menor a 150_000
	t.Run("Mayor a 50_000 y menor a 150_000", func(t *testing.T) {
		//Arrange - Se define el escenario de la prueba, donde tenemos un sueldo mayor a 50_000 y menor a 150_000
		//en este caso 100_000
		salario := 100_000.0
		resultadoEsperado := salario * 0.17

		//Act - Se realiza la prueba
		resultadoObtenido := impuestoSalario(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "El impuesto no es el esperado, Debería ser del 0.17")
	})

	//En este subtest se comprueba el impuesto cuando el salario es mayor a 150_000
	t.Run("Mayor a 150_000", func(t *testing.T) {
		//Arrange - Se define el escenario de la prueba, donde tenemos un sueldo mayor a 150_000
		//en este caso 200_000
		salario := 200_000.0
		resultadoEsperado := salario * 0.27

		//Act - Se realiza la prueba
		resultadoObtenido := impuestoSalario(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "El impuesto no es el esperado, Debería ser del 0.27")
	})
}
