package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
