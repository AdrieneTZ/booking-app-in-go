package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still aviliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// declare variables

	// var bookings [50]string
	// change Array to Slice for better flexibility
	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for inputting their first name, last name, email address and numbers of ticket
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceTickets)

	fmt.Printf("These are all our bookings: %v.\n", bookings)
}
