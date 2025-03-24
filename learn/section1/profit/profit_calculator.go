package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("ERROR: Revenue", err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("ERROR: Expenses", err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("ERROR: Tax Rate", err)
		return
	}

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt) // Earnings before Tax
	storeResult(ebt, "ebt")
	fmt.Printf("PROFIT: %.2f\n", profit)
	storeResult(profit, "profit")
	fmt.Printf("RATIO: %.2f\n", ratio)
	storeResult(ratio, "ratio")

	storeResults(ebt, profit, ratio)
}

func calculateProfits(rev, exp, tr float64) (float64, float64, float64) {
	ebt := rev - exp
	profit := ebt * (1 - tr/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("input must be a positive number")
	}

	return userInput, nil
}

func storeResult(data float64, filename string) {
	byteData := fmt.Sprintf("%.2f", data)
	filePath := fmt.Sprintf("./cmd/section1/profit/%s.txt", filename)

	os.WriteFile(filePath, []byte(byteData), 0644)
}

func storeResults(ebt, profit, ratio float64) {
	byteData := fmt.Sprintf("EBT: %.2f\nPROFIT: %.2f\nRATION: %.2f", ebt, profit, ratio)
	filePath := "./cmd/section1/profit/results.txt"

	os.WriteFile(filePath, []byte(byteData), 0644)
}
