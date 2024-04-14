package usecase

import "kafka-go-microservice/internal/domain/model"

type CreditAccountUseCase struct {
	repository model.AccountRepository
}

type CreditAccountDTO struct {
	accountId string
	value     float64
}

func (uc *CreditAccountUseCase) Execute(dto CreditAccountDTO) error {
	account, err := uc.repository.GetAccountById(dto.accountId)
	if err != nil {
		return err
	}
	if err = account.Credit(dto.value); err != nil {
		return err
	}
	uc.repository.Save(*account)
	return nil
}
