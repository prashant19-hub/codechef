package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println(n, "is not prime")
		return
	}

	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(n, "is prime")
	} else {
		fmt.Println(n, "is not prime")
	}
}
