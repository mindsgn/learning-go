package main

import "fmt"

func main(){
	applicationName := "HackDay"
	const numberOfTickets uint = 10
	var remainingTickets uint = numberOfTickets
	var username string
	var userTickets int

	fmt.Printf("Welcome to %v\n", applicationName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", numberOfTickets, remainingTickets)

	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Printf("Enter your username: ")
	fmt.Scan(&username)

	fmt.Printf("Enter number of Tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v, you will receive a confirmation sooner to the event date.", username, userTickets)
	//fmt.Printf("User %v book %v tickets. \n", username, userTickets)
}