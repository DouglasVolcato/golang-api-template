package middlewares

import "app/src/domain/abstract"

var BaseMiddleware = Middleware{
	Execute: func(data abstract.DtoType) (abstract.DtoType, error) {
		data["token"] = "fb7362fg28346fg234"
		return data, nil
	},
}
