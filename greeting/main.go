package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, I am", p.Name)
}

func main() {
	p := Person{Name: "Amit"}
	p.Greet()
}
