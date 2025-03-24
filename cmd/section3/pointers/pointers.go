package main

import "fmt"

func main() {
	age := 28
	fmt.Println("Age:", age)

	adultYears := getAdultYears(&age)
	fmt.Println("Adult Years:", adultYears)

	modifyAge(&age)              // Mutating age
	fmt.Println("New Age:", age) // This will no longer be 28

	var input string
	fmt.Scan(&input) // pass a pointer to directly mutate input
	fmt.Println("Your Input:", input)
}

// Using age variable value without creating a copy
func getAdultYears(age *int) int {
	return *age - 18
}

// Mutating age variable directly
func modifyAge(age *int) {
	*age++
}
