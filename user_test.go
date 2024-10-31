package main

import (
	"testing"
	"github.com/Oluwaseyieniola/booking-app-go/helper"
)

// Test validation logic with valid and invalid inputs
func TestValidateUsers(t *testing.T) {
	var remainingTickets uint = 10

	tests := []struct {
		firstName    string
		lastName     string
		email        string
		ticketNumber uint
		expectedName bool
		expectedEmail bool
		expectedTicket bool
	}{
		{"John", "Doe", "john.doe@example.com", 1, true, true, true},
		{"Jo", "D", "johnexample.com", 1, false, false, true},
		{"Jane", "Smith", "jane.smith@example.com", 11, true, true, false},
	}

	for _, tt := range tests {
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUsers(tt.firstName, tt.lastName, tt.email, tt.ticketNumber, remainingTickets)

		if isValidName != tt.expectedName {
			t.Errorf("got %v, want %v for name validation", isValidName, tt.expectedName)
		}
		if isValidEmail != tt.expectedEmail {
			t.Errorf("got %v, want %v for email validation", isValidEmail, tt.expectedEmail)
		}
		if isValidTicketNumber != tt.expectedTicket {
			t.Errorf("got %v, want %v for ticket number validation", isValidTicketNumber, tt.expectedTicket)
		}
	}
}

// Test booking logic and check remaining tickets count
func TestBookUser(t *testing.T) {
	var testBookings = []UserData{}
	var testRemainingTickets uint = 10
	firstName := "Alice"
	lastName := "Johnson"
	email := "alice.johnson@example.com"
	ticketNumber := uint(2)

	bookUser(ticketNumber, firstName, lastName, email)

	if testRemainingTickets != 8 {
		t.Errorf("remaining tickets is %v, expected 8", testRemainingTickets)
	}

	if len(testBookings) != 1 {
		t.Errorf("booking count is %v, expected 1", len(testBookings))
	}

	if testBookings[0].firstName != firstName || testBookings[0].lastName != lastName || testBookings[0].email != email || testBookings[0].bookedTicket != ticketNumber {
		t.Errorf("booking details do not match expected values")
	}
}

// Test extracting first names from bookings
func TestGetFirstNames(t *testing.T) {
	bookings = []UserData{
		{firstName: "Alice", lastName: "Smith", email: "alice@example.com", bookedTicket: 1},
		{firstName: "Bob", lastName: "Jones", email: "bob@example.com", bookedTicket: 2},
	}

	firstNames := getFirstNames()

	expectedFirstNames := []string{"Alice", "Bob"}
	for i, name := range firstNames {
		if name != expectedFirstNames[i] {
			t.Errorf("got %v, want %v", name, expectedFirstNames[i])
		}
	}
}
