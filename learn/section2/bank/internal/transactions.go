package internal

import "fmt"

func DepositMoney(balance *float64) (float64, float64) {
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

func WithdrawMoney(balance *float64) (float64, float64) {
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
