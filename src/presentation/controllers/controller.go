package controllers

import (
	"app/src/infra/database"
	"database/sql"
)

type Controller struct {
	usecaseFunc func(data interface{}) (interface{}, error)
}

func NewController(usecaseFunc func(data interface{}) (interface{}, error)) *Controller {
	return &Controller{usecaseFunc: usecaseFunc}
}

func (controller *Controller) Handle(databaseConnection *sql.DB, data interface{}) interface{} {
	var transaction = database.NewTransaction(databaseConnection)
	var err error

	err = transaction.BeginTransaction()
	if err != nil {
		return struct {
			error string
		}{error: "Falha ao iniciar transação"}
	}

	response, err := controller.usecaseFunc(data)

	if err != nil {
		transaction.RollbackTransaction()
		return struct {
			error string
		}{error: err.Error()}
	}

	err = transaction.CommitTransaction()
	if err != nil {
		return struct {
			error string
		}{error: "Falha ao finalizar transação"}
	}

	return response
}
