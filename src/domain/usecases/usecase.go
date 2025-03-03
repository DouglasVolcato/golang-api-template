package usecases

import (
	"app/src/domain/abstract"
	"app/src/infra/database"
	"app/src/validation"
)

type UseCase struct {
	Validators []*validation.ValidatorBuilder
	Execute    func(transaction *database.Transaction, data abstract.DtoType) (abstract.DtoType, error)
}

func (usecase *UseCase) Validate(data abstract.DtoType) error {
	for _, validator := range usecase.Validators {
		error := validator.Data(data).Validate()
		if error != nil {
			return error
		}
	}
	return nil
}
