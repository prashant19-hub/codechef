package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
	Position   string
	Salary     float64
}

var employees []Employee
var nextID int = 1
var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println("\n===== HR Management System =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View All Employees")
		fmt.Println("3. Search Employee by ID")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choice := readInt()

		switch choice {
		case 1:
			addEmployee()
		case 2:
			viewEmployees()
		case 3:
			searchEmployee()
		case 4:
			updateEmployee()
		case 5:
			deleteEmployee()
		case 6:
			fmt.Println("Exiting HR Management System...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addEmployee() {
	fmt.Println("\n--- Add Employee ---")

	fmt.Print("Enter Name: ")
	name := readString()

	fmt.Print("Enter Age: ")
	age := readInt()

	fmt.Print("Enter Department: ")
	department := readString()

	fmt.Print("Enter Position: ")
	position := readString()

	fmt.Print("Enter Salary: ")
	salary := readFloat()

	employee := Employee{
		ID:         nextID,
		Name:       name,
		Age:        age,
		Department: department,
		Position:   position,
		Salary:     salary,
	}

	employees = append(employees, employee)
	nextID++

	fmt.Println("Employee added successfully.")
}

func viewEmployees() {
	fmt.Println("\n--- Employee List ---")

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, emp := range employees {
		printEmployee(emp)
	}
}

func searchEmployee() {
	fmt.Println("\n--- Search Employee ---")
	fmt.Print("Enter Employee ID: ")
	id := readInt()

	for _, emp := range employees {
		if emp.ID == id {
			fmt.Println("Employee found:")
			printEmployee(emp)
			return
		}
	}

	fmt.Println("Employee not found.")
}

func updateEmployee() {
	fmt.Println("\n--- Update Employee ---")
	fmt.Print("Enter Employee ID to update: ")
	id := readInt()

	for i, emp := range employees {
		if emp.ID == id {
			fmt.Print("Enter New Name: ")
			employees[i].Name = readString()

			fmt.Print("Enter New Age: ")
			employees[i].Age = readInt()

			fmt.Print("Enter New Department: ")
			employees[i].Department = readString()

			fmt.Print("Enter New Position: ")
			employees[i].Position = readString()

			fmt.Print("Enter New Salary: ")
			employees[i].Salary = readFloat()

			fmt.Println("Employee updated successfully.")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func deleteEmployee() {
	fmt.Println("\n--- Delete Employee ---")
	fmt.Print("Enter Employee ID to delete: ")
	id := readInt()

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted successfully.")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func printEmployee(emp Employee) {
	fmt.Println("----------------------------")
	fmt.Println("ID        :", emp.ID)
	fmt.Println("Name      :", emp.Name)
	fmt.Println("Age       :", emp.Age)
	fmt.Println("Department:", emp.Department)
	fmt.Println("Position  :", emp.Position)
	fmt.Println("Salary    :", emp.Salary)
}

func readString() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt() int {
	for {
		input := readString()
		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Print("Invalid number, enter again: ")
	}
}

func readFloat() float64 {
	for {
		input := readString()
		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num
		}
		fmt.Print("Invalid salary, enter again: ")
	}
}