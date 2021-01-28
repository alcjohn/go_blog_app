package config

import (
	"html/template"
	"log"

	"github.com/alcjohn/go_fullstack/html_functions"
	"github.com/alcjohn/go_fullstack/middlewares"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func NewApp() *gin.Engine {
	r := gin.Default()
	binding.Validator = new(DefaultValidator)
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}
	r.Use(sessions.Sessions("user_session", store))
	r.Use(middlewares.AuthMiddleware())
	r.HTMLRender = ginview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/base",
		Funcs: template.FuncMap{
			"script": html_functions.ScriptTag,
			"link":   html_functions.LinkTag,
			"vitejs": html_functions.Vitejs,
		},
		DisableCache: true,
	})
	r.Static("/public", "./public")
	return r
}
