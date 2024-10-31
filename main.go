package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Oluwaseyieniola/booking-app-go/helper"
)

var conferencename = "Osapolor"
var remainingTickets uint = 50
var bookings = make([]Userdata, 0) // empty list of maps

const conferenceTickets int = 50

var wg = sync.WaitGroup{}

type Userdata struct {
	firstname    string
	lastname     string
	email        string
	bookedticket uint
}

func main() {

	IlovemyUsers(conferencename)

	firstname, lastname, email, bookedticket := getUserinput()

	// validate user
	isvalidName, isvalidEmail, isvalidUserticketsNumber := helper.ValidateUsers(firstname, lastname, email, bookedticket, remainingTickets)

	if isvalidName && isvalidEmail && isvalidUserticketsNumber {
		fmt.Printf("%v %v, with ticket number %v and email %v will be attending your event\n", firstname, lastname, bookedticket, email)

		BookingUser(bookedticket, firstname, lastname, email)
		wg.Add(1)
		go sendTicket(bookedticket, firstname, lastname, email)

		getfirstnames()
		fmt.Printf("The firstname include %v\n", firstname)

		if remainingTickets == 0 {
			fmt.Println("Tickets are soldout thank you!")

		}

	} else {
		if !isvalidName {
			fmt.Println("The firstname and last name you entered is too short")
		}
		fmt.Println("Invalid details")

		if !isvalidEmail {
			fmt.Println("The email address you entered is invalid it does not contain an @ sign")
		}
		if !isvalidUserticketsNumber {
			fmt.Printf("the amount of tickets %v you entered is more than available tickets %v\n ", bookedticket, remainingTickets)
		}

	}
	wg.Wait()

}

func IlovemyUsers(confName string) {
	fmt.Printf("Welcome to %v Live", confName)
}

func getfirstnames() []string {
	firstnames := []string{}
	for _, booking := range bookings {

		//var names = strings.Fields(booking)
		firstnames = append(firstnames, booking.firstname)
	}
	return firstnames
}

func getUserinput() (string, string, string, uint) {

	var firstname string
	var lastname string

	var email string
	var bookedticket uint

	fmt.Println("Please enter your firstname:")
	fmt.Scan(&firstname)
	fmt.Println("Please enter your lastname:")
	fmt.Scan(&lastname)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like:")
	fmt.Scan(&bookedticket)

	return firstname, lastname, email, bookedticket
}

func BookingUser(bookedticket uint, firstname string, lastname string, email string) {
	remainingTickets = remainingTickets - bookedticket

	var Userdata = Userdata{
		firstname:    firstname,
		lastname:     lastname,
		email:        email,
		bookedticket: bookedticket,
	}
	//not a map anymore
	//Userdata["Firstname"] = firstname
	//Userdata["lastname"] = lastname
	//Userdata["email"] = email
	//Userdata["bookedticket"] = strconv.FormatUint(uint64(bookedticket), 10)

	bookings = append(bookings, Userdata)

	fmt.Printf("The user includes %v\n", bookings)
	fmt.Printf("we have %v tickets left\n", remainingTickets)

}

func sendTicket(bookedticket uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("###########")
	var ticket = fmt.Sprintf(" %v tickets for %v  %v ", bookedticket, firstname, lastname)
	fmt.Printf("seding tickets %v to %v email", ticket, email)
	wg.Done()
}
