package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func (user User) outputUserDatails() {
	fmt.Printf("My name is %v %v (born on %v) ", user.firstName, user.lastName, user.birthDate)
	//fmt.Println("My name is " + user.firstName + " " + user.firstName)
}

func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthDate:   bDate,
		createdDate: created,
	}

	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser = NewUser(firstName, lastName, birthdate)

	// ... do something awesome with that gathered data!

	newUser.outputUserDatails()
}

func getUserData(promtText string) string {
	fmt.Print(promtText)

	userInput, _ := reader.ReadString('\n')
	//cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return userInput
}
