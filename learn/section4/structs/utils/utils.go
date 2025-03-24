package utils

import "fmt"

func PromptUser() (string, string, string, string) {
	firstName := GetUserData("Enter your first name: ")
	lastName := GetUserData("Enter your last name: ")
	birthDate := GetUserData("Enter your birthdate (MM/DD/YYY): ")
	age := GetUserData("Enter your Age: ")

	return firstName, lastName, birthDate, age
}

func GetUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input) // Scanln when you want the input to end with enter key
	return input
}
