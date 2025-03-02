package usecases

import (
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type BaseUsecaseInput struct {
	name string
}

type BaseUsecaseOutput struct {
	message string
}

var BaseUsecase = UseCase[BaseUsecaseInput, BaseUsecaseOutput]{
	Validators: []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().Property("test").Validators([]string{validation.ValidatorTypes.IsRequired}),
	},
	Handle: func(transaction *database.Transaction, data BaseUsecaseInput) (BaseUsecaseOutput, error) {
		var repository = repositories.BaseRepository
		var id = utils.GenerateUuid()
		var name = data.name

		repository.Insert(transaction, []any{id, name})

		return BaseUsecaseOutput{message: data.name + " created successfully with id: " + id}, nil
	},
}
