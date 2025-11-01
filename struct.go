package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDay  string
	createAt  time.Time
}

func Struct() {
	var userInfo user = user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthDay:  getUserData("Please enter your birthday: "),
		createAt:  time.Now(),
	}
	fmt.Println(userInfo.firstName, userInfo.lastName, userInfo.birthDay)
	userInfo.setUserInfoEmpty()
	fmt.Println(userInfo.firstName, userInfo.lastName, userInfo.birthDay)

}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
func (u *user) setUserInfoEmpty() { //set pointer to change exactly variable not change copy
	u.birthDay = ""
	u.firstName = ""
	u.lastName = ""

}
