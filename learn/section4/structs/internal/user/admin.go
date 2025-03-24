package user

import (
	"errors"
	"strconv"
	"time"
)

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(
	firstName, lastName, birthDate, age, email, password string,
) (*Admin, error) {
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
	return &Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			BirthDate: birthDate,
			Age:       ageStr,
			CreatedAt: time.Now().Format("2006-01-02"),
		},
	}, nil
}
