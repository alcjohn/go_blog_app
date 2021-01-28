package config

import (
	"github.com/alcjohn/go_fullstack/routes"
	"github.com/gin-gonic/gin"
	"github.com/goava/di"
)

type Routes struct{}

var RoutesProviders = di.Options(
	di.Provide(NewRoutes),
	di.Provide(routes.ArticlesRoutes),
	di.Provide(routes.MeRoutes),
)

func NewRoutes(
	r *gin.Engine,
) *Routes {
	router := r.Group("/")
	routes.BlogRoutes(router)
	me := routes.MeRoutes(router)
	{
		routes.ArticlesRoutes(me)
	}
	routes.AuthRoutes(router)

	return &Routes{}

}
