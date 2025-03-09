package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
)

type CreateUsecaseInput struct {
	Name string `json:"name"`
}

type CreateUsecaseOutput struct {
	Message string `json:"message"`
}

var CreateUsecase = UseCase{
	Validators: []*utils.ValidatorBuilder{
		utils.NewValidatorBuilder().
			Property("name", "Name").
			Validators([]string{
				utils.ValidatorTypes.IsRequired,
				utils.ValidatorTypes.IsString,
			}),
	},
	Execute: func(transaction *database.Transaction, data dtos.DtoType) (dtos.DtoType, error) {
		var repository = repositories.BaseRepository
		var id = utils.GenerateUuid()
		var name = data["name"].(string)

		err := repository.Insert(transaction, []any{id, name})

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"message": name + " created successfully with id " + id}, nil
	},
}
