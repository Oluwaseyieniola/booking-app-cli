# booking-app-cli
A quick and easy to learn cli application for golang beginners to learn about the basics of Golang. Including concepts like; Formatting, Go concurrency, Arrays, slices, structs, loops, conditional statements and final tests
The CLI Booking App allows users to book tickets for a conference, providing a real-time view of available tickets and validating user inputs. It’s built with Golang and leverages concurrency for sending booking confirmation asynchronously.

Features
-Interactive CLI: Users can enter their details and book tickets via the command line.
-Validation: Ensures that user input (name, email, and ticket quantity) meets specific requirements.
-Concurrency: Sends booking confirmation emails in the background, showcasing Golang’s Goroutines and sync.WaitGroup.
-User-Friendly Output: Provides clear messages about available tickets, errors, and booking confirmations.
## INSTALLATION
`git clone https://github.com/Oluwaseyieniola/booking-app-go`
`cd booking-app-go`
`go version`
`go run main.go`

*How It Works*
User Input: The user is prompted to enter their first name, last name, email, and the number of tickets they wish to book.
Validation: The app validates user input. If the input is invalid (e.g., too many tickets requested or an invalid email format), it displays relevant error messages.
Concurrency: If the input is valid, it adds the booking to the list and then spawns a Goroutine to simulate sending a ticket confirmation email.
Ticket Availability: The app dynamically adjusts the number of remaining tickets after each booking.

booking-app-go/
├── main.go                # Main application logic
├── helper/
│   └── validation.go      # Helper functions for input validation
├── user_test.go           # Unit tests for validation, booking, and utility functions
└── README.md              # Detailed explanation and instructions


*Detailed Walkthrough*
1. Variables and Constants
conferenceName: Stores the name of the conference.
remainingTickets: Tracks the available tickets.
conferenceTickets: The total number of tickets (constant).
2. Data Structure
The UserData struct stores user information such as firstName, lastName, email, and bookedTicket count. A slice (bookings) holds all bookings.

3. Functions
main()
The main function manages the program’s flow, including greeting users, gathering inputs, validating data, and booking tickets. It utilizes Goroutines to handle sending emails in the background.

welcomeUsers()
Displays a welcome message with conference details.

getUserInput()
Prompts the user for booking details and returns them for validation.

validateUsers()
Checks if user input meets the following criteria:

First and last names are longer than two characters.
Email contains an "@" symbol.
Requested tickets do not exceed the number of remaining tickets.
bookUser()
Adds a valid booking to the list and decreases remainingTickets. It uses a mutex (mu) to avoid race conditions when multiple bookings occur simultaneously.

sendTicket()
Simulates sending a booking confirmation. The function waits 10 seconds, then prints a confirmation message. It uses sync.WaitGrou

