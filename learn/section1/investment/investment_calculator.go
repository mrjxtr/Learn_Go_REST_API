package main

import (
	"fmt"
	"math"
)

const inflation = 2.4

func main() {
	var investment float64
	var years float64
	interestRate := 4.2

	fmt.Print("Investment: ")
	fmt.Scan(&investment)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Interest Rate: ")
	// fmt.Scan(&interestRate)
	fmt.Scanln(&interestRate) // do this to use the set default value

	roi, aroi := calculateROIs(investment, interestRate, years)

	fmt.Printf("This is your total return: %.2f\n", roi)
	fmt.Printf("This is your total return (adjusted for inflation): %.2f\n", aroi)
}

func calculateROIs(inv, intR, yrs float64) (float64, float64) {
	roi := inv * math.Pow(1+intR/100, yrs)
	rfv := roi / math.Pow(1+inflation/100, yrs)
	return roi, rfv
}
