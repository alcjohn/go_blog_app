package main

import (
	"html/template"
	"log"

	"github.com/alcjohn/go_fullstack/middlewares"

	"github.com/alcjohn/go_fullstack/config"
	"github.com/alcjohn/go_fullstack/html_functions"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"

	method "github.com/bu/gin-method-override"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
)

var r *gin.Engine

var posts = []string{
	"First post",
	"Second Post",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	binding.Validator = new(config.DefaultValidator)
	r = gin.Default()
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(method.ProcessMethodOverride(r))
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
	// r.LoadHTMLGlob("views/**/*")
	initializeRoutes()
	r.Run(":8081")

}
