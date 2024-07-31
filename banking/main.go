package main

import (
	"errors"
)

func main() {
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

func transferFunds(fromAccount, toAccount *BankAccount, transferAmount int) error {
	if transferAmount < 0 {
		return errors.New("transfer amount must be positive")
	}
	if transferAmount > fromAccount.balance {
		return errors.New("Insufficient funds")
	}
	fromAccount.balance -= transferAmount
	toAccount.balance += transferAmount
	return nil
}
