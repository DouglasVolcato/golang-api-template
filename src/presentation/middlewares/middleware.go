package middlewares

import "app/src/domain/abstract/dtos"

type Middleware struct {
	Execute func(data dtos.DtoType) (dtos.DtoType, error)
}
