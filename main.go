package main

// fmt stands for format package
import (
	"fmt"
	"strings"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData = struct {
	fName           string
	lName           string
	email           string
	numberOfTickets uint
}

// func main is the entry point of a go application , A program can have only on main function
func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// fmt.Println("We have total of ", conferenceTickets,"tickets and ",remainingTickets,"are still available")

	for len(bookings) < 50 {
		//User inputs
		fName, lName, email, userTickets := getUserInput()

		//Validations
		isValidName, isValidEmail, isValidTixNumber := validateUserInput(fName, lName, email, userTickets, remainingTickets)

		//Printing first names
		var firstNames = printFirstNames()
		fmt.Printf("First Names are : %v\n", firstNames)

		if isValidName && isValidEmail && isValidTixNumber {
			fmt.Printf("Thanks %v %v for booking %v tickets , a confirmation email has been sent to %v \n", fName, lName, email, userTickets)

			// creating a map
			/*
				var userData = make(map[string]string)
				userData["fname"] = fName
				userData["lname"] = lName
				userData["email"] = email
				userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
			*/

			// Creating a struct for user
			var userData = UserData{
				fName:           fName,
				lName:           lName,
				email:           email,
				numberOfTickets: userTickets,
			}

			//Adding values to the bookings array
			bookings = append(bookings, userData)
			//calculating remaining tickets
			remainingTickets = remainingTickets - userTickets

			// create a map for a user

			fmt.Printf("No of tickets remaining : %v\n", remainingTickets)
			if remainingTickets == 0 {
				fmt.Println("Conference Full , Come back next year ")
				break
			}

		} else {
			// fmt.Printf("We only have %v tickets remaining so you can't book %v tickets \n",remainingTickets,userTickets)
			fmt.Printf("Input data invalid \n")
			if !isValidName {
				fmt.Print("First Name and Last Name should be greater than 2")
			}
			if !isValidEmail {
				fmt.Print("Email Adress should contain @")
			}
			if !isValidTixNumber {
				fmt.Print("Enter valid number of tickets , should be more than or equal to 0")
			}

		}
		go sendTicket(userTickets, fName, lName, email) // go keyword creates a go routine which is light weight thread managed by go runtime
	}
}

func greetUsers(confname string, conftickets uint, remtickets uint) {
	fmt.Printf("Hello , Welcome to the booking app ! , Book your tickets for the %v\n", confname)
	fmt.Printf("We have total of %v tickets and %v are available\n", conftickets, remtickets)
	fmt.Println("Get your tickets here")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["fName"])
		firstNames = append(firstNames, booking.fName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains("@", email)
	isValidTixNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTixNumber
}

func getUserInput() (string, string, string, uint) {
	var fName string
	var lName string
	var email string
	var userTickets uint
	fmt.Println("Enter your firstName : ")
	fmt.Scan(&fName)
	fmt.Println("Enter your lName : ")
	fmt.Scan(&lName)
	fmt.Println("Enter your email : ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want : ")
	fmt.Scan(&userTickets)

	return fName, lName, email, userTickets
}

func sendTicket(userTickets uint, fname string, lname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, fname, lname)
	fmt.Println("########################")
	fmt.Printf("Sending Tickets \n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
}

//Arrays
//var bookings = [50] string{} //array of data type string and size 50 OR var bookings [50] string

//----------------------------------//
//Getting value from the user
// var userName string
// var userTix int
// fmt.Scan(&userName) //& is a pointer , the value gets stored in memory .
// fmt.Scan(&userTix)

// fmt.Printf("%v bought %v tickets", userName, userTix)

// Add Two numbers
// var num1 , num2 , num3 int
// num3 = num1+num2
// fmt.Println("Enter the first number")
// fmt.Scan(&num1)
// fmt.Println("Enter the second number")
// fmt.Scan(&num2)
// fmt.Printf("The sume is %v", num3)

//WAP to greet a user
// var fName string
// fmt.Println("Enter your name : ")
// fmt.Scan(&fName)
// fmt.Printf("Hi %v , Good Morning",fName)

//Arrays in Go
//WAP to build a fruit array
// var fruits[]string
// var fruitName string
// fmt.Printf("Enter Fruit Name : %v\n ",fruitName)
// fmt.Scan(&fruitName)
// fruits = append(fruits,fruitName )
// fmt.Printf("The fruits are :  %v\n",fruits)

//Loops in go
// var employeeName string
// var empList [] string
// for{
// 	fmt.Printf("Enter employee name : \n")
// 	fmt.Scan(&employeeName)
// 	empList = append(empList,employeeName)
// 	firstNames:=[]string{}
// 	for _,employees:=range empList{
// 		var empNames = strings.Fields(employees)
// 		firstNames = append(firstNames,empNames[0])
// 	}
// 	fmt.Printf("The first Names area %v \n",firstNames)
// }

//Switch case in Go
/*
	var num1,num2 int
	var choice int
	fmt.Println("Enter two nos")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Println("1.Addition")
	fmt.Println("2.Subtraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")
	fmt.Println("Enter a choice")
	fmt.Scan(&choice)
	switch choice{
		case 1:
			fmt.Print(num1+num2)
		case 2:
			fmt.Print(num1-num2)
		case 3:
			fmt.Print(num1 * num2)
		case 4:
			fmt.Print(num1 / num2)
		default:
			fmt.Print("Invalid Choice")
	}
*/
