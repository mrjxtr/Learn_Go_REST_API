package prices

import "fmt"

type PricesMap map[string]float64

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	Result      PricesMap
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10.0, 20.0, 30.0},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(PricesMap)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}
