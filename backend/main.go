package main

import (
	"gin_backend/controller"
	"gin_backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Corsの設定
	middleware.Cors(router)

	// ルーティング
	controller.Router(router)

	// ポートの設定
	router.Run(":8080")
}
