package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("PORT", "9090")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api is running",
		})
	})
	r.Run()
}
