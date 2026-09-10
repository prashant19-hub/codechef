package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("Enter 5 numbers:")

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println("Sum:", sum)
}
