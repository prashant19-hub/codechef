package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GoInfo struct {
	Topic   string
	Details string
}

func main() {
	data := []GoInfo{
		{"go", "Go ek fast, simple aur compiled programming language hai."},
		{"goroutine", "Goroutine lightweight concurrent function hoti hai jo go keyword se chalti hai."},
		{"channel", "Channel goroutines ke beech data safely bhejne ke kaam aata hai."},
		{"struct", "Struct related fields ko ek jagah group karta hai."},
		{"interface", "Interface methods ka set define karta hai aur Go me implicit implementation hoti hai."},
		{"slice", "Slice dynamic collection hoti hai jo array se zyada use hoti hai."},
		{"map", "Map key-value pair store karta hai."},
		{"pointer", "Pointer kisi value ka memory address store karta hai."},
		{"package", "Package Go code ko organize karta hai. main package executable program ke liye hota hai."},
		{"gin", "Gin ek popular Go web framework hai jo API banane me use hota hai."},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Data Question Answer Program")
	fmt.Println("Type 'exit' to stop")

	for {
		fmt.Print("\nApna sawal likho: ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "exit" {
			fmt.Println("Program band ho gaya.")
			break
		}

		found := false

		for _, item := range data {
			if strings.Contains(input, item.Topic) {
				fmt.Println("Answer:", item.Details)
				found = true
			}
		}

		if !found {
			fmt.Println("Answer nahi mila. Try words like: go, goroutine, channel, struct, interface, slice, map, pointer, package, gin")
		}
	}
}
