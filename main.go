package main

import (
	"bookingapp/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	//var conferenceName string = "Go Conference"

	//var bookings [50]string
	//var bookings []string

	greetUsers()
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T,conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)
	//fmt.Println("Welcome to ", conferenceName, " booking application")
	//fmt.Printf("Welcome to %v booking application\n", conferenceName)
	//fmt.Println("We have total ", conferenceTickets, " tickets and ", remainingTickets, " are still available")
	//fmt.Println("Get your tickets here to attend")
	for {
		//code for user input fn calling
		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)
		//isValidCity := city == "Singapore" || city == "Landon"

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTicket, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

			//call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings are:%v\n", firstNames)
			//var noTicketsRemaining bool = remainingTickets == 0
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//endprgm
				fmt.Println("Our conference is booked out.Come back next year")
				break
			}
			//} else if userTicket == remainingTickets {
			//dosomethingelse
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address  you entered does'nt contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("No of tickets you entered is invalid ")
			}

			//fmt.Printf("We only have %v tickets remaining,So you cant book %v tickets\n", remainingTickets, userTicket)
			//fmt.Printf("Your Input data is Invalid,try again\n")
			//continue
		}

	}
	wg.Wait()
}

/*
SWITCH STATEMENT

city:=Landon;
switch city {
case "Newyork":
//execute code for Newyork conference tickets
case "Singapore", "Hong kong":
//execute code for Singapore & Hong Kong conference tickets
case "Landon", case "Landon":
//some code here :
case "Mexico City":
//execute code for Mexico city conference tickets


default:
fmt.Print("No valid city selected")
}*/
func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceTickets)
	fmt.Printf("We have total %v tickets and %v   are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)

		firstNames = append(firstNames, booking.firstName)
	}
	//fmt.Printf("These are all the bookings:%v\n", bookings)
	return firstNames
}

//validateuserinput

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	//serTicket = 2
	//fmt.Println(userName)
	//fmt.Printf("user %v booked %v tickets\n", firstName, userTicket)
	//ask foruser input

	/*fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)*/
	fmt.Println("enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter your email")
	fmt.Scan(&email)
	fmt.Println("enter number of tickets")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}
func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket
	//create a map for a user

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}
	/*userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)*/
	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	/*fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("slice type: %T\n", bookings)
	fmt.Printf("slice length: %v\n", len(bookings))*/

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at  %v\n ", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("############")
	wg.Done()
}
