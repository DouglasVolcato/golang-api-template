package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/abstract/entities"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
)

type GetOneUsecaseInput struct {
	Id string `json:"id"`
}

type GetOneUsecaseOutput struct {
	Data []entities.BaseEntity `json:"data"`
}

var GetOneUsecase = UseCase{
	Validators: []*utils.ValidatorBuilder{
		utils.NewValidatorBuilder().
			Property("id", "Id").
			Validators([]string{
				utils.ValidatorTypes.IsRequired,
				utils.ValidatorTypes.IsString,
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
