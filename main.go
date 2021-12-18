package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "Wecome to Tiny URL Service",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}