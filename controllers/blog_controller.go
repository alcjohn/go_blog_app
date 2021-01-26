package controllers

import (
	"net/http"

	"github.com/alcjohn/go_fullstack/utils"

	"github.com/gin-gonic/gin"
)

type BlogController struct{}

func (controller BlogController) Index(c *gin.Context) {
	utils.Render(c, http.StatusOK, "blog/index", gin.H{
		"title": "Liste des articles de blog",
	})
}
