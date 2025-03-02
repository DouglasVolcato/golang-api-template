package middlewares

var BaseMiddleware = Middleware{
	Execute: func(data any) (any, error) {
		return data, nil
	},
}
