package main

import (
	"gin_backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routerをcontrollerに渡してルーティング
	controller.Router(router)

	// ポートの設定
	router.Run(":8080")
}
