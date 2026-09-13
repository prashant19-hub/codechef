package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func makeSound(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSound(dog)
	makeSound(cat)
}
