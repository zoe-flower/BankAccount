package main

import (
	"errors"
)

func main() {
}

type BankAccount struct {
	Balance      int
	AccountId    int
	AccountName  string
	AccountType  AccountType
	Transactions Transactions
}

type AccountType string

const (
	Savings AccountType = "savings"
	Current AccountType = "current"
	Credit  AccountType = "credit"
)

type Transactions struct {
	Date            int
	Amount          int
	TransactionType TransactionType
}

type TransactionType string

const (
	Deposit    TransactionType = "Deposit"
	Withdrawal TransactionType = "Withdrawal"
	Transfer   TransactionType = "Transfer"
)

func openAccount(accountName string, accountType AccountType) *BankAccount {
	return &BankAccount{
		Balance:     0,
		AccountId:   1,
		AccountName: accountName,
		AccountType: accountType,
	}
}

func (ba *BankAccount) deposit(depositAmount int) error {
	if depositAmount < 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.Balance += depositAmount
	return nil
}

func (ba *BankAccount) withdraw(withdrawAmount int) error {
	if withdrawAmount >= 0 {
		return errors.New("deposit amount must be negative")
	}
	ba.Balance = ba.Balance + withdrawAmount
	return nil
}

func (ba *BankAccount) checkBalance() int {
	return ba.Balance
}

func transferFunds(fromAccount, toAccount *BankAccount, transferAmount int) error {
	if transferAmount < 0 {
		return errors.New("transfer amount must be positive")
	}
	if transferAmount > fromAccount.Balance {
		return errors.New("Insufficient funds")
	}
	fromAccount.Balance -= transferAmount
	toAccount.Balance += transferAmount
	return nil
}

func viewTransactionHistory(account *BankAccount) *BankAccount {
	return nil
}
