package main

import "fmt"

func main() {
	marks := map[string]int{
		"Amit":  85,
		"Priya": 92,
		"Rahul": 78,
	}

	fmt.Println("Priya ke marks:", marks["Priya"])

	marks["Rohan"] = 88
	fmt.Println("All marks:", marks)
}
