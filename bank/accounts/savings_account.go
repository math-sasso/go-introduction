package accounts

import (
	"matheus/bank/clients"
)

type SavingsAccount struct {
	Client                                 clients.Client
	NumberAgency, NumberAccount, Operation int
	balance                                float32
}

func (account *SavingsAccount) Withdraw(value float32) (string, float32) {
	allowed := value > 0 && value <= account.balance
	if allowed {
		account.balance -= value
		return "Withdraw done", account.balance
	} else {
		return "balance is not enough", account.balance
	}
}

func (account *SavingsAccount) Deposit(value float32) (string, float32) {
	if value > 0 {
		account.balance += value
		return "Deposit done", account.balance
	} else {
		return "Deposit did not worked", account.balance
	}
}
func (account *SavingsAccount) GetBalance() float32 {
	return account.balance
}
