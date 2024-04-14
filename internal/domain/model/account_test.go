package model_test

import (
	"kafka-go-microservice/internal/domain/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	holderName := "Gustavo"

	account := model.OpenAccount(holderName)

	const expected float64 = 0

	require.Equal(t, account.HolderName, holderName)
	require.Equal(t, account.GetBalance(), expected)
}

func TestGetBalance(t *testing.T) {
	holderName := "Gustavo"

	account := model.OpenAccount(holderName)

	account.Credit(90)

	require.Equal(t, account.GetBalance(), float64(90))
}

func ValidDebit(t *testing.T) {
	holderName := "Gustavo"

	account := model.OpenAccount(holderName)

	isValid := account.Debit(90)

	require.Nil(t, isValid)
}
