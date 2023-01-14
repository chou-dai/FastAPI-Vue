package controller

import (
	"github.com/gin-gonic/gin"
)

/*
エンドポイントの設定を行います
*/
func Router(router *gin.Engine) {
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", GetAllUsers)
			users.POST("", CreateUser)
			users.PUT("")
			users.DELETE("")
		}

		memories := api.Group("/memories")
		{
			memories.GET("", GetAllMemories)
			memories.POST("", CreateMemory)
			memories.PUT("/:id", UpdateMemory)
			memories.DELETE("/:id", DeleteMemory)
		}
	}
}
