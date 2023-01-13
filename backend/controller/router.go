package controller

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", GetAllUsers)
			users.POST("/", CreateUser)
			users.PUT("/")
			users.DELETE("/")
		}

		memories := api.Group("/memories")
		{
			memories.GET("/")
			memories.POST("/")
			memories.PUT("/")
			memories.DELETE("/")
		}
	}
}
