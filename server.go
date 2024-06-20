package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start() {
	fmt.Println("Server started")
	r := gin.Default()

	counter := 0

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": counter,
			"content": "hello",
		})
		counter++
	})
	r.Run(":8080")
}