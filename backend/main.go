package main

import (
	"gin_backend/controller"
	"gin_backend/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Corsの設定
	middleware.Cors(router)

	// ルーティング
	controller.Router(router)

	// ポートの設定
	router.Run(":8080")
}
