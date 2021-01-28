package main

import (
	"log"

	"github.com/alcjohn/go_fullstack/global"

	"github.com/goava/di"

	"github.com/alcjohn/go_fullstack/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var posts = []string{
	"First post",
	"Second Post",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	global.DB = config.ConnectDB()
	c, err := di.New(
		di.Provide(config.NewApp),
		di.Provide(config.NewRoutes),
	)
	if err := c.Invoke(start); err != nil {
		log.Fatal(err)
	}

}
func start(r *gin.Engine, _ *config.Routes) {
	defer func() {
		global.DB.Close()
	}()
	r.Run(":8081")
}
