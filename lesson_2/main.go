package main

import (
	"fmt"
	"strings"
)

func main(){
	applicationName := "HackDay"
	const numberOfTickets uint = 10
	var remainingTickets uint = numberOfTickets
	var username string
	var userTickets uint
	var bookings []string

	fmt.Printf("Welcome to %v\n", applicationName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", numberOfTickets, remainingTickets)

	for remainingTickets > 0 /*&& len(bookings) < 10*/ {
		fmt.Printf("\n")
		fmt.Printf("\n")

		fmt.Printf("Enter your username: \n")
		fmt.Scan(&username)

		fmt.Printf("Enter number of Tickets: \n")
		fmt.Scan(&userTickets)

		bookings = append(bookings, username+" public")

		var isUsernameNotEmpty bool = len(username) >= 2
		var isValid bool = strings.Contains(username, "@")
		var isValidTicketNumber bool = userTickets <= remainingTickets

		if isUsernameNotEmpty && isValid && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
		
			fmt.Printf("Thank you %v for booking %v, you will receive a confirmation sooner to the event date.\n", username, userTickets)
			
			names := []string{}
			for _, booking := range bookings{
				var name = strings.Fields(booking)
				names = append(names, name[0])
			}

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, applicationName)
			fmt.Printf("%v ticket holders\n\n", names)
			
			noTicketsRemaining := remainingTickets == 0

			if  noTicketsRemaining {
				fmt.Println("full booked goodbye")
				break
			}
		} else{
			if !isValid{
				fmt.Println("user name does not contain @.")
			}
			if !isUsernameNotEmpty{
				fmt.Println("user name is too short.")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid.")
			}
		}		
	}
} 