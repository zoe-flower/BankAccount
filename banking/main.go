package main

import (
	"errors"
	"fmt"
)

func main() {
	account1 := openAccount("Zoe Flower", "savings")
	account2 := openAccount("Zoe Flower", "current")
	fmt.Println(account1, account2)
}

type BankAccount struct {
	balance     int
	accountId   int
	accountName string
	accountType accountType
}

type accountType string

func openAccount(accountName string, accountType accountType) *BankAccount {
	return &BankAccount{
		balance:     0,
		accountId:   1,
		accountName: accountName,
		accountType: accountType,
	}
}

func (ba *BankAccount) deposit(depositAmount int) error {
	if depositAmount < 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.balance += depositAmount
	return nil
}

func (ba *BankAccount) withdraw(withdrawAmount int) error {
	if withdrawAmount >= 0 {
		return errors.New("deposit amount must be negative")
	}
	ba.balance = ba.balance + withdrawAmount
	return nil
}

func (ba *BankAccount) checkBalance() int {
	return ba.balance
}

func (ba *BankAccount) transferFunds(fromAccount, toAccount BankAccount) error {

	return nil
}
