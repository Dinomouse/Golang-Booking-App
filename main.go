package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	//Printf allows us to format a string with %v and \n creates a new line explicitly

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	


	var firstName string
	var lastName string
	var email string
	var userTickets uint

	var bookings [50]string

	//ask for input
	fmt.Println("Enter your first name: ")
	//& is a pointer and points to the memory address of where userName is stored
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName 
	

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("There are %v tickets remaining for %v\n",remainingTickets,conferenceName)
	fmt.Println(bookings)
}
