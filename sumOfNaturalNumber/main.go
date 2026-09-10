package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("Sum of first", n, "natural numbers:", sum)
}
