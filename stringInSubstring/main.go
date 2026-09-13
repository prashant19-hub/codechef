package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Welcome to Go programming"
	sub := "Go"

	if strings.Contains(text, sub) {
		fmt.Println(sub, "found in text")
	} else {
		fmt.Println(sub, "not found")
	}
}
