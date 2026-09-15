package main

import "fmt"

func main() {
	x := 5
	p1 := &x
	p2 := &p1

	fmt.Println("x:", x)
	fmt.Println("p1:", p1, "->", *p1)
	fmt.Println("p2:", p2, "->", *p2, "->", **p2)

	**p2 = 15
	fmt.Println("After **p2 = 15, x:", x)
}
