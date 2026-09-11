package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	rect := Rectangle{Length: 10, Width: 5}
	fmt.Println("Area:", rect.Area())
}
