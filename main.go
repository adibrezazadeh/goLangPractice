package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmount, backRate, year, total float64
	investAmount = getUserInput("Enter Invest Amount: ")
	backRate = getUserInput("Enter Back Rate: ")
	year = getUserInput("Enter Number Years : ")
	total = calculateTotal(investAmount, backRate, year)
	fmt.Printf("TOTAL INVESTMENT AFTER %v YEAR IS %.2f ", year, total)
}

func calculateTotal(investAmount, backRate, year float64) float64 {
	return investAmount * math.Pow((1+backRate/100), year)
}
func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}
