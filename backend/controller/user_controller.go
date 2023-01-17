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
	if service.IsUserExist(user) {
		c.Status(http.StatusBadRequest)
		return
	}

	sessionId := uuid.New()
	user.SessionId = sessionId.String()
	newUser := service.CreateUser(user)
	// セッションに登録
	session.SetLoginUserToSession(c, newUser)
	c.Status(http.StatusOK)
	log.Print(session.GetLoginUserFromSession(c).Name)
}

func Login(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	if !service.IsUserExist(user) {
		c.Status(http.StatusBadRequest)
		return
	}

	sessionId := uuid.New()
	updateUser := service.UpdateSessionIdByName(sessionId.String(), user.Name)
	// セッションに登録
	session.SetLoginUserToSession(c, updateUser)
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	// セッションをクリア
	session.ClearSession(c)
	c.Status(http.StatusOK)
}

func GetMe(c *gin.Context) {
	loginUser := session.GetLoginUserFromSession(c)

	var user model.LoginUserResponse
	user.Id = loginUser.UserId
	user.Name = loginUser.Name

	c.JSON(http.StatusOK, user)
}
