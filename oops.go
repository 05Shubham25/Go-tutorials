package main

import (
	"fmt"
)

// BankAccount struct
type BankAccount struct {
	accountHolder string
	balance       float64 // Private (unexported, lowercase)
}

// Constructor function
func NewBankAccount(holder string, initialBalance float64) *BankAccount {
	return &BankAccount{accountHolder: holder, balance: initialBalance}
}

// Method to deposit money
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
		fmt.Println("Deposited:", amount)
	} else {
		fmt.Println("Invalid deposit amount")
	}
}

// Method to withdraw money (controlled access)
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
		fmt.Println("Withdrawn:", amount)
	} else {
		fmt.Println("Insufficient balance or invalid amount")
	}
}

// Getter method (Read-Only access)
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	// Creating a bank account
	account := NewBankAccount("Alice", 1000)

	// Accessing the private field using a method
	fmt.Println("Initial Balance:", account.GetBalance())

	// Deposit and Withdraw
	account.Deposit(500)
	account.Withdraw(300)
	fmt.Println("Final Balance:", account.GetBalance())

	// Trying to access balance directly (not allowed)
	fmt.Println("Trying to access balance directly",account.balance) // Error: balance is private
}
