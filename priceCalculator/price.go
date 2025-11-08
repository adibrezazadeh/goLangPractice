package priceCalculator

import "example.com/first/priceCalculator/prices"

func PriceCalculator() {
	taxRate := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
