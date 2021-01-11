package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	username := ctx.PostForm("username")
	println("username: " + username)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}
