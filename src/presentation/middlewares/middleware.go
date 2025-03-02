package middlewares

type Middleware struct {
	Execute func(data any) (any, error)
}
