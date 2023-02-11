package main

import (
	"fmt"
	"strings"
)

//go needs to know where to start executing the program(entrypoint)
func maina() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsersa(conferenceName, conferenceTickets, remainingTickets)

	//Ask user for ther input
	//create an infinity loop cooz we use cli to keep asking mutiple users new bookng after a booking has done
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		//Validating inputs
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v for booking %v ticket(s), you will receive confirmation email at %v \n", firstName, userTickets, email)
			fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
			//print first
			firstNames := []string{}
			//for each value take first name/range provides the index(_ as blanck identifier/unused variabled) and value for each element
			for _, booking := range bookings {
				//splite the string with white space as separator using built in func/Field("timotheo Mhoja"-> "timotheo","Mhoja")
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of the bookings are: %v \n", firstNames)
			if remainingTickets == 0 {
				//End program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address you intered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
		}

	}

}

func greetUsersa(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking applications \n", confName)
	fmt.Printf("We have  total of %v tickets and %v are stll available \n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

//refactored snippet
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// //package level variables
// const conferenceTickets = 50

// var conferenceName = "Go conference"
// var remainingTickets uint = 50
// var bookings []string

// //go needs to know where to start executing the program(entrypoint)
// func main() {

// 	greetUsers()

// 	//Ask user for ther input
// 	//create an infinity loop cooz we use cli to keep asking mutiple users new bookng after a booking has done
// 	for {
// 		//get user inputs
// 		firstName, lastName, email, userTickets := getUserInput()
// 		//vaidata user input func
// 		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

// 		if isValidName && isValidEmail && isValidTicketNumber {
// 			//bookinf func
// 			bookTicket(userTickets, firstName, lastName, email)
// 			//print first name func
// 			firstNames := getFirstNames()
// 			fmt.Printf("The first names of the bookings are: %v \n", firstNames)

// 			if remainingTickets == 0 {
// 				//End program
// 				fmt.Println("Our conference is booked out. Come back next year.")
// 				break
// 			}

// 		} else {
// 			if !isValidName {
// 				fmt.Println("First name or Last name you entered is too short!")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Email address you intered doesn't contain @ sign")
// 			}
// 			if !isValidTicketNumber {
// 				fmt.Println("Number of tickets you entered is invalid!")
// 			}
// 		}

// 	}

// }

// func greetUsers() {
// 	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
// 	fmt.Printf("We have  total of %v tickets and %v are stll available \n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	//for each value take first name/range provides the index(_ as blanck identifier/unused variabled) and value for each element
// 	for _, booking := range bookings {
// 		//splite the string with white space as separator using built in func/Field("timotheo Mhoja"-> "timotheo","Mhoja")
// 		var names = strings.Fields(booking)
// 		firstNames = append(firstNames, names[0])
// 	}
// 	//input and output params
// 	return firstNames
// }

// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	//Validating inputs
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Enter your email: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)
// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []string) {
// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, firstName+" "+lastName)
// 	fmt.Printf("Thank you %v for booking %v ticket(s), you will receive confirmation email at %v \n", firstName, userTickets, email)
// 	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
// 	return remainingTickets, bookings
// }
//Change map to struct usage
// package main

// import (
// 	"booking-app/helper"
// 	"fmt"
// 	"strconv"
// )

// //package level variables
// const conferenceTickets = 50

// var conferenceName = "Go conference"
// var remainingTickets uint = 50
// var bookings = make([]map[string]string, 0) //creating empty list of map[]
// //define the structure which can handle mixed data types at same time
// //cretate a type called 'UserData' based on a structure=class in java of firstName, lastName,email,numberOfTickets
// type UserData struct {
// 	firstName       string
// 	lastName        string
// 	email           string
// 	numberOfTickets uint
// }

// //go needs to know where to start executing the program(entrypoint)
// func main() {

// 	greetUsers()

// 	//Ask user for ther input
// 	//create an infinity loop cooz we use cli to keep asking mutiple users new bookng after a booking has done
// 	for {
// 		//get user inputs
// 		firstName, lastName, email, userTickets := getUserInput()
// 		//vaidata user input func
// 		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

// 		if isValidName && isValidEmail && isValidTicketNumber {
// 			//bookinf func
// 			bookTicket(userTickets, firstName, lastName, email)
// 			//print first name func
// 			firstNames := getFirstNames()
// 			fmt.Printf("The first names of the bookings are: %v \n", firstNames)

// 			if remainingTickets == 0 {
// 				//End program
// 				fmt.Println("Our conference is booked out. Come back next year.")
// 				break
// 			}

// 		} else {
// 			if !isValidName {
// 				fmt.Println("First name or Last name you entered is too short!")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Email address you intered doesn't contain @ sign")
// 			}
// 			if !isValidTicketNumber {
// 				fmt.Println("Number of tickets you entered is invalid!")
// 			}
// 		}

// 	}

// }

// func greetUsers() {
// 	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
// 	fmt.Printf("We have  total of %v tickets and %v are stll available \n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	//for each value take first name/range provides the index(_ as blanck identifier/unused variabled) and value for each element
// 	for _, booking := range bookings {
// 		//splite the string with white space as separator using built in func/Field("timotheo Mhoja"-> "timotheo","Mhoja")
// 		firstNames = append(firstNames, booking["firstName"])
// 	}
// 	//input and output params
// 	return firstNames
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Enter your email: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)
// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []string) {
// 	remainingTickets = remainingTickets - userTickets
// 	//create a map(data type) for user(map[key data type]value data type)/and create an empty map by using make
// 	var userData = make(map[string]string)
// 	userData["firstName"] = firstName
// 	userData["lastName"] = lastName
// 	userData["email"] = email
// 	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

// 	//add all key value data to the map
// 	bookings = append(bookings, userData)
// 	fmt.Printf("List of bookins is:  %v \n", bookings)
// 	fmt.Printf("Thank you %v for booking %v ticket(s), you will receive confirmation email at %v \n", firstName, userTickets, email)
// 	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
// 	return 0, nil
// 	// return remainingTickets, sbookings
// }
