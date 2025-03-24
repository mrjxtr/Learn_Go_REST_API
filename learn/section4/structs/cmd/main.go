package main

import (
	"fmt"

	"github.com/mrjxtr/Learn_Go_REST_API/learn/section4/structs/internal/user"
	"github.com/mrjxtr/Learn_Go_REST_API/learn/section4/structs/utils"
)

func main() {
	// i: Using the constructor method is much more idomatic
	firstName, lastName, birthDate, age := utils.PromptUser()

	// Using constructor pattern to make the code below more concise
	appUser, err := user.New(firstName, lastName, birthDate, age)
	if err != nil {
		panic(fmt.Sprintf("Error while creating a new user: %s", err))
	}

	appAdmin, err := user.NewAdmin(
		firstName,
		lastName,
		birthDate,
		age,
		"example@admin.com", // Example admin email and password
		"examplepassword",
	)
	if err != nil {
		panic(fmt.Sprintf("Error while creating a new admin: %s", err))
	}

	// i: Using this method is longer
	// firstName := utils.GetUserData("Enter your first name: ")
	// lastName := utils.GetUserData("Enter your last name: ")
	// birthDate := utils.GetUserData("Enter your birthdate (MM/DD/YYY): ")
	// age := utils.GetUserData("Enter your Age: ")
	// ageStr, err := strconv.Atoi(age)
	// if err != nil {
	// 	panic("Age must be a whole number")
	// }
	// createdAt := time.Now().Format("2006-01-02")

	// i: You have to individually set values
	// appUser := user.User{
	// 	FirstName: firstName,
	// 	LastName:  lastName,
	// 	BirthDate: birthDate,
	// 	Age:       ageStr,
	// 	CreatedAt: createdAt,
	// }

	// Do something with the data...

	fmt.Printf("\nUser Details")
	appUser.OutputUserDetails() // First output
	// appUser.ClearUserName()
	// appUser.OutputUserDetails() // Second output after clearing name

	fmt.Printf("\nAdmin Details:")
	appAdmin.OutputUserDetails() // Third output for admin
}
