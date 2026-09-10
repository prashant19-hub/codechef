package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Positive")
	} else if n < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
