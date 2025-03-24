package user

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	Age       int
	CreatedAt string
}

// Create a Go struct constructor function
func New(firstName, lastName, birthDate, age string) (*User, error) {
	// Adding validation on a constructor
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New(
			"first name, last name, and birth date should not be empty",
		)
	}

	ageStr, err := strconv.Atoi(age)
	if err != nil {
		return nil, errors.New("age must be a whole number")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Age:       ageStr,
		CreatedAt: time.Now().Format("2006-01-02"),
	}, nil
}

func (u *User) OutputUserDetails() {
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
func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
