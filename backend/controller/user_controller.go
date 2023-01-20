package controller

import (
	"gin_backend/model"
	"gin_backend/service"
	"gin_backend/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignUp(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	// ユーザー名重複チェック
	if service.IsUserExistByName(user.Name) {
		c.Status(http.StatusBadRequest)
		return
	}
	sessionId := uuid.New()
	user.SessionId = sessionId.String()
	newUser := service.CreateUser(user)
	// セッションに登録
	session.SetLoginUserToSession(c, newUser)
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	//ユーザー名とパスワードチェック
	if !service.IsUserExistForLogin(user) {
		c.Status(http.StatusBadRequest)
		return
	}
	sessionId := uuid.New()
	updateUser := service.UpdateSessionIdByName(sessionId.String(), user.Name)
	// セッションに登録
	session.SetLoginUserToSession(c, updateUser)
	c.Status(http.StatusOK)
}

// 以下Authユーザー

func Logout(c *gin.Context) {
	// セッションをクリア
	session.ClearSession(c)
	c.Status(http.StatusOK)
}

func GetMe(c *gin.Context) {
	loginUser := session.GetLoginUserFromSession(c)
	// セッションで管理しているユーザーをレスポンス型にキャスト
	var user model.LoginUserResponse
	user.Id = loginUser.UserId
	user.Name = loginUser.Name
	c.JSON(http.StatusOK, user)
}

func UpdateUserName(c *gin.Context) {
	log.Print("dsfjaklie")
	var updateUser model.User
	c.BindJSON(&updateUser)
	loginUser := session.GetLoginUserFromSession(c)
	if !service.UpdateUserName(loginUser, updateUser) {
		c.Status(http.StatusBadRequest)
		return
	}
	// セッションで管理しているユーザーを更新
	updateUser.Id = loginUser.UserId
	updateUser.SessionId = loginUser.SessionId
	session.SetLoginUserToSession(c, updateUser)
	c.Status(http.StatusOK)
}
