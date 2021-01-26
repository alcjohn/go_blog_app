package controllers

import (
	"net/http"

	"github.com/alcjohn/go_fullstack/utils"
	"github.com/gin-gonic/gin"
)

type MeController struct{}

func (controller MeController) Index(c *gin.Context) {
	utils.Render(c, http.StatusOK, "admin/index", gin.H{
		"title": "Admin Page",
	})
}
