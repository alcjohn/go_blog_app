package controllers

import (
	"net/http"
	"time"

	"github.com/alcjohn/go_fullstack/forms"

	"github.com/alcjohn/go_fullstack/global"
	"github.com/alcjohn/go_fullstack/utils"

	"github.com/alcjohn/go_fullstack/models"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
}

func (controller ArticlesController) Index(c *gin.Context) {
	var articles []models.Article
	// user := c.MustGet("user").(models.User)
	global.DB.Where("user_id = ?", 1).Find(&articles)
	utils.Render(c, http.StatusOK, "articles/index", gin.H{
		"title":    "Liste des Articles",
		"articles": articles,
	})
}
func (controller ArticlesController) Show(c *gin.Context) {
	var article models.Article
	global.DB.Where("id = ?", c.Param("article_id")).First(&article)

	utils.Render(c, http.StatusOK, "articles/show", gin.H{
		"title":   article.Title,
		"article": article,
	})
}
func (controller ArticlesController) New(c *gin.Context) {
	utils.Render(c, http.StatusOK, "articles/new", gin.H{
		"title": "Créer un Article",
	})
}
func (controller ArticlesController) Edit(c *gin.Context) {}

func (controller ArticlesController) Create(c *gin.Context) {
	var input forms.ArticleCreate
	user := c.MustGet("user").(models.User)
	if err := c.ShouldBind(&input); err != nil {
		utils.Render(c, http.StatusBadRequest, "articles/new", gin.H{
			"title":  "Créer un Article",
			"errors": err,
			"input":  input,
		})
	}
	article := models.Article{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		User:        user,
		PublishedAt: time.Now(),
		IsPublished: true,
	}
	global.DB.Create(&article)

	c.Redirect(http.StatusFound, "/me/articles")
}
func (controller *ArticlesController) Update(c *gin.Context) {}
func (controller *ArticlesController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from delete",
	})
}
