package main
import "fmt"

func main(){
	var conferenceName = "YesConference"
	const conferenceTickets = 50
	var remainingTickets = 50 

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	
}
