package middlewares

import "app/src/domain/abstract"

type Middleware struct {
	Execute func(data abstract.DtoType) (abstract.DtoType, error)
}
