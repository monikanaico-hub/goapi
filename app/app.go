package app

import (
	"github.com/gin-gonic/gin"
	"github.com/monikanaico-hub/goapi/handler"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	// Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.POST("/users", handler.CreateUserHandler)
		v1.GET("/users", handler.GetallUserHandler)
	}
}
