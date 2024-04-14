package usecase

import "kafka-go-microservice/internal/domain/model"

type OpenAccountUseCase struct {
	Repository model.AccountRepository
}

type OpenAccountDTO struct {
	HolderName string
}

func (uc *OpenAccountUseCase) Execute(dto OpenAccountDTO) error {
	account := model.OpenAccount(dto.HolderName)

	if error := uc.Repository.Save(account); error != nil {
		return error
	}

	return nil
}
