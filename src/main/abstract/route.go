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

type Route struct {
	Name         string
	Path         string
	Method       string
	Controller   controllers.Controller
	Middlewares  []middlewares.Middleware
	TemplatePath string
	RequestType  interface{}
	ResponseType interface{}
}
