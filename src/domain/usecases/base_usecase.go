package usecases

import (
	"app/src/domain/utils"
	"app/src/infra/database"
	"app/src/infra/repositories"
	"app/src/validation"
)

type input struct {
	name string
}

type output struct {
	message string
}

func BaseUsecase(transaction *database.Transaction, data input) (output, error) {
	var validators = []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().Property("test").Validators([]string{validation.ValidatorTypes.IsRequired}).Data(map[string]interface{}{"test": "test"}),
	}
	for _, validator := range validators {
		error := validator.Validate()
		if error != nil {
			return output{}, error
		}
	}

	var repository = repositories.BaseRepository
	var id = utils.GenerateUuid()
	var name = data.name

	repository.Insert(transaction, []any{id, name})

	return output{message: data.name + " created successfully with id: " + id}, nil
}
