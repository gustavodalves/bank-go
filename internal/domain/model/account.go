package model

import "errors"

type Account struct {
	Base         Base
	HolderName   string
	Transactions []Transaction `json:"transactions"`
}

type AccountRepository interface {
	GetAccountById(id string) (*Account, error)
	Save(t Account) error
}

func (a *Account) GetBalance() float64 {
	total := 0.0

	for _, transaction := range a.Transactions {
		transaction.Calculate(&total)
	}

	return total
}

func NewAccount(holderName string, transactions []Transaction) Account {
	return Account{
		Transactions: transactions,
		HolderName:   holderName,
	}
}

func OpenAccount(holderName string) *Account {
	return &Account{
		Base:       *NewBase(),
		HolderName: holderName,
	}
}

func (a *Account) Debit(value float64) error {
	transaction := &TransactionDebit{
		Value: value,
	}

	if balance := a.GetBalance() - value; balance < 0 {
		return errors.New("")
	}

	isValid := transaction.IsValid()

	if isValid != nil {
		return isValid
	}

	a.Transactions = append(a.Transactions, transaction)
	return nil
}

func (a *Account) Credit(value float64) error {
	transaction := &TransactionCredit{
		Value: value,
	}
	isValid := transaction.IsValid()

	if isValid != nil {
		return isValid
	}

	a.Transactions = append(a.Transactions, transaction)
	return nil
}
