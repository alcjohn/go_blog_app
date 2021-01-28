package routes

import (
	"github.com/alcjohn/go_fullstack/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	var authController controllers.AuthController

	auth := r.Group("/")
	{
		auth.GET("/signin", authController.SignIn)
		auth.GET("/signup", authController.SignUp)
		auth.POST("/signin", authController.Login)
		auth.POST("/signup", authController.Register)
		auth.POST("/logout", authController.Logout)
	}
	return auth
}
