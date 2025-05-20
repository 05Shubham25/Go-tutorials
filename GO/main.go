package main

import (
	"fmt" // Import the bank package (replace with your actual project path)
)

func main() {
	// Creating a bank account
	account := NewBankAccount("Alice", 1000)

	// Access balance using getter
	fmt.Println("Initial Balance:", account.GetBalance())

	// Deposit and Withdraw
	account.Deposit(500)
	account.Withdraw(300)
	fmt.Println("Final Balance:", account.GetBalance())

	// Trying to access balance directly (this will cause an error)
	fmt.Println(account.balance) // ❌ ERROR: balance is not accessible outside the `bank` package
}
