package main

import "fmt"

type Employee struct {
	Name   string
	ID     int
	Salary float64
}

func main() {
	e := Employee{
		Name:   "Rahul",
		ID:     101,
		Salary: 50000,
	}

	fmt.Printf("Name: %s, ID: %d, Salary: %.2f\n", e.Name, e.ID, e.Salary)
}
