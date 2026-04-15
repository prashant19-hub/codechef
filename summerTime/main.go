package main

import "fmt"

func main() {
var X int
fmt.Scan(&X)

if X > 35 {
fmt.Println("Yes")
} else {
fmt.Println("No")
}
}


//Question:

// Summer Time:
// Mamalesh likes to drink mango lassi when it's hot, and only when it's hot. If (and only if) the temperature on a given day is strictly greater than 
// 35
// 35 degrees, Mamalesh will drink mango lassi.

// Mamalesh sees that today's temperature is 
// X
// X degrees Celsius. Will he drink mango lassi today?
// Print "Yes" if he will, and "No" otherwise (without quotes).