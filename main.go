package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/vaansh/gorest/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/items", handlers.GetItems)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
