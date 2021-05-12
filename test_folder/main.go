package main

import "fmt"

func main() {
	e1 := Emp{ID: 100, Name: "Nui", Address: "Phrapradaeng", Salary: 10000}
	e1.PrintName()
	e1.PrintSalary()
}

// Emp user-defined type
type Emp struct {
	ID      int
	Name    string
	Address string
	Salary  float64
}

// Declare an Interface Type and methods does not have a body
type Employee interface {
	PrintName()    // Method with string return type
	PrintAddress() // Method with int parameter
	PrintSalary()  // Method with parameters and return type
}

// PrintName method to print employee name
func (e *Emp) PrintName() {
	fmt.Println("Employee Id: ", e.ID)
	fmt.Println("Employee Name: ", e.Name)
}

// PrintSalary method to calculate employee salary
func (e *Emp) PrintSalary() {
	fmt.Println("Employee Id: ", e.ID)
	fmt.Println("Employee Salary: ", e.Salary)
}
