package usecase

import "kafka-go-microservice/internal/domain/model"

type DebitAccountUseCase struct {
	Repository model.AccountRepository
}

type DebitAccountDTO struct {
	AccountId string
	Value     float64
}

func (uc *DebitAccountUseCase) Execute(dto DebitAccountDTO) error {
	account, err := uc.Repository.GetAccountById(dto.AccountId)
	if err != nil {
		return err
	}
	if err = account.Debit(dto.Value); err != nil {
		return err
	}
	uc.Repository.Save(account)
	return nil
}
