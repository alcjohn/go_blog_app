package routes

import (
	"github.com/alcjohn/go_fullstack/controllers"
	"github.com/alcjohn/go_fullstack/middlewares"
	"github.com/gin-gonic/gin"
)

func MeRoutes(
	r *gin.RouterGroup,
) *gin.RouterGroup {
	var meController controllers.MeController
	me := r.Group("me")
	me.Use(middlewares.MustBeAuth())
	{
		me.GET("/", meController.Index)
	}
	return me
}
