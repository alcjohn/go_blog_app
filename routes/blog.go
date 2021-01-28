package routes

import (
	"github.com/alcjohn/go_fullstack/controllers"
	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	var blogController controllers.BlogController
	blog := r.Group("/")
	{
		blog.GET("/", blogController.Index)
	}
	return blog
}
