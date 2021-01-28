package controllers

import (
	"net/http"

	"github.com/alcjohn/go_fullstack/global"
	"github.com/alcjohn/go_fullstack/models"

	"github.com/alcjohn/go_fullstack/utils"

	"github.com/gin-gonic/gin"
)

type BlogController struct{}

func (controller BlogController) Index(c *gin.Context) {
	var articles []models.Article
	global.DB.Where("is_published = true").Find(&articles)
	utils.Render(c, http.StatusOK, "blog/index", gin.H{
		"title":    "Liste des articles de blog",
		"articles": articles,
	})
}
