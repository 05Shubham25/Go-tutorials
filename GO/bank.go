package main

import "fmt"

type BankAccount struct {
	accountHolder string
	balance       float64 // Private field (not accessible outside this package)
}

// Constructor function to create a new account
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

// Method to withdraw money
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
		fmt.Println("Withdrawn:", amount)
	} else {
		fmt.Println("Insufficient balance or invalid amount")
	}
}

// Getter method to access balance (Read-Only)
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}
