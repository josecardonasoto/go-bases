package main

import (
	"errors"
	"fmt"
)

// ---- EJERCICIO 1 ----
type ErrSalaryMinimum struct{}

func (e ErrSalaryMinimum) Error() string {
	return "the salary entered does not reach the taxable minimum\n"
}

// ---- EJERCICIO 2 ----
type ErrSalaryLess struct{}

func (e ErrSalaryLess) Error() string {
	return "salary is less than 100_000\n"

}

// ---- EJERCICIO 3 ----
var (
	ErrSalaryLessThan100_000 = errors.New("salary is less than 100_000")
	ErrSalaryLessThan150_000 = errors.New("the minimum taxable amount is 150_000")
)

func MustPayTax(salary int) (err error) {
	if salary < 100_000 {
		err = ErrSalaryLess{}
		fmt.Println(err)
		return
	} else if salary < 150_000 {
		err = ErrSalaryMinimum{}
		return
	}
	return
}

func main() {
	var salary int = 140_000
	if salary < 100_000 {
		fmt.Println(ErrSalaryLessThan100_000)
	} else if salary < 150_000 {
		err := fmt.Errorf("Error: %w and the salary entered is %d", ErrSalaryLessThan150_000, salary)
		fmt.Println(err)
	}
	fmt.Println("You must pay taxes")

}
