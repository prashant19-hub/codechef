package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateAge(u *User, newAge int) {
	u.Age = newAge
}

func main() {
	user := User{Name: "Alice", Age: 25}
	fmt.Println("Before:", user)

	updateAge(&user, 30)
	fmt.Println("After:", user)
}
