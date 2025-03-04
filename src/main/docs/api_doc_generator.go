package docs

import (
	"app/src/main/abstract"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

type ApiDocGenerator struct {
	router *gin.Engine
}

func NewApiDocGenerator(router *gin.Engine) *ApiDocGenerator {
	return &ApiDocGenerator{router: router}
}

func (s *ApiDocGenerator) RegisterRoutes(routes []abstract.Route) {
	endpoints := []*swagger.Endpoint{}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	handler := func(context *gin.Context) {
	}

	for _, route := range routes {
		newRoute := endpoint.New(
			route.Method,
			route.Path,
			route.Name,
			endpoint.Handler(handler),
			endpoint.Body(route.RequestType, "Request Payload", true),
			endpoint.Response(200, route.ResponseType, "Default Response"),
			endpoint.Response(400, ErrorResponse{}, "Request error"),
			endpoint.Response(403, ErrorResponse{}, "Authentication error"),
			endpoint.Response(500, ErrorResponse{}, "Server error"),
		)

		endpoints = append(endpoints, newRoute)
	}

	endpoints = s.combine(endpoints)

	api := swag.New(
		swag.Endpoints(endpoints...),
		swag.Description("This is the api"),
		swag.Version("1.0.0"),
		swag.Title("Api"),
	)

	enableCors := true
	s.router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))
	s.router.GET("/docs", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		scheme := "http://"
		if ctx.Request.TLS != nil {
			scheme = "https://"
		}
		content := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Scalar API Reference</title>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
		</head>
		<body>
			<script
			id="api-reference"
			type="application/json"
			data-url="%s"
			></script>
			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		</body>
		</html>
		`, scheme+ctx.Request.Host+"/swagger")
		ctx.String(http.StatusOK, content)
	})
}

func (s *ApiDocGenerator) combine(endpoints []*swagger.Endpoint) []*swagger.Endpoint {
	return append(endpoints, endpoints...)
}
