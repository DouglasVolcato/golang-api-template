package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type DeleteUsecaseInput struct {
	Id string `json:"id"`
}

type DeleteUsecaseOutput struct {
	Message string `json:"message"`
}

var DeleteUsecase = UseCase{
	Validators: []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().
			Property("id", "Id").
			Validators([]string{
				validation.ValidatorTypes.IsRequired,
				validation.ValidatorTypes.IsString,
			}),
	},
	Execute: func(transaction *database.Transaction, data dtos.DtoType) (dtos.DtoType, error) {
		var repository = repositories.BaseRepository
		var id = data["id"].(string)

		err := repository.Delete(transaction, id)

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"message": "Item deleted successfully"}, nil
	},
}
