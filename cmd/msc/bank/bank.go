package main

import "fmt"

func main() {
	var choice int
	balance := 1000.0

	fmt.Println("WELCOME TO THE BANK!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsExit := choice == 4
	if !wantsExit {
		if choice == 1 {
			fmt.Println("Your balance is:", balance)
		} else if choice == 2 {
			deposit := depositMoney()

			fmt.Printf("You are depositing: %.2f\n", deposit)
			fmt.Printf("Your new balance is : %.2f\n", deposit+balance)
		} else if choice == 3 {

			withdrawAmount := withdrawMoney()

			fmt.Printf("You are withdrawing: %.2f\n", withdrawAmount)
			fmt.Printf("Your new balance is : %.2f\n", balance-withdrawAmount)
		}
	} else {
		fmt.Println("Exiting Bank, Thank you!")
	}
}

func depositMoney() float64 {
	var deposit float64
	fmt.Print("How much do you want to deposit?: ")
	fmt.Scan(&deposit)
	return deposit
}

func withdrawMoney() float64 {
	var withdrawAmount float64
	fmt.Print("How much do you want to withdraw?: ")
	fmt.Scan(&withdrawAmount)
	return withdrawAmount
}
