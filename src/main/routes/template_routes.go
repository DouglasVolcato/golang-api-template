package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var TemplateRoutes = []abstract.Route{
	{
		Name:         "Get all (HTML)",
		Path:         "/html/get-all",
		Method:       abstract.RouteMethodTypes.GET,
		Controller:   controllers.NewController(usecases.GetAllUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		TemplatePath: "index.html",
		RequestType:  usecases.GetAllUsecaseInput{},
		ResponseType: usecases.GetAllUsecaseOutput{},
	},
}
