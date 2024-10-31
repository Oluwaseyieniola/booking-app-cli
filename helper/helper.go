package helper

import "strings"

func ValidateUsers(firstname string, lastname string, email string, bookedticket uint, remainingTickets uint) (bool, bool, bool) {
	isvalidName := len(firstname) > 2 && len(lastname) > 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidUserticketsNumber := bookedticket > 0 && bookedticket <= remainingTickets 
	return isvalidEmail, isvalidName, isvalidUserticketsNumber
}
