package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 1

	//Printf allows us to format a string with %v and \n creates a new line explicitly

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName,userTickets)


}
