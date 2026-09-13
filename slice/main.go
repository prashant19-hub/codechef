package main

import "fmt"

func main() {
	nums := []int{12, 5, 89, 3, 45}

	min := nums[0]
	max := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
