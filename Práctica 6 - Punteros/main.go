package main

import "fmt"

// ---- EJERCICIO 1 ----

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) Detalle() {
	fmt.Printf("Detalles del alumno\nNombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

// ---- EJERCICIO 2 ----

type Product interface {
	Price() float64
}

const (
	SMALL  = "Small"
	MEDIUM = "Medium"
	LARGE  = "Large"
)

type Small struct {
	Precio float64
}

func (s Small) Price() (precio float64) {
	precio = s.Precio
	return
}

type Medium struct {
	Precio float64
}

func (m Medium) Price() (precio float64) {
	precio = m.Precio + 0.3*m.Precio
	return
}

type Large struct {
	Precio float64
}

func (l Large) Price() (precio float64) {
	precio = l.Precio + 0.6*l.Precio + 2_500
	return
}

func Factory(tipo string, precio float64) Product {
	switch tipo {
	case SMALL:
		return &Small{Precio: precio}
	case MEDIUM:
		return Medium{Precio: precio}
	case LARGE:
		return Large{Precio: precio}
	default:
		return nil
	}
}

func main() {
	alumno1 := Alumno{
		Nombre:   "Juan",
		Apellido: "Perez",
		DNI:      12345678,
		Fecha:    "01/01/2000",
	}
	alumno1.Detalle()

	producto1 := Factory(SMALL, 4_000)
	fmt.Println(producto1.Price())
	producto2 := Factory(MEDIUM, 10_000)
	fmt.Println(producto2.Price())
	producto3 := Factory(LARGE, 1_000)
	fmt.Println(producto3.Price())

	r, ok := producto1.(Small)
	fmt.Println(r, ok)

}
