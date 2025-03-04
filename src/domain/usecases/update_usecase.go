package usecases

import (
	"app/src/domain/abstract/dtos"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type UpdateUsecaseInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUsecaseOutput struct {
	Message string `json:"message"`
}

var UpdateUsecase = UseCase{
	Validators: []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().
			Property("name", "Name").
			Validators([]string{
				validation.ValidatorTypes.IsRequired,
				validation.ValidatorTypes.IsString,
			}),
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
		var name = data["name"].(string)

		err := repository.Update(transaction, id, []any{name})

		if err != nil {
			return nil, err
		}

		return dtos.DtoType{"message": name + " updated successfully with id " + id}, nil
	},
}
