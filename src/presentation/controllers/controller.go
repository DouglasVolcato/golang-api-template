package controllers

import (
	"app/src/domain/abstract"
	"app/src/domain/usecases"
	"app/src/infra/database"
	"database/sql"
)

type Controller struct {
	usecase *usecases.UseCase
}

func NewController(usecase *usecases.UseCase) *Controller {
	return &Controller{usecase: usecase}
}

func (controller *Controller) Execute(databaseConnection *sql.DB, data abstract.DtoType) (abstract.DtoType, error, int) {
	var transaction = database.NewTransaction(databaseConnection)
	var err error

	err = transaction.BeginTransaction()
	if err != nil {
		return nil, err, 500
	}

	err = controller.usecase.Validate(data)
	if err != nil {
		transaction.RollbackTransaction()
		return nil, err, 400
	}

	response, err := controller.usecase.Execute(transaction, data)
	if err != nil {
		transaction.RollbackTransaction()
		return nil, err, 500
	}

	err = transaction.CommitTransaction()
	if err != nil {
		return nil, err, 500
	}

	return response, nil, 200
}
