package usecase

import "kafka-go-microservice/internal/domain/model"

type CreditAccountUseCase struct {
	Repository model.AccountRepository
}

type CreditAccountDTO struct {
	AccountId string
	Value     float64
}

func (uc *CreditAccountUseCase) Execute(dto CreditAccountDTO) error {
	account, err := uc.Repository.GetAccountById(dto.AccountId)
	if err != nil {
		return err
	}
	if err = account.Credit(dto.Value); err != nil {
		return err
	}
	uc.Repository.Save(account)
	return nil
}
