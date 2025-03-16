package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("WELCOME TO THE BANK!\n\n")

	const balanceFile = "./cmd/msc/bank/balance.txt"
	balance, err := readBalanceFromFile(balanceFile)
	if err != nil {
		fmt.Print("ERROR:")
		fmt.Println(err)
	}

	// i: for testing purposes
	// balance := 1000.0

	wantsExit := false
	var choice int

	for !wantsExit {

		fmt.Printf("What do you want to do?\n\n")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n\n", balance)
		case 2:
			depositAmount, newBalance := depositMoney(&balance)
			fmt.Printf("You are depositing: %.2f\n", depositAmount)
			fmt.Printf("Your balance is : %.2f\n\n", newBalance)
			writeBalanceToFile(balanceFile, newBalance)
		case 3:
			withdrawAmount, newBalance := withdrawMoney(&balance)
			fmt.Printf("You are withdrawing: %.2f\n", withdrawAmount)
			fmt.Printf("Your balance is : %.2f\n\n", newBalance)
			writeBalanceToFile(balanceFile, newBalance)
		case 4:
			wantsExit = true
			fmt.Println("Exiting Bank, Thank you!")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

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

func readBalanceFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("failed to find balanceFile")
	}

	balanceStr := string(data)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0, errors.New("failed to parse balance into float")
	}

	return balance, nil
}

func writeBalanceToFile(filename string, balance float64) {
	balanceStr := fmt.Sprint(balance)

	os.WriteFile(filename, []byte(balanceStr), 0644)
}

func depositMoney(balance *float64) (float64, float64) {
	var depositAmount float64
	fmt.Print("How much do you want to deposit?: ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Error: Invalid deposit amount")
		return 0, *balance
	}

	*balance += depositAmount
	return depositAmount, *balance
}

func withdrawMoney(balance *float64) (float64, float64) {
	var withdrawAmount float64
	fmt.Print("How much do you want to withdraw?: ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		fmt.Println("Error: Invalid withdraw amount")
		return 0, *balance
	}

	if withdrawAmount > *balance {
		fmt.Println("Error: Insufficient funds")
		return 0, *balance
	}

	*balance -= withdrawAmount
	return withdrawAmount, *balance
}
