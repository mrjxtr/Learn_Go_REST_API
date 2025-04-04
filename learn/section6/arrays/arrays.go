package main

import "fmt"

func main() {
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
