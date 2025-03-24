package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birthdate (MM/DD/YYY): ")
	age, _ := strconv.Atoi(getUserData("Enter your Age: "))
	createdAt := time.Now().Format("2006-01-02")

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Age:       age,
		CreatedAt: createdAt,
	}

	// Do something with the data...

	user.outputUserDetails()
	user.clearUserName()
	user.outputUserDetails()
}

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	Age       int
	CreatedAt string
}

func (u *User) outputUserDetails() {
	fmt.Printf(
		"\nYour Name: %v %v\nYour Birth Date: %v\nYour Age: %v\n\nDate Created: %v\n",
		u.FirstName,
		u.LastName,
		u.BirthDate,
		u.Age,
		u.CreatedAt,
	)
}

// Make sure to receive a pointer to a struct when mutating
func (u *User) clearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}
