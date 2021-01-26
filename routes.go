package main

import (
	"github.com/alcjohn/go_fullstack/middlewares"

	"github.com/alcjohn/go_fullstack/controllers"
)

var articlesController controllers.ArticlesController
var blogController controllers.BlogController
var authController controllers.AuthController
var meController controllers.MeController

func initializeRoutes() {

	r.GET("/", blogController.Index)
	me := r.Group("/me", middlewares.MustBeAuth())
	{
		me.GET("/", meController.Index)
		articles := me.Group("/articles")
		articlesController.InitRoutes(articles)
	}
	r.GET("/signin", authController.SignIn)
	r.GET("/signup", authController.SignUp)
	r.POST("/signin", authController.Login)
	r.POST("/signup", authController.Register)
	r.POST("/logout", authController.Logout)

}
