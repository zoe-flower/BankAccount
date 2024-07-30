package main

import (
	"testing"
)

func TestOpenAccount(t *testing.T) {
	tests := []struct {
		accountName         string
		accountType         accountType
		expectedBankAccount *BankAccount
	}{
		{"Zoe Flower", "savings", &BankAccount{0, 1, "Zoe Flower", "savings"}},
	}
	for _, tt := range tests {
		t.Run(tt.accountName, func(t *testing.T) {
			actualBankAccount := openAccount(tt.accountName, tt.accountType)
			if actualBankAccount.accountName != tt.expectedBankAccount.accountName {
				t.Errorf("expected accountName: %s, got: %s", tt.expectedBankAccount.accountName, actualBankAccount.accountName)
			}
			if actualBankAccount.accountType != tt.expectedBankAccount.accountType {
				t.Errorf("expected accountType: %s, got: %s", tt.expectedBankAccount.accountType, actualBankAccount.accountType)
			}
			if actualBankAccount.balance != tt.expectedBankAccount.balance {
				t.Errorf("expected balance: %d, got: %d", tt.expectedBankAccount.balance, actualBankAccount.balance)
			}
			if actualBankAccount.accountId != tt.expectedBankAccount.accountId {
				t.Errorf("expected accountId: %d, got: %d", tt.expectedBankAccount.accountId, actualBankAccount.accountId)
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
		{"Valid deposit", 5, false, 5},
		{"Invalid deposit (negative amount)", -5, true, 0},
		{"Valid deposit (zero amount)", 0, false, 0},
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
		{"Valid withdraw amount", -5, false, -5},
		{"Invalid withdraw amount (positive amount)", 5, true, 0},
		{"Invalid deposit (zero amount)", 0, true, 0},
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
	tests := []struct {
		name                    string
		account1                BankAccount
		account2                BankAccount
		expectedAccount1Balance int
		expectedAccount2Balance int
		transferAmount          int
		expectedError           bool
	}{
		{"Valid transfer amount", *account1, *account2, 0, 0, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := transferFunds(*account1, *account2, tt.transferAmount)
			//function yet to be written, but will reduce balance from acc1 and add to acc2, dummy amounts atm
			// function will set the balance for accounts which I will check

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
