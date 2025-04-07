package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mrjxtr/price-calc/conversion"
)

type PricesMap map[string]string

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	Result      PricesMap
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	var lines []string

	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	err = s.Err()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		file.Close()
		return
	}

	prices, err := conversion.StrsToFloats(lines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(PricesMap)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	fmt.Println(result)
}
