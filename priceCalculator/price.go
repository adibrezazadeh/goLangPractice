package pricecalculator

import "fmt"

func PriceCalculator() {
	var prices []float64 = []float64{10, 20, 30}
	taxRate := []float64{0, 0.7, 0.1, 0.15}
	result := make(map[float64][]float64)
	for _, taxRate := range taxRate {
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)

		}
		result[taxRate] = taxIncludedPrices
	}
	fmt.Println(result)
}
