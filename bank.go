package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

)

func Bank() {
	var amountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("Sorry Cant Continue")
	}
	fmt.Println("Welcome To SunCafe Bank!")
	for {
		choice := menu()
		switch choice {
		case 1:
			fmt.Println("Your Amount Balance Is: ", amountBalance)
		case 2:
			var addMoney float64
			fmt.Print("How Much You Love To Add : ")
			fmt.Scan(&addMoney)
			if addMoney <= 0 {
				fmt.Println("Invalid Amount")
				continue
			}
			amountBalance += addMoney
			fmt.Println("Money Added Your new Amount Balance Is: ", amountBalance)
			writeBalanceToFile(amountBalance)
		case 3:
			var withdrawMoney float64
			fmt.Print("How Much You Love To Withdraw : ")
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 {
				fmt.Println("Invalid Amount")
				continue
			} else if withdrawMoney > amountBalance {
				fmt.Println("Yor Request Is More Than Your Balance")
				continue
			}
			amountBalance -= withdrawMoney
			fmt.Println("Money Added Your new Amount Balance Is: ", amountBalance)
			writeBalanceToFile(amountBalance)
		default:
			fmt.Println("Good By BaBy!")
			return
		}

	}
}

func menu() int {
	var menuChoose int
	fmt.Println("What Do You Want To Do ?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. EXIT")
	fmt.Print("4. Your Choice: ")
	fmt.Scan(&menuChoose)
	return menuChoose
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("can not find file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("can not parse balance")
	}
	return balance, nil

}
