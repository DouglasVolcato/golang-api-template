package controllers

import (
	"app/src/domain/usecases"
	"app/src/infra/database"
	"database/sql"
)

type Controller[I any, O any] struct {
	usecase *usecases.UseCase[I, O]
}

func NewController[I any, O any](usecase *usecases.UseCase[I, O]) *Controller[I, O] {
	return &Controller[I, O]{usecase: usecase}
}

func (controller *Controller[I, O]) Execute(databaseConnection *sql.DB, data I) (O, error) {
	var transaction = database.NewTransaction(databaseConnection)
	var err error

	err = transaction.BeginTransaction()
	if err != nil {
		var empty O
		return empty, err
	}

	response, err := controller.usecase.Execute(transaction, data)
	if err != nil {
		transaction.RollbackTransaction()
		var empty O
		return empty, err
	}

	err = transaction.CommitTransaction()
	if err != nil {
		var empty O
		return empty, err
	}

	return response, nil
}
