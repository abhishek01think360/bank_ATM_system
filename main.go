package main

import (
	"fmt"
)

type account struct {
	username  string
	password  string
	balance   int
	statement []string
}

var idx int = 0
var account_list = [10]account{}

func register() {
	var user string
	var pass string

	fmt.Println("UserName:")
	fmt.Scanln(&user)
	fmt.Println("Password:")
	fmt.Scanln(&pass)
	var temp account
	temp.username = user
	temp.password = pass
	temp.balance = 100
	account_list[idx] = temp
	idx++
	fmt.Println(account_list[idx-1])
	fmt.Println("registration sucessful")
}
func account_exist(user_id string, pass string) (i int) {
	for i := 0; i < len(account_list); i++ {
		if (account_list[i].username == user_id) && (account_list[i].password == pass) {
			return i
		}
	}
	return -1
}

// add_new_age := append(student_age, 25)
func check_balance() {
	var user_id string
	var pass string
	fmt.Println("Enter User Id")
	fmt.Scanln(&user_id)
	fmt.Println("Enter Password")
	fmt.Scanln(&pass)
	var i = account_exist(user_id, pass)
	if i != -1 {
		fmt.Println(account_list[i].balance)
		return
	}
	fmt.Println("Account Not Found")
	return
}
func amount_withdraw() {
	var user_id string
	var pass string
	var amt int
	fmt.Println("Enter User Id")
	fmt.Scanln(&user_id)
	fmt.Println("Enter Password")
	fmt.Scanln(&pass)
	fmt.Println("Enter Amount")
	fmt.Scanln(&amt)
	var i int = account_exist(user_id, pass)
	if account_list[i].balance >= amt {
		account_list[i].balance = account_list[i].balance - amt
		// var trans string = string(amt) + " debited.avaible balance " + " " + string(account_list[i].balance)
		// account_list[i].statement = append(account_list[i].statement, trans)
		fmt.Println("Transaction Sucessfull")
	} else {
		fmt.Println("amount Insufficient.")
	}
}
func statement_print() {
	var user_id string
	var pass string
	fmt.Println("Enter User Id")
	fmt.Scanln(&user_id)
	fmt.Println("Enter Password")
	fmt.Scanln(&pass)
	var i = account_exist(user_id, pass)
	for k := len(account_list[i].statement) - 5; k < len(account_list[i].statement); k++ {
		fmt.Println(account_list[i].statement[k])
	}
}

func main() {
	var respose string
	var q string = "a"
	for q != "e" {
		fmt.Println("please enter number accordingly:")
		fmt.Println("a for new users")
		fmt.Println("b for Check balance")
		fmt.Println("c for amount withdraw")
		fmt.Println("d for Transacton")
		fmt.Println("e exist program")
		fmt.Scanln(&respose)
		if respose == "a" {
			register()
		} else if respose == "b" {
			check_balance()
		} else if respose == "c" {
			amount_withdraw()
		} else if respose == "d" {
			statement_print()
		} else if respose == "e" {
			q = "e"
		} else {
			fmt.Println("Entered Number is invalid!")
		}
	}
}
