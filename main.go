package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("PORT", "9090")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
