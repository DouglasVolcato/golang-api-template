package docs

import (
	"app/src/main/abstract"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

type ApiDocGenerator struct {
	name        string
	description string
	router      *gin.Engine
}

func NewApiDocGenerator(name string, description string, router *gin.Engine) *ApiDocGenerator {
	return &ApiDocGenerator{name: name, description: description, router: router}
}

func (s *ApiDocGenerator) RegisterRoutes(routes []abstract.Route) {
	endpoints := []*swagger.Endpoint{}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	handler := func(context *gin.Context) {
	}

	for _, route := range routes {
		var newRoute *swagger.Endpoint

		var requestData []endpoint.Option

		if route.Method == "GET" {
			queryParams := make([]endpoint.Option, 0)
			if structs.IsStruct(route.RequestType) {
				for _, field := range structs.Fields(route.RequestType) {
					queryParams = append(queryParams, endpoint.Query(
						strings.ToLower(field.Name()),
						fmt.Sprintf("%s %s", strings.ToLower(field.Name()), field.Kind().String()),
						"query",
						true,
					))
				}
			}
			requestData = queryParams
		} else {
			requestData = []endpoint.Option{endpoint.Body(route.RequestType, "Request Payload", true)}
		}

		commonParams := []endpoint.Option{
			endpoint.Handler(handler),
		}

		if route.TemplatePath != "" {
			commonParams = append(commonParams,
				endpoint.Response(200, "text/html", "Default Response - HTML template"),
				endpoint.Response(400, "text/html", "Request error - HTML template"),
				endpoint.Response(403, "text/html", "Authentication error - HTML template"),
				endpoint.Response(500, "text/html", "Server error - HTML template"),
			)
		} else {
			commonParams = append(commonParams,
				endpoint.Response(200, route.ResponseType, "Default Response"),
				endpoint.Response(400, ErrorResponse{}, "Request error"),
				endpoint.Response(403, ErrorResponse{}, "Authentication error"),
				endpoint.Response(500, ErrorResponse{}, "Server error"),
			)
		}

		// Append request data and create the new route
		commonParams = append(commonParams, requestData...)

		newRoute = endpoint.New(route.Method, route.Path, route.Name, commonParams...)

		endpoints = append(endpoints, newRoute)
	}

	endpoints = s.combine(endpoints)

	api := swag.New(
		swag.Endpoints(endpoints...),
		swag.Description(s.description),
		swag.Title(s.name),
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
			<title>Docs - `+s.name+`</title>
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
