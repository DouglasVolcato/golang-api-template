package abstract

import (
	"app/src/presentation/controllers"
	"app/src/presentation/middlewares"
)

var RouteMethodTypes = struct {
	POST   string
	GET    string
	PUT    string
	DELETE string
	PATCH  string
}{
	POST:   "POST",
	GET:    "GET",
	PUT:    "PUT",
	DELETE: "DELETE",
	PATCH:  "PATCH",
}

type Route[I any, O any] struct {
	Path        string
	Method      string
	Controller  *controllers.Controller[I, O]
	Middlewares []*middlewares.Middleware
}
