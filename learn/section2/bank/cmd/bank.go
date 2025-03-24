package main

import (
	"fmt"

	"github.com/mrjxtr/Learn_Go_REST_API/cmd/section2/bank/internal"
	"github.com/mrjxtr/Learn_Go_REST_API/cmd/section2/bank/utils"
)

func main() {
	fmt.Printf("WELCOME TO THE BANK!\n\n")

	const balanceFile = "./cmd/section2/bank/data/balance.txt"

	balance, err := utils.GetValueFromFile(balanceFile)
	if err != nil {
		fmt.Print("ERROR:")
		fmt.Println(err)
	}

	// i: for testing purposes
	// balance := 1000.0

	wantsExit := false
	var choice int

	for !wantsExit {

		utils.PresentOptions()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n\n", balance)
		case 2:
			depositAmount, newBalance := internal.DepositMoney(&balance)
			fmt.Printf("You are depositing: %.2f\n", depositAmount)
			fmt.Printf("Your balance is : %.2f\n\n", newBalance)
			utils.WriteValueToFile(balanceFile, newBalance)
		case 3:
			withdrawAmount, newBalance := internal.WithdrawMoney(&balance)
			fmt.Printf("You are withdrawing: %.2f\n", withdrawAmount)
			fmt.Printf("Your balance is : %.2f\n\n", newBalance)
			utils.WriteValueToFile(balanceFile, newBalance)
		case 4:
			wantsExit = true
			fmt.Println("Exiting Bank, Thank you!")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// i: This is an alternative to the code above
		// if choice == 1 {
		// 	fmt.Printf("Your balance is: %.2f\n\n", balance)
		// } else if choice == 2 {
		//
		// 	deposit, newBalance := depositMoney(&balance)
		//
		// 	fmt.Printf("You are depositing: %.2f\n", deposit)
		// 	fmt.Printf("Your balance is : %.2f\n\n", newBalance)
		// } else if choice == 3 {
		//
		// 	withdrawAmount, newBalance := withdrawMoney(&balance)
		//
		// 	fmt.Printf("You are withdrawing: %.2f\n", withdrawAmount)
		// 	fmt.Printf("Your balance is : %.2f\n\n", newBalance)
		// } else if choice == 4 {
		//
		// 	wantsExit = true
		//
		// 	fmt.Println("Exiting Bank, Thank you!")
		// } else {
		// 	fmt.Println("Invalid choice. Please try again.")
		// }
	}
}
