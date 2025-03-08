package main

import (
	"fmt"
	"strings"

	"app/src/domain/abstract/dtos"
	"app/src/infra/database"
	"app/src/main/docs"
	"app/src/main/routes"

	"github.com/gin-gonic/gin"
)

func formatMessage(message string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(message[:1]), message[1:])
}

func main() {
	var router = gin.Default()
	router.LoadHTMLGlob("src/presentation/templates/**/*")

	var databaseConnection = database.InitializeDatabaseConnection()

	err := database.ExecuteDatabaseMigrations(databaseConnection)
	if err != nil {
		panic(err)
	}

	var routes = append(routes.BaseRoutes, routes.TemplateRoutes...)

	docGenerator := docs.NewApiDocGenerator("Api", "Api description", router)
	docGenerator.RegisterRoutes(routes)

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Api is running",
		})
	})

	for _, route := range routes {
		var path = route.Path
		var method = route.Method
		var controller = route.Controller
		var middlewares = route.Middlewares

		router.Handle(method, path, func(context *gin.Context) {
			var data = dtos.DtoType{}
			var err error

			for key, value := range context.Request.Header {
				data[key] = value
			}

			for key, value := range context.Request.URL.Query() {
				data[key] = value[0]
			}

			if method != "GET" {
				body := dtos.DtoType{}
				if context.BindJSON(&body) == nil {
					for key, value := range body {
						data[key] = value
					}
				}
			}

			for _, middleware := range middlewares {
				data, err = middleware.Execute(data)
				if err != nil {
					context.JSON(400, gin.H{
						"error": formatMessage(err.Error()),
					})
					return
				}
			}

			response, err, status := controller.Execute(databaseConnection, data)

			if route.TemplatePath != "" {
				if err != nil {
					context.HTML(status, "index.html", err)
					return
				}
				context.HTML(status, route.TemplatePath, response)
				return
			} else {
				if err != nil {
					context.JSON(status, gin.H{
						"error": formatMessage(err.Error()),
					})
					return
				}
				context.JSON(status, response)
				return
			}
		})
	}

	router.Run()
}
