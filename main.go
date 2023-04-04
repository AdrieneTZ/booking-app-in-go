package main

// import Go packages
import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string // Slice

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still aviliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
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
		bookings = append(bookings, firstName+" "+lastName+",")

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceTickets)

		firstNames := []string{} // empty slice
		// "range" allows us to iterate over elements for different data structure
		// not only arrays and slices
		// for arrays and slices, "range" provides the index and value for each element
		// for <index>, <item> := range <arrays or slices> {logic}
		// "_": to ignore the variable that's not be used
		// unused variables need to be explicted in Go
		for _, booking := range bookings {
			// "strings.Fields": splits the string with white space as separator
			// and returns a slice with the split elements
			// "Tom Liz" => ["Tom","Liz"]
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("The first names of bookings are: %v.\n", firstNames)
	}
}
