package main

import (
	"fmt"
)

// Employee struct
type Employee struct {
	Name    string
	PassKey string
}

// CheckPassKey method: passkey verify karta hai
func (e Employee) CheckPassKey(inputKey string) bool {
	return e.PassKey == inputKey
}

func main() {
	// Employees ki list (hardcoded)
	employees := []Employee{
		{"Rahul", "RAH123"},
		{"Priya", "PRI456"},
		{"Amit", "AMI789"},
	}

	var inputKey string
	fmt.Print("Enter your passkey: ")
	fmt.Scan(&inputKey) // user se passkey input

	found := false

	for _, emp := range employees {
		if emp.CheckPassKey(inputKey) {
			found = true
			fmt.Println("Access granted. Welcome,", emp.Name)
			break
		}
	}

	if !found {
		fmt.Println("Wrong passkey. Access denied.")
	}
}
