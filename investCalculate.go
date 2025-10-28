package main

import (
	"errors"
	"fmt"
	"math"

)

func InvestCalculate() {
	investAmount, err := getUserInput("Enter Invest Amount: ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("sorry we cant continue")
	}
	backRate, err2 := getUserInput("Enter Back Rate: ")
	if err2 != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("sorry we cant continue")
	}
	year, err3 := getUserInput("Enter Number Years : ")
	if err3 != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("sorry we cant continue")
	}
	total := calculateTotal(investAmount, backRate, year)
	fmt.Printf("TOTAL INVESTMENT AFTER %v YEAR IS %.2f ", year, total)
}

func calculateTotal(investAmount, backRate, year float64) float64 {
	return investAmount * math.Pow((1+backRate/100), year)
}
func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("the value must be more than zero")
	}
	return userInput, nil
}
