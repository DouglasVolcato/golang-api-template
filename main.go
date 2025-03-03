package main

import (
	"fmt"
	"os"
	"strings"

	"app/src/domain/abstract"
	"app/src/infra/database"
	"app/src/main/routes"

	"github.com/gin-gonic/gin"
)

func formatMessage(message string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(message[:1]), message[1:])
}

func main() {
	os.Setenv("PORT", "9090")
	var router = gin.Default()

	var databaseConnection = database.InitializeDatabaseConnection()

	router.GET("/", func(context *gin.Context) {
		fmt.Print(routes.BaseRoutes)
		context.JSON(200, gin.H{
			"message": "Api is running",
		})
	})

	for _, route := range routes.BaseRoutes {

		fmt.Println(route)

		var path = route.Path
		var method = route.Method
		var controller = route.Controller
		var middlewares = route.Middlewares

		router.Handle(method, path, func(context *gin.Context) {
			var data = abstract.DtoType{}
			var err error

			for key, value := range context.Request.Header {
				data[key] = value
			}

			for key, value := range context.Request.URL.Query() {
				data[key] = value
			}

			body := abstract.DtoType{}
			if err := context.BindJSON(&body); err == nil {
				for key, value := range body {
					data[key] = value
				}
			}

			for _, middleware := range middlewares {
				data, err = middleware.Execute(data)
				if err != nil {
					context.JSON(400, gin.H{
						"error": formatMessage(err.Error()),
					})
				}
			}

			response, err, status := controller.Execute(databaseConnection, data)

			if err != nil {
				context.JSON(status, gin.H{
					"error": formatMessage(err.Error()),
				})
				return
			}

			context.JSON(status, response)
		})
	}

	router.Run()
}
