package accounts

import (
	"matheus/bank/clients"
)

type CheckingAccount struct {
	Client                      clients.Client
	NumberAgency, NumberAccount int
	balance                     float32
}

func (account *CheckingAccount) Withdraw(value float32) (string, float32) {
	allowed := value > 0 && value <= account.balance
	if allowed {
		account.balance -= value
		return "Withdraw done", account.balance
	} else {
		return "balance is not enough", account.balance
	}
}

func (account *CheckingAccount) Deposit(value float32) (string, float32) {
	if value > 0 {
		account.balance += value
		return "Deposit done", account.balance
	} else {
		return "Deposit did not worked", account.balance
	}
}

func (originCheckingAccount *CheckingAccount) Transfer(value float32, destinationCheckingAccount *CheckingAccount) bool {
	if value < originCheckingAccount.balance && value > 0 {
		originCheckingAccount.Withdraw(value)
		destinationCheckingAccount.Deposit(value)
		return true
	} else {
		return false
	}
}

func (account *CheckingAccount) GetBalance() float32 {
	return account.balance
}
