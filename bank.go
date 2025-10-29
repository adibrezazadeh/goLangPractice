package main

import (
	"example.com/first/fileOs"
	"fmt"
)

func Bank() {
	var amountBalance, err = fileos.ReadFloatFromFile("balance.txt")
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
			fileos.WriteFloatToFile("balance.txt", amountBalance)
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
			fileos.WriteFloatToFile("balance.txt", amountBalance)
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
