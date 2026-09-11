package main

import "fmt"

func main() {
	users := map[string]string{
		"admin": "admin123",
		"rahul": "rahul@123",
		"priya": "priya#456",
	}

	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if pass, ok := users[username]; ok {
		if pass == password {
			fmt.Println("Login successful. Welcome,", username)
		} else {
			fmt.Println("Wrong password")
		}
	} else {
		fmt.Println("User not found")
	}
}
