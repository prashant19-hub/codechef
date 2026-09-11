package main

import "fmt"

func main() {
	var n, rev int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	temp := n
	for temp != 0 {
		digit := temp % 10
		rev = rev*10 + digit
		temp /= 10
	}

	fmt.Println("Reverse:", rev)
}
