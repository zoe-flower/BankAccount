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
	err4 := account1.Deposit(1656)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	err5 := account2.Deposit(4631)
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	err3 := account2.Withdraw(-123)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println(account1.Balance, account2.Balance)
	err := transferFunds(account1, account2, 345)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account1.Balance, account2.Balance)
	err2 := transferFunds(account2, account1, 234)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(account1.Balance, account2.Balance)
	fmt.Println(account1.Transactions)
	fmt.Println(account2.Transactions)
}

type BankAccount struct {
	Balance      int
	AccountId    string
	AccountName  string
	AccountType  AccountType
	Transactions []Transaction

	//timeProviderThingy
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

func (ba *BankAccount) Deposit(depositAmount int) error {
	if depositAmount < 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.Balance += depositAmount
	err := ba.addTransaction(Deposit, depositAmount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ba *BankAccount) Withdraw(withdrawAmount int) error {
	if withdrawAmount >= 0 {
		return errors.New("deposit amount must be negative")
	}
	ba.Balance = ba.Balance + withdrawAmount
	err := ba.addTransaction(Withdraw, withdrawAmount)
	if err != nil {
		fmt.Println(err)
		return err
	}
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
	err2 := fromAccount.addTransaction(Transfer, -1*transferAmount)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	err := toAccount.addTransaction(Transfer, transferAmount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ba *BankAccount) addTransaction(transactionType TransactionType, amount int) error {
	// now:= ba.timeProviderThingy.Now()
	if transactionType == Deposit && amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	if transactionType == Withdraw && amount > 0 {
		return errors.New("withdrawal amount must be negative")
	}
	ba.Transactions = append(ba.Transactions, Transaction{transactionType, amount, time.Now().Local()})
	return nil
}

func (ba *BankAccount) viewTransactions() []Transaction {
	return ba.Transactions
}
