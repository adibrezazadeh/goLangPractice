package main

import "fmt"

func Bank() {
	var amountBalance float64 = 0
	choice := menu()
	if choice == 1 {
		fmt.Println("Your Amount Balance Is: ", amountBalance)
		menu()
	}
}

func menu() int {
	var menuChoose int
	fmt.Println("Welcome To SunCafe Bank!")
	fmt.Println("What Do You Want To Do ?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. EXIT")
	fmt.Print("4. Your Choice: ")
	fmt.Scan(&menuChoose)
	return menuChoose
}
