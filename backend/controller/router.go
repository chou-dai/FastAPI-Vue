package controller

import (
	"gin_backend/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	api := router.Group("/api")
	{
		// 認証チェック
		api.GET("/auth", middleware.IsAuth)

		// ログインしていないユーザー
		api.POST("/signup", SignUp)
		api.POST("/login", Login)
		api.GET("/memories", GetPublicMemories)

		// ログインユーザー
		auth := api.Group("/auth", middleware.AuthenticateBySession())
		{
			auth.POST("/logout", Logout)
			users := auth.Group("/users")
			{
				users.GET("", GetMe)
				users.PUT("")
			}
			memories := auth.Group("/memories")
			{
				memories.GET("", GetMyMemories)
				memories.POST("", CreateMemory)
				memories.PUT("/:id", UpdateMemory)
				memories.DELETE("/:id", DeleteMemory)
			}
		}
	}
}
