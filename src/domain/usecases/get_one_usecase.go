package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/abstract/entities"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type GetOneUsecaseInput struct {
	Id string `json:"id"`
}

type GetOneUsecaseOutput struct {
	Data []entities.BaseEntity `json:"data"`
}

var GetOneUsecase = UseCase{
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

		result, err := repository.Select(transaction, id)

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"data": result}, nil
	},
}
