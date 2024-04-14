package model_test

import (
	"kafka-go-microservice/internal/domain/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewTransactionCredit(t *testing.T) {
	transaction := model.TransactionCredit{Value: -90}

	err := transaction.IsValid()

	require.Error(t, err)
}

func NewTransactionCreditValid(t *testing.T) {
	transaction := model.TransactionCredit{Value: 90}

	err := transaction.IsValid()

	require.Nil(t, err)
}

func NewTransactionDebit(t *testing.T) {
	transaction := model.TransactionCredit{Value: -90}

	err := transaction.IsValid()

	require.Nil(t, err)
}
