package main

import (
	"fmt"
	"os"
)

var balance float64

func main() {
	var userChoice uint
	var amount float64
	message()
	fmt.Print("input your choice.... ")
	fmt.Scan(&userChoice)
	switch userChoice {
	case 1:
		fmt.Print("enter amount to deposit... ")
		fmt.Scan(&amount)
		deposit(amount)
		fmt.Printf("you deposited %v your account balance is %v\n", amount, balance)
		main()
	case 2:
		fmt.Print("enter amount to withdraw... ")
		fmt.Scan(&amount)
		withdraw(amount)
		fmt.Printf("you withdrew %v your account balance is %v\n", amount, balance)
		main()

	case 3:
		fmt.Print("Thank you for Banking With us")
		os.Exit(0)
	default:
		fmt.Print("unknown Command!")
	}
}

func message() {
	message := `Welcome to Our Bank
				Choose an Action to Perform
				1. Deposit.
				2. Withdraw.
				3. Exit.`
	fmt.Println(message)

}

func deposit(amount float64) float64 {
	if amount < 0 {
		fmt.Println("Invalid amount")
	}
	balance += amount
	return balance
}

func withdraw(amount float64) float64 {
	if amount < 0 || amount > balance {
		fmt.Println("invalid amount")
	}
	balance -= amount
	return balance
}
