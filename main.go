package main

import (
	"net/http"

	"github.com/38tter/yutabe/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	user := router.Group("/User")
	{
		user.POST("/signup", routes.UserSignUp)
	}

	router.Run(":8080")
}
