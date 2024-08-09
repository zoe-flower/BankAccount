package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	newAccount := openAccount("Zoe Flower", Current)
	fmt.Println(newAccount)
}

type BankAccount struct {
	Balance      int
	AccountId    string
	AccountName  string
	AccountType  AccountType
	Transactions []Transaction
}

type AccountType string

const (
	Savings AccountType = "Savings"
	Current AccountType = "Current"
	Credit  AccountType = "Credit"
)

type Transaction struct {
	TransactionType TransactionType
	amount          int
	date            time.Time
}

type TransactionType string

const (
	Deposit  TransactionType = "Deposit"
	Withdraw TransactionType = "Withdraw"
	Transfer TransactionType = "Transfer"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func openAccount(accountName string, accountType AccountType) *BankAccount {
	return &BankAccount{
		Balance:     0,
		AccountId:   GenerateUUID(),
		AccountName: accountName,
		AccountType: accountType,
	}
}

func (ba *BankAccount) deposit(depositAmount int) error {
	if depositAmount < 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.Balance += depositAmount
	ba.addTransaction(Deposit, depositAmount)
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
		return errors.New("insufficient funds")
	}
	fromAccount.Balance -= transferAmount
	toAccount.Balance += transferAmount
	return nil
}

func (ba *BankAccount) addTransaction(transationType TransactionType, amount int) {
	ba.Transactions = append(ba.Transactions, Transaction{transationType, amount, time.Now()})
}
