package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var TemplateRoutes = []abstract.Route{
	{
		Name:         "Get All with template",
		Path:         "template/get-all",
		Method:       abstract.RouteMethodTypes.GET,
		Controller:   controllers.NewController(usecases.GetAllUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		TemplatePath: "index.tmpl",
		RequestType:  usecases.GetAllUsecaseInput{},
		ResponseType: usecases.GetAllUsecaseOutput{},
	},
}
