package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {
	// slice of pointers to Employee
	emps := []*Employee{
		{ID: 1, Name: "Amit"},
		{ID: 2, Name: "Priya"},
	}

	for _, e := range emps {
		fmt.Printf("ID: %d, Name: %s\n", e.ID, e.Name)
	}

	// Update via pointer
	emps[0].Name = "Amit Kumar"
	fmt.Println("After update:", emps[0])
}
