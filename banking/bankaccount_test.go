package main

import (
	"testing"
)

func TestOpenAccount(t *testing.T) {
	tests := []struct {
		accountName         string
		accountType         AccountType
		expectedBankAccount *BankAccount
	}{
		{"Zoe Flower", "savings", &BankAccount{0, 1, "Zoe Flower", "savings", []Transactions{}}},
	}
	for _, tt := range tests {
		t.Run(tt.accountName, func(t *testing.T) {
			actualBankAccount := openAccount(tt.accountName, tt.accountType)
			if actualBankAccount.AccountName != tt.expectedBankAccount.AccountName {
				t.Errorf("expected accountName: %s, got: %s", tt.expectedBankAccount.AccountName, actualBankAccount.AccountName)
			}
			if actualBankAccount.AccountType != tt.expectedBankAccount.AccountType {
				t.Errorf("expected accountType: %s, got: %s", tt.expectedBankAccount.AccountType, actualBankAccount.AccountType)
			}
			if actualBankAccount.Balance != tt.expectedBankAccount.Balance {
				t.Errorf("expected balance: %d, got: %d", tt.expectedBankAccount.Balance, actualBankAccount.Balance)
			}
			if actualBankAccount.AccountId != tt.expectedBankAccount.AccountId {
				t.Errorf("expected accountId: %d, got: %d", tt.expectedBankAccount.AccountId, actualBankAccount.AccountId)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	account := openAccount("Zoe Flower", "savings")

	tests := []struct {
		name            string
		date            int
		depositAmount   int
		expectedError   bool
		expectedBalance int
	}{
		{"Valid deposit", 0, 5, false, 5},
		{"Invalid deposit (negative amount)", 1, -5, true, 0},
		{"Valid deposit (zero amount)", 1, 0, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.deposit(tt.depositAmount, tt.date)
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err != nil)
			}
			if account.checkBalance() != tt.expectedBalance {
				t.Errorf("Expected balance: %d, got: %d", tt.expectedBalance, account.checkBalance())
			}
		})
	}
}

func TestWithdraw(t *testing.T) {
	account := openAccount("Zoe Flower", "current")

	tests := []struct {
		name            string
		date            int
		withdrawAmount  int
		expectedError   bool
		expectedBalance int
	}{
		{"Valid withdraw amount", 1, -5, false, -5},
		{"Invalid withdraw amount (positive amount)", 1, 5, true, 0},
		{"Invalid deposit (zero amount)", 2, 0, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.withdraw(tt.withdrawAmount, tt.date)
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err != nil)
			}
			if account.checkBalance() != tt.expectedBalance {
				t.Errorf("Expected balance: %d, got: %d", tt.expectedBalance, account.checkBalance())
			}
		})
	}
}

func TestTransferFunds(t *testing.T) {
	account1 := openAccount("Zoe Flower", "savings")
	account2 := openAccount("Zoe Flower", "current")
	tests := []struct {
		name                    string
		initialBalanceAccount1  int
		initialBalanceAccount2  int
		expectedAccount1Balance int
		expectedAccount2Balance int
		transferAmount          int
		expectedError           bool
	}{
		{"Valid transfer amount", 50, 20, 40, 30, 10, false},
		{"Insufficient funds", 5, 20, 5, 20, 10, true},
		{"Transfer amount is zero", 30, 30, 30, 30, 0, false},
		{"Negative transfer amount", 30, 30, 30, 30, -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			account1.Balance = tt.initialBalanceAccount1
			account2.Balance = tt.initialBalanceAccount2

			err := transferFunds(account1, account2, tt.transferAmount)

			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}

			if got := account1.checkBalance(); got != tt.expectedAccount1Balance {
				t.Errorf("Expected balance for account1: %d, got: %d", tt.expectedAccount1Balance, got)
			}
			if got := account2.checkBalance(); got != tt.expectedAccount2Balance {
				t.Errorf("Expected balance for account2: %d, got: %d", tt.expectedAccount2Balance, got)
			}
		})
	}
}

//create goland project from first principles, add dependency.
// handle dev-tool go run main.go
// echo labstack middleware cookbook

func TestViewTransactionHistory(t *testing.T) {
	account := openAccount("Zoe Flower", "current")
	// do I then run a few deposit and withdraw functions? Or add the transactions in manually to set up the test?
	tests := []struct {
		name          string
		date          int
		expectedError bool
	}{
		{"Valid date with transaction", 2, false},
		{"Multiple transactions", 1, false},
		{"Future date", -1, true},
		{"No transaction on date ", -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.viewTransactionHistory(tt.withdrawAmount)
			if (err != nil) != tt.expectedError {
				return
			}
		})
	}
}

// need to add date to withdraw/deposit etc?
// how to I attach/save transactions to an account. Have a transactions struct? I append details on each transaction?
// dummy dates for now
