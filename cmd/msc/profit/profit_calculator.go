package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("PROFIT: %.2f\n", profit)
	fmt.Printf("RATIO: %.2f\n", ratio)
}

func calculateProfits(rev, exp, tr float64) (float64, float64, float64) {
	ebt := rev - exp
	profit := ebt * (1 - tr/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
