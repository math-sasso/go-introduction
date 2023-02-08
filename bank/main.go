package main

import (
	"fmt"
	"matheus/bank/accounts"
	"matheus/bank/clients"
)

type verifyAccount interface {
	Withdraw(value float32) (string, float32)
}

func PayBill(account verifyAccount, billValue float32) {
	account.Withdraw(billValue)
}

func main() {
	contaDoBruno := accounts.CheckingAccount{Client: clients.Client{
		Name: "Bruno",
		ID:   "123456",
		Job:  "Desenvolvedor"},
		NumberAgency: 123, NumberAccount: 123456}

	contaDoBruno.Deposit(100)

	fmt.Println(contaDoBruno.GetBalance())

	contaDoPedro := accounts.SavingsAccount{Client: clients.Client{
		Name: "Pedro",
		ID:   "1234568",
		Job:  "Desenvolvedor"},
		NumberAgency: 123, NumberAccount: 123456}

	contaDoPedro.Deposit(200)

	fmt.Println(contaDoPedro.GetBalance())

	PayBill(&contaDoBruno, 10)

	fmt.Println(contaDoBruno.GetBalance())

}
