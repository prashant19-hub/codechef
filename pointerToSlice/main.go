package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 999
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Before:", nums)

	modifySlice(nums)
	fmt.Println("After:", nums)
}
