package main

import (
	"fmt"
	"os"
)

func makeArrays() {
	// Use this if you know in advance how much you will add to the array
	userNames := make([]string, 2, 6) // creates a slice with 2 slots and 6 capacity
	fmt.Println(userNames)

	userNames[0] = "Jester" // assign values to existing slots
	userNames[1] = "Sione"
	fmt.Println(userNames)

	userNames = append(userNames, "Jesus") // append values
	fmt.Println(userNames)
}

func main() {
	exit := true
	makeArrays()

	if exit {
		fmt.Printf("\nexiting early\n")
		os.Exit(0)
	}

	// arrays are collection of data with the same type i.e different numbers
	var emptyArray [4]string    // creating an empty string
	emptyArray[2] = "Something" // adding a value to a specific index of an array

	fmt.Println(emptyArray) // We have empty values plus "Something" in idx 2

	producNames := [4]string{"Coffee", "Iced Tea", "Milk Tea", "Water"}

	fmt.Println(producNames)
	fmt.Println(producNames[0]) // Print specific elements in array

	prices := [4]float64{10.99, 11.5, 9.5, 5.1}
	fmt.Println("price:", prices)
}
