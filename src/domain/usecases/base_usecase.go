package usecases

import (
	"app/src/domain/abstract"
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

var BaseUsecase = UseCase{
	Validators: []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().Property("name").Validators([]string{validation.ValidatorTypes.IsRequired}),
	},
	Execute: func(transaction *database.Transaction, data abstract.DtoType) (abstract.DtoType, error) {
		var repository = repositories.BaseRepository
		var id = utils.GenerateUuid()
		var name = data["name"].(string)

		err := repository.Insert(transaction, []any{id, name})

		if err != nil {
			return nil, err
		}

		return abstract.DtoType{"message": name + " created successfully with id " + id}, nil
	},
}
