package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var BaseRoutes = []abstract.Route[usecases.BaseUsecaseInput, usecases.BaseUsecaseOutput]{
	{
		Path:        "/",
		Method:      abstract.RouteMethodTypes.GET,
		Controller:  controllers.NewController(&usecases.BaseUsecase),
		Middlewares: []*middlewares.Middleware{&middlewares.BaseMiddleware},
	},
}
