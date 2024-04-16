package usecase

import "kafka-go-microservice/internal/domain/model"

type GetAccountUseCase struct {
	Repository model.AccountRepository
}

type GetAccountDTO struct {
	AccountId string
}

func (us *GetAccountUseCase) Execute(dto GetAccountDTO) (*model.Account, error) {
	account, err := us.Repository.GetAccountById(dto.AccountId)
	if err != nil {
		return nil, err
	}
	return account, nil
}
