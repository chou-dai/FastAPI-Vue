package controller

import (
	"gin_backend/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	api := router.Group("/api")
	{
		// ログインしていないユーザー
		api.POST("/signup", SignUp)
		api.POST("/login", Login)

		// ログインユーザー
		auth := api.Group("/auth", middleware.AuthenticateBySession())
		{
			auth.POST("/logout", Logout)

			users := auth.Group("/users")
			{
				users.GET("")
				users.POST("")
				users.PUT("")
				users.DELETE("")
			}

			memories := auth.Group("/memories")
			{
				memories.GET("", GetAllMemories)
				memories.POST("", CreateMemory)
				memories.PUT("/:id", UpdateMemory)
				memories.DELETE("/:id", DeleteMemory)
			}
		}
	}
}
