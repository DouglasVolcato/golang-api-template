package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var BaseRoutes = []abstract.Route{
	{
		Path:         "/test",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.BaseUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.BaseUsecaseInput{},
		ResponseType: usecases.BaseUsecaseOutput{},
	},
}
