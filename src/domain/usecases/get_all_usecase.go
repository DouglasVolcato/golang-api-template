package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/abstract/entities"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type GetAllUsecaseInput struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type GetAllUsecaseOutput struct {
	Data []entities.BaseEntity `json:"data"`
}

var GetAllUsecase = UseCase{
	Validators: []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().
			Property("limit", "Limit").
			Validators([]string{
				validation.ValidatorTypes.IsRequired,
				validation.ValidatorTypes.IsInteger,
			}),
		validation.NewValidatorBuilder().
			Property("offset", "Offset").
			Validators([]string{
				validation.ValidatorTypes.IsRequired,
				validation.ValidatorTypes.IsInteger,
			}),
	},
	Execute: func(transaction *database.Transaction, data dtos.DtoType) (dtos.DtoType, error) {
		var repository = repositories.BaseRepository
		var limit = int(data["limit"].(float64))
		var offset = int(data["offset"].(float64))

		result, err := repository.SelectAll(transaction, limit, offset)

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"data": result}, nil
	},
}
