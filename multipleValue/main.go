package main

import "fmt"

func calc(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	s, d := calc(25, 10)
	fmt.Println("Sum:", s, "Diff:", d)
}
