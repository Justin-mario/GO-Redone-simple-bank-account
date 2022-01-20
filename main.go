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
	_, err := fmt.Scan(&userChoice)
	if err != nil {
		return
	}
	switch userChoice {
	case 1:
		fmt.Print("enter amount to deposit... ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			return
		}
		deposit(amount)
		if amount >= 0 {
			fmt.Printf("you deposited %v\n", amount)
		}

		main()
	case 2:
		fmt.Print("enter amount to withdraw... ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			return
		}
		withdraw(amount)
		if amount >= 0 || balance <= amount {
			fmt.Printf("you withdrew %v\n", amount)
		}
		main()
	case 3:
		fmt.Printf("your Balance is %v\n", balance)
		main()
	case 4:
		fmt.Print("Thank you for Banking With us")
		os.Exit(0)
	default:
		fmt.Print("unknown Command!")
		main()
	}
}

func message() {
	message := `Welcome to Our Bank
				Choose an Action to Perform
				1. Deposit.
				2. Withdraw.
				3. Balance.
				4.Exit`
	fmt.Println(message)

}

func deposit(amount float64) float64 {
	if amount < 0 {
		fmt.Println("Invalid amount")
	} else {
		balance += amount
	}

	return balance
}

func withdraw(amount float64) float64 {
	if amount < 0 || amount > balance {
		fmt.Println("invalid amount")

	} else {
		balance -= amount
	}

	return balance
}
