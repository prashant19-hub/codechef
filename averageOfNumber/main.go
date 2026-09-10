package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	avg := float64(sum) / float64(len(nums))
	fmt.Println("Average:", avg)
}
