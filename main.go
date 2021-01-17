package main

import (
	"fmt"

	"github.com/38tter/yutabe/routes"
	"github.com/38tter/yutabe/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/products", usecase.GetProducts)

	user := router.Group("/User")
	{
		user.POST("/signup", routes.UserSignUp)
	}
	fmt.Printf("connect to db")

	router.Run(":8081")
}
