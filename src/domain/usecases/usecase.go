package usecases

import (
	"app/src/infra/database"
	"app/src/validation"
)

type UseCase[I any, O any] struct {
	Validators []*validation.ValidatorBuilder
	Handle     func(transaction *database.Transaction, data I) (O, error)
}

func (usecase *UseCase[I, O]) Execute(transaction *database.Transaction, data I) (O, error) {
	var validators = []*validation.ValidatorBuilder{
		validation.NewValidatorBuilder().Property("test").Validators([]string{validation.ValidatorTypes.IsRequired}),
	}
	for _, validator := range validators {
		error := validator.Data(data).Validate()
		if error != nil {
			var zero O
			return zero, error
		}
	}

	return usecase.Execute(transaction, data)
}
