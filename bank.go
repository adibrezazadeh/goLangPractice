package main

import "fmt"

func Bank() {
	menu()
}

func menu() int {
	var menuChoose int
	fmt.Println("Welcome To SunCafe Bank!")
	fmt.Println("What Do You Want To Do ?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. EXIT")
	fmt.Scan(&menuChoose)
	return menuChoose
}
