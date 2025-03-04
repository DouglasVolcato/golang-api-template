package middlewares

import "app/src/domain/abstract/dtos"

var BaseMiddleware = Middleware{
	Execute: func(data dtos.DtoType) (dtos.DtoType, error) {
		data["token"] = "fb7362fg28346fg234"
		return data, nil
	},
}
