package main

import (
	"testing"
)

func TestOpenAccount(t *testing.T) {
	tests := []struct {
		name                string
		accountName         string
		accountType         AccountType
		expectedBankAccount *BankAccount
	}{
		{name: "Happy Path",
			accountName:         "Zoe Flower",
			accountType:         "savings",
			expectedBankAccount: &BankAccount{0, "1", "Zoe Flower", "savings"}},
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
		})
	}
}

func TestDeposit(t *testing.T) {
	account := openAccount("Zoe Flower", "savings")

	tests := []struct {
		name            string
		depositAmount   int
		expectedError   bool
		expectedBalance int
	}{
		{name: "Valid deposit",
			depositAmount:   5,
			expectedError:   false,
			expectedBalance: 5},
		{name: "Invalid deposit (negative amount)",
			depositAmount:   -5,
			expectedError:   true,
			expectedBalance: 0},
		{name: "Valid deposit (zero amount)",
			depositAmount:   0,
			expectedError:   false,
			expectedBalance: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.deposit(tt.depositAmount)
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
		withdrawAmount  int
		expectedError   bool
		expectedBalance int
	}{
		{name: "Valid withdraw amount",
			withdrawAmount:  -5,
			expectedError:   false,
			expectedBalance: -5},
		{name: "Invalid withdraw amount (positive amount)",
			withdrawAmount:  5,
			expectedError:   true,
			expectedBalance: 0},
		{name: "Invalid deposit (zero amount)",
			withdrawAmount:  0,
			expectedError:   true,
			expectedBalance: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.withdraw(tt.withdrawAmount)
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
	var tests = []struct {
		name                    string
		initialBalanceAccount1  int
		initialBalanceAccount2  int
		expectedAccount1Balance int
		expectedAccount2Balance int
		transferAmount          int
		expectedError           bool
	}{
		{name: "Valid transfer amount",
			initialBalanceAccount1:  50,
			initialBalanceAccount2:  20,
			expectedAccount1Balance: 40,
			expectedAccount2Balance: 30,
			transferAmount:          10,
			expectedError:           false},
		{name: "Insufficient funds",
			initialBalanceAccount1:  5,
			initialBalanceAccount2:  20,
			expectedAccount1Balance: 5,
			expectedAccount2Balance: 20,
			transferAmount:          10,
			expectedError:           true},
		{name: "Transfer amount is zero",
			initialBalanceAccount1:  30,
			initialBalanceAccount2:  30,
			expectedAccount1Balance: 30,
			expectedAccount2Balance: 30,
			transferAmount:          0,
			expectedError:           false},
		{
			name:                    "Negative transfer amount",
			initialBalanceAccount1:  30,
			initialBalanceAccount2:  30,
			expectedAccount1Balance: 30,
			expectedAccount2Balance: 30,
			transferAmount:          -10,
			expectedError:           true,
		},
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
