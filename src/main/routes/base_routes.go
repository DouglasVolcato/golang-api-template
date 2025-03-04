package routes

import (
	"app/src/domain/usecases"
	"app/src/main/abstract"
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var BaseRoutes = []abstract.Route{
	{
		Name:         "Test Create",
		Path:         "/create",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.CreateUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.CreateUsecaseInput{},
		ResponseType: usecases.CreateUsecaseOutput{},
	},
	{
		Name:         "Test Update",
		Path:         "/update",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.UpdateUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.UpdateUsecaseInput{},
		ResponseType: usecases.UpdateUsecaseOutput{},
	},
	{
		Name:         "Test Delete",
		Path:         "/delete",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.DeleteUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.DeleteUsecaseInput{},
		ResponseType: usecases.DeleteUsecaseOutput{},
	},
	{
		Name:         "Test Get One",
		Path:         "/get-one",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.GetOneUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.GetOneUsecaseInput{},
		ResponseType: usecases.GetOneUsecaseOutput{},
	},
	{
		Name:         "Test Get All",
		Path:         "/get-all",
		Method:       abstract.RouteMethodTypes.POST,
		Controller:   controllers.NewController(&usecases.GetAllUsecase),
		Middlewares:  []*middlewares.Middleware{&middlewares.BaseMiddleware},
		RequestType:  usecases.GetAllUsecaseInput{},
		ResponseType: usecases.GetAllUsecaseOutput{},
	},
}
