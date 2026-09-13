package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value at address:", *p)

	*p = 20
	fmt.Println("After *p = 20, x:", x)
}
