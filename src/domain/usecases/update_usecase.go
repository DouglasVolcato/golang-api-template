package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
)

type UpdateUsecaseInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUsecaseOutput struct {
	Message string `json:"message"`
}

var UpdateUsecase = UseCase{
	Validators: []*utils.ValidatorBuilder{
		utils.NewValidatorBuilder().
			Property("name", "Name").
			Validators([]string{
				utils.ValidatorTypes.IsRequired,
				utils.ValidatorTypes.IsString,
			}),
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
		var name = data["name"].(string)

		err := repository.Update(transaction, id, []any{name})

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"message": name + " updated successfully with id " + id}, nil
	},
}
