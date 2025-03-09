package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/utils"
	"app/src/infra/database"
)

type UseCase struct {
	Validators []*utils.ValidatorBuilder
	Execute    func(transaction *database.Transaction, data dtos.DtoType) (dtos.DtoType, error)
}

func (usecase *UseCase) Validate(data dtos.DtoType) error {
	for _, validator := range usecase.Validators {
		error := validator.Data(data).Validate()
		if error != nil {
			return error
		}
	}
	return nil
}
