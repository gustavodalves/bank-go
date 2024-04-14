package model

import "errors"

type Transaction interface {
	Calculate(total *float64)
}

type TransactionDebit struct {
	Base  Base
	Value float64
}

type TransactionCredit struct {
	Base  Base
	Value float64
}

func (t *TransactionCredit) IsValid() error {
	if t.Value < 0 {
		return errors.New("")
	}

	return nil
}

func (t *TransactionDebit) IsValid() error {
	if t.Value < 0 {
		return errors.New("")
	}

	return nil
}

func (t *TransactionDebit) Calculate(total *float64) {
	*total -= t.Value
}

func (t *TransactionCredit) Calculate(total *float64) {
	*total += t.Value
}
