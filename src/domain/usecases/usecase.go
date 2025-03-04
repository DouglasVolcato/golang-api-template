package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/infra/database"
	"app/src/validation"
)

type UseCase struct {
	Validators []*validation.ValidatorBuilder
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
