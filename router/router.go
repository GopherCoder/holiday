package router

import (
	"holiday/handler/holiday"
	"holiday/handler/user"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) InitRouter(g *gin.Engine, handlerFunctions ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(handlerFunctions...)

	holidayResource := g.Group("/v1/api/holiday")
	{
		holiday.RegisterHoliday(holidayResource)
	}

	userResource := g.Group("/v1/api/user")
	{
		user.RegisterUser(userResource)
	}

	return g
}
