package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
)

type DeleteUsecaseInput struct {
	Id string `json:"id"`
}

type DeleteUsecaseOutput struct {
	Message string `json:"message"`
}

var DeleteUsecase = UseCase{
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

		err := repository.Delete(transaction, id)

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"message": "Item deleted successfully"}, nil
	},
}
