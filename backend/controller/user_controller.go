package controller

import (
	exception "gin_backend/exceptions"
	"gin_backend/model"
	"gin_backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
全てのユーザーを取得します

	Returns
	-------
		status: 200
		users: 全ユーザー
*/
func GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

/*
ユーザーを作成します

	Parameters
	----------
		user: UserRequestModel

	Returns
	-------
		status: 201

	Exceptions
	----------
		status: 400 リクエストエラーs
*/
func CreateUser(c *gin.Context) {
	var user model.UserRequestModel
	err := c.BindJSON(&user)

	// 400: バリデーションエラー
	if err != nil {
		exception.BadRequest(c)
	}

	service.CreateUser(user)
	c.Status(http.StatusCreated)
}
