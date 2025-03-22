package utils

import "fmt"

// Spliting packages into multiple files
func PresentOptions() {
	fmt.Printf("What do you want to do?\n\n")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
