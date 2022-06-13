package main

import (
	"fmt"
	"sync"
	"time"
)

// Defining variables
var (
	conferenceName = "Go Conference"
	remainingTickets = 50
	// wg Waits for the launched goroutine to finish
	wg = sync.WaitGroup{}
	// bookings creating an empty list of userData struct
	bookings = make([]UserData, 0)
)

const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	// Calling the greetUser function
	greetUsers()

	// Getting the user input
	firstName, lastName, email, userTickets := getUserInput()

	// Validating the user input - calling a function to do this
	isValidName, isValidEmail, isValidUserTicketInput := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidUserTicketInput && isValidEmail && isValidName {

		// Calling the booking tickets function
		bookTicket(userTickets, firstName, lastName, email)
		// Sending the ticket to the user (go starts a new goroutine and goroutine is a lightweight thread managed by the Go runtime)
		// Add: Sets the number of goroutines to wait for(increases the counter by the provided number)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// call function print FirstNames
		firstNames := getFirstNames()
		fmt.Printf("The first name of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end the programm whenever the tickets becomes 0
			fmt.Printf("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered does not contain @ sign")
		}
		if !isValidUserTicketInput {
			fmt.Printf("Number of tickets you entered is invalid\n")
		}
	}
	// It waits until the counter becomes 0
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Inside the brakets we have the input parameters and outside the brakets we have the output parameters
func getFirstNames() []string {
	firstNames := []string{}

	// The _ is used as a Blank indentifier -> To ignor a variable we don't want to use
	// In Go we need to make unused variables explicit
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Creating a userData object of the Struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Adding a value to a slice instead of an array
	// Slices are dynamic arrays with dynamic size
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket: %v to email address %v\n", ticket, email)
	fmt.Println("############")
	// When the function is done this tells to the wait function to stop waiting and to continue with the new thread
	wg.Done()
}
