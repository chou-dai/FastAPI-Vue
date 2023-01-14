package controller

// ユーザーのバックエンド処理のエントリーポイント
// バリデーション + service層へのパイプライン

import (
	"gin_backend/model"
	"gin_backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	service.CreateUser(user)
	c.Status(http.StatusCreated)
}
