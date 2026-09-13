package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function done")
}
