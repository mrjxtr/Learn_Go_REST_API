package main

import (
	"github.com/mrjxtr/price-calc/prices"
)

func main() {
	// pricesList := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
