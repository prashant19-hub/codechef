package main

import (
	"fmt"
)

type Flight struct {
	FlightNo    string
	Source      string
	Destination string
	TotalSeats  int
	BookedSeats int
}

func (f Flight) ShowDetails() {
	fmt.Println("----------------------------")
	fmt.Println("Flight No   :", f.FlightNo)
	fmt.Println("Source      :", f.Source)
	fmt.Println("Destination :", f.Destination)
	fmt.Println("Total Seats :", f.TotalSeats)
	fmt.Println("Booked Seats:", f.BookedSeats)
	fmt.Println("Available   :", f.TotalSeats-f.BookedSeats)
}

func main() {
	flights := []Flight{
		{"AI101", "Delhi", "Mumbai", 100, 20},
		{"IN202", "Lucknow", "Bangalore", 120, 50},
	}

	var choice int

	for {
		fmt.Println("\n===== Airline Management System =====")
		fmt.Println("1. Show all flights")
		fmt.Println("2. Add new flight")
		fmt.Println("3. Book seat")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\nAvailable Flights:")
			for i := 0; i < len(flights); i++ {
				flights[i].ShowDetails()
			}

		case 2:
			var flightNo, source, destination string
			var totalSeats int

			fmt.Print("Enter flight number: ")
			fmt.Scan(&flightNo)

			fmt.Print("Enter source: ")
			fmt.Scan(&source)

			fmt.Print("Enter destination: ")
			fmt.Scan(&destination)

			fmt.Print("Enter total seats: ")
			fmt.Scan(&totalSeats)

			newFlight := Flight{
				FlightNo:    flightNo,
				Source:      source,
				Destination: destination,
				TotalSeats:  totalSeats,
				BookedSeats: 0,
			}

			flights = append(flights, newFlight)
			fmt.Println("New flight added successfully.")

		case 3:
			var flightNo string
			found := false

			fmt.Print("Enter flight number to book seat: ")
			fmt.Scan(&flightNo)

			for i := 0; i < len(flights); i++ {
				if flights[i].FlightNo == flightNo {
					found = true
					if flights[i].BookedSeats < flights[i].TotalSeats {
						flights[i].BookedSeats++
						fmt.Println("Seat booked successfully.")
						fmt.Println("Remaining seats:", flights[i].TotalSeats-flights[i].BookedSeats)
					} else {
						fmt.Println("No seats available.")
					}
					break
				}
			}

			if !found {
				fmt.Println("Flight not found.")
			}

		case 4:
			fmt.Println("Exiting Airline Management System.")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
