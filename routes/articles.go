package routes

import (
	"github.com/alcjohn/go_fullstack/controllers"
	"github.com/gin-gonic/gin"
)

func ArticlesRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	var articlesController controllers.ArticlesController

	articles := r.Group("/articles")
	{
		articles.GET("/", articlesController.Index)
		articles.GET("/show/:article_id", articlesController.Show)
		articles.GET("/new", articlesController.New)
		articles.GET("/edit/:article_id", articlesController.Edit)
		articles.POST("/new", articlesController.Create)
		articles.POST("/edit/:article_id", articlesController.Update)
		articles.POST("delete/:article_id", articlesController.Destroy)
	}
	return articles
}
