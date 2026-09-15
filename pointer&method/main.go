package main

import "fmt"

type Counter struct {
	Value int
}

// value receiver
func (c Counter) Show() {
	fmt.Println("Value:", c.Value)
}

// pointer receiver (modify kar sakta hai)
func (c *Counter) Increment() {
	c.Value++
}

func main() {
	c := Counter{Value: 0}
	c.Show()

	c.Increment()
	c.Increment()
	c.Show()
}
