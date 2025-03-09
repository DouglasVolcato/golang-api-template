package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var BaseRoutes = []abstract.Route{
	{
		Name:         "Create (JSON)",
		Path:         "/base/create",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(usecases.CreateUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		RequestType:  usecases.CreateUsecaseInput{},
		ResponseType: usecases.CreateUsecaseOutput{},
	},
	{
		Name:         "Update (JSON)",
		Path:         "/base/update",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(usecases.UpdateUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		RequestType:  usecases.UpdateUsecaseInput{},
		ResponseType: usecases.UpdateUsecaseOutput{},
	},
	{
		Name:         "Delete (JSON)",
		Path:         "/base/delete",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(usecases.DeleteUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		RequestType:  usecases.DeleteUsecaseInput{},
		ResponseType: usecases.DeleteUsecaseOutput{},
	},
	{
		Name:         "Get one (JSON)",
		Path:         "/base/get-one",
		Method:       abstract.RouteMethodTypes.GET,
		Controller:   controllers.NewController(usecases.GetOneUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		RequestType:  usecases.GetOneUsecaseInput{},
		ResponseType: usecases.GetOneUsecaseOutput{},
	},
	{
		Name:         "Get all (JSON)",
		Path:         "/base/get-all",
		Method:       abstract.RouteMethodTypes.GET,
		Controller:   controllers.NewController(usecases.GetAllUsecase),
		Middlewares:  []middlewares.Middleware{middlewares.BaseMiddleware},
		RequestType:  usecases.GetAllUsecaseInput{},
		ResponseType: usecases.GetAllUsecaseOutput{},
	},
}
