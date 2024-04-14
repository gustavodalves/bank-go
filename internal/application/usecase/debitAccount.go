package usecase

import "kafka-go-microservice/internal/domain/model"

type DebitAccountUseCase struct {
	repository model.AccountRepository
}

type DebitAccountDTO struct {
	accountId string
	value     float64
}

func (uc *DebitAccountUseCase) Execute(dto DebitAccountDTO) error {
	account, err := uc.repository.GetAccountById(dto.accountId)
	if err != nil {
		return err
	}
	if err = account.Debit(dto.value); err != nil {
		return err
	}
	uc.repository.Save(*account)
	return nil
}
