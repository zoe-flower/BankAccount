package main

import (
	"errors"
	"fmt"
)

func main() {
	newAccount := openAccount("Zoe Flower", Current)
	fmt.Println(newAccount)
	newAccount.deposit(5, 1)
	fmt.Println(newAccount.Transactions)
	newAccount.withdraw(-1, 2)
	fmt.Println(newAccount.Transactions)
	fmt.Println(newAccount.Balance)

}

type BankAccount struct {
	Balance      int
	AccountId    int
	AccountName  string
	AccountType  AccountType
	Transactions []Transactions
}

type AccountType string

const (
	Savings AccountType = "Savings"
	Current AccountType = "Current"
	Credit  AccountType = "Credit"
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
		Balance:      0,
		AccountId:    1,
		AccountName:  accountName,
		AccountType:  accountType,
		Transactions: []Transactions{},
	}
}

func (ba *BankAccount) deposit(depositAmount int, date int) error {
	if depositAmount < 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.Balance += depositAmount
	ba.Transactions = append(ba.Transactions, Transactions{
		Date:            date,
		Amount:          depositAmount,
		TransactionType: Deposit,
	})
	return nil
}

func (ba *BankAccount) withdraw(withdrawAmount int, date int) error {
	if withdrawAmount >= 0 {
		return errors.New("deposit amount must be negative")
	} //not working?
	ba.Balance = ba.Balance + withdrawAmount
	ba.Transactions = append(ba.Transactions, Transactions{
		Date:            date,
		Amount:          withdrawAmount,
		TransactionType: Withdrawal,
	})
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

// func viewTransactionHistory(account *BankAccount) *BankAccount {
// 	return nil
// }
