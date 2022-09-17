package main

import "fmt"

type BankAccount struct {
	id            string
	typeOfAccount string
	balance       float64
}

func CreateAccount(typeOfAccount string) *BankAccount {
	return &BankAccount{
		id:            "1",
		typeOfAccount: typeOfAccount,
	}
}

type Customer struct {
	id   string
	name string
}

func CreateCustomer(name string) *Customer {
	return &Customer{
		id:   "1",
		name: name,
	}
}

type Transaction struct {
	id       string
	amount   float64
	sender   *BankAccount
	receiver *BankAccount
}

func CreateTransaction(sender *BankAccount, receiver *BankAccount, amount float64) *Transaction {
	return &Transaction{
		id:       "1",
		amount:   amount,
		sender:   sender,
		receiver: receiver,
	}
}

type BankTeller struct {
	account     *BankAccount
	customer    *Customer
	transaction *Transaction
}

func NewBankTeller() *BankTeller {
	return &BankTeller{}
}

func (b *BankTeller) CreateAccount(typeOfAccount string) {
	b.customer = CreateCustomer("John Doe")
	b.account = CreateAccount(typeOfAccount)
}

func (b *BankTeller) Transfer(receiver *BankAccount, amount float64) {
	b.transaction = CreateTransaction(b.account, receiver, amount)
}

func main() {
	bankTeller := NewBankTeller()
	bankTeller.CreateAccount("Checking")
	bankTeller.Transfer(CreateAccount("Savings"), 100)
	fmt.Printf("%+v\n", *bankTeller.account)
	fmt.Printf("%+v\n", *bankTeller.customer)
	fmt.Printf("%+v\n", *bankTeller.transaction)
}
