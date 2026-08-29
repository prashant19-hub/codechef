package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getResponse(input string) string {
	msg := strings.ToLower(strings.TrimSpace(input))

	switch {
	case msg == "hi" || msg == "hello":
		return "Hello! How can I help you?"
	case strings.Contains(msg, "your name"):
		return "I am a Go chatbot."
	case strings.Contains(msg, "how are you"):
		return "I am fine. Thanks for asking."
	case msg == "bye" || msg == "exit":
		return "Goodbye!"
	default:
		return "Sorry, I don't understand that yet."
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Go Chatbot started. Type 'bye' to exit.")

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()
		reply := getResponse(userInput)

		fmt.Println("Bot:", reply)

		if strings.ToLower(strings.TrimSpace(userInput)) == "bye" ||
			strings.ToLower(strings.TrimSpace(userInput)) == "exit" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}


//chatBot/main.go