package main

import "fmt"

func main() {
	var d, t int
	fmt.Scan(&d, &t)

	if d > t {
		fmt.Println(d - t)
	} else {
		fmt.Println(0)
	}
}


// Teleport Home:-

// Chef has traveled a long way, and now wants to get home.

// Chef is D kilometers away from home, and he can walk at a speed of 
// 1 kilometer per hour.
// Chef also has the ability to teleport. He can teleport for a distance of at most 
// T kilometers, which happens instantly and doesn't require any time.

// The teleport can be used at most once.

// Find the minimum time, in hours, that Chef needs to reach home.