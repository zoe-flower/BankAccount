package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	account1 := openAccount("Zoe Flower", Current)
	account2 := openAccount("Zoe Flower", Savings)
	fmt.Println(account1, account2)
	account1.deposit(1656)
	account2.deposit(4631)
	fmt.Println(account1.Balance, account2.Balance)
	err := transferFunds(account1, account2, 345)
	if err != nil {
		return
	}
	fmt.Println(account1.Balance, account2.Balance)
	err2 := transferFunds(account2, account1, 234)
	if err2 != nil {
		return
	}
	fmt.Println(account1.Balance, account2.Balance)
	fmt.Println(account1.Transactions)
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
	ba.addTransaction(Withdraw, withdrawAmount)
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
	fromAccount.addTransaction(Transfer, -1*transferAmount)
	toAccount.addTransaction(Transfer, transferAmount)
	return nil
}

func (ba *BankAccount) addTransaction(transactionType TransactionType, amount int) {
	ba.Transactions = append(ba.Transactions, Transaction{transactionType, amount, time.Now().Local()})
}

func (ba *BankAccount) viewTransactions() []Transaction {
	return ba.Transactions
}
