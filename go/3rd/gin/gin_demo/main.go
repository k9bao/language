package main

import (
	"github.com/gin-gonic/gin"
)

func main_simple() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	// main_simple()
	main_method()
}
