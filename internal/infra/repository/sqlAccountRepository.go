package repository

import (
	"database/sql"
	"kafka-go-microservice/internal/domain/model"
)

type SqlAccountRepository struct {
	DB *sql.DB
}

func (r *SqlAccountRepository) Save(account *model.Account) error {
	_, err := r.DB.Exec("INSERT INTO account (holder_name, id) VALUES (?, ?)", account.HolderName, account.Base.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqlAccountRepository) GetByAccountId(accountId string) (*model.Account, error) {
	var account model.Account
	err := r.DB.QueryRow("SELECT id, holder_name FROM accounts WHERE id = ? LIMIT 1", accountId).Scan(account)

	if err != nil {
		return nil, err
	}

	return &account, err
}
