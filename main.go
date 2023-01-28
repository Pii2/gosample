package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := conferenceTickets
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets\n", firstName, lastName, userTickets)
		fmt.Printf("%v tickets are still available\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("Ther are all our bookings: %v\n", firstNames)
	}

}
