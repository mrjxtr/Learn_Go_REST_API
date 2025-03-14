package main

import "fmt"

func main() {
	fmt.Printf("WELCOME TO THE BANK!\n\n")

	balance := 1000.0
	wantsExit := false

	for !wantsExit {
		var choice int
		var newBalance float64

		fmt.Printf("What do you want to do?\n\n")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is: %v\n\n", balance)
		} else if choice == 2 {
			var deposit float64

			deposit, newBalance = depositMoney(balance)
			balance = newBalance

			fmt.Printf("You are depositing: %.2f\n", deposit)
			fmt.Printf("Your new balance is : %.2f\n\n", newBalance)
		} else if choice == 3 {
			var withdrawAmount float64

			withdrawAmount, newBalance := withdrawMoney(balance)
			balance = newBalance

			fmt.Printf("You are withdrawing: %.2f\n", withdrawAmount)
			fmt.Printf("Your new balance is : %.2f\n\n", newBalance)
		} else {
			wantsExit = true
			fmt.Println("Exiting Bank, Thank you!")
		}
	}
}

func depositMoney(balance float64) (float64, float64) {
	var deposit float64
	fmt.Print("How much do you want to deposit?: ")
	fmt.Scan(&deposit)
	newBalance := deposit + balance
	return deposit, newBalance
}

func withdrawMoney(balance float64) (float64, float64) {
	var withdrawAmount float64
	fmt.Print("How much do you want to withdraw?: ")
	fmt.Scan(&withdrawAmount)
	newBalance := balance - withdrawAmount
	return withdrawAmount, newBalance
}
