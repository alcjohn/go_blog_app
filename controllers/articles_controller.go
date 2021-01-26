package controllers

import (
	"net/http"
	"time"

	"github.com/alcjohn/go_fullstack/utils"

	"github.com/alcjohn/go_fullstack/config"
	"github.com/alcjohn/go_fullstack/models"

	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}
type ArticlesController struct{}

func (controller ArticlesController) InitRoutes(r *gin.RouterGroup) {
	r.GET("/", controller.Index)
	r.GET("/show/:article_id", controller.Show)
	r.GET("/new", controller.New)
	r.GET("/edit/:article_id", controller.Edit)
	r.POST("/new", controller.Create)
	r.POST("/edit/:article_id", controller.Update)
	r.POST("/delete/:article_id", controller.Destroy)
}

func (controller ArticlesController) Index(c *gin.Context) {
	var articles []models.Article
	user := c.MustGet("user").(models.User)
	config.DB.Where("user_id = ?", user.ID).Find(&articles)
	utils.Render(c, http.StatusOK, "articles/index", gin.H{
		"title":    "Liste des Articles",
		"articles": articles,
	})
}
func (controller ArticlesController) Show(c *gin.Context) {
	var article models.Article
	config.DB.Where("id = ?", c.Param("article_id")).First(&article)

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
	var input CreateInput
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
	config.DB.Create(&article)

	c.Redirect(http.StatusFound, "/me/articles")
}
func (controller ArticlesController) Update(c *gin.Context) {}
func (controller ArticlesController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from delete",
	})
}
