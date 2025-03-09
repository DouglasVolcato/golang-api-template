package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/domain/abstract/entities"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"strconv"
)

type GetAllUsecaseInput struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type GetAllUsecaseOutput struct {
	Data []entities.BaseEntity `json:"data"`
}

var GetAllUsecase = UseCase{
	Validators: []*utils.ValidatorBuilder{
		utils.NewValidatorBuilder().
			Property("limit", "Limit").
			Validators([]string{
				utils.ValidatorTypes.IsRequired,
				utils.ValidatorTypes.IsInteger,
			}),
		utils.NewValidatorBuilder().
			Property("offset", "Offset").
			Validators([]string{
				utils.ValidatorTypes.IsRequired,
				utils.ValidatorTypes.IsInteger,
			}),
	},
	Execute: func(transaction *database.Transaction, data dtos.DtoType) (dtos.DtoType, error) {
		var repository = repositories.BaseRepository
		limit, _ := strconv.Atoi(data["limit"].(string))
		offset, _ := strconv.Atoi(data["offset"].(string))

		result, err := repository.SelectAll(transaction, limit, offset)

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"data": result}, nil
	},
}
