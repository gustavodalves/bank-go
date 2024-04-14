package usecase

import "kafka-go-microservice/internal/domain/model"

type OpenAccountUseCase struct {
	repository model.AccountRepository
}

type OpenAccountDTO struct {
	holderName string
}

func (uc *OpenAccountUseCase) Execute(dto OpenAccountDTO) error {
	account := model.OpenAccount(dto.holderName)

	if error := uc.repository.Save(*account); error != nil {
		return error
	}

	return nil
}
