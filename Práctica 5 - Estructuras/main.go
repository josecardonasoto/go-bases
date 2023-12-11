package main

import "fmt"

// ---- EJERCICIO 1 ----

type Product struct {
	ID          uint
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d\n Name: %s\n Price: %f\n Description: %s\n Category: %s\n\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id uint) (product Product) {
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}
	return
}

// ---- EJERCICIO 2 ----

type Person struct {
	ID          uint
	Name        string
	DateOfBirth string
}

type Employee struct {
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Datos del empleado:\n"+
		" ID: %d\n Name: %s\n Date of birth: %s\n Position: %s\n\n", e.ID, e.Name, e.DateOfBirth, e.Position)
}

func main() {
	// ---- EJERCICIO 1 ----
	producto1 := Product{
		ID:          1,
		Name:        "Producto 1",
		Price:       100,
		Description: "Producto 1",
		Category:    "Categoria 1",
	}
	producto1.Save()
	producto2 := Product{
		ID:          2,
		Name:        "Producto 2",
		Price:       200,
		Description: "Producto 2",
		Category:    "Categoria 2",
	}
	producto2.Save()
	producto3 := Product{
		ID:          3,
		Name:        "Producto 3",
		Price:       300,
		Description: "Producto 3",
		Category:    "Categoria 3",
	}
	producto3.Save()
	producto1.GetAll()
	//fmt.Println(getById(2).Name)

	// ---- EJERCICIO 2 ----

	empleado1 := Employee{
		Position: "Cajero",
		Person: Person{
			ID:          1,
			Name:        "John",
			DateOfBirth: "01/01/2000",
		},
	}

	empleado1.PrintEmployee()

}
