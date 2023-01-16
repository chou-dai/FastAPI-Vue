package session

import (
	"gin_backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetLoginUserToSession(c *gin.Context, user model.User) {
	session := sessions.Default(c)
	session.Set("UserId", user.Id)
	session.Set("Name", user.Name)
	session.Set("SessionId", user.SessionId)
	session.Save()
}

func GetLoginUserFromSession(c *gin.Context) model.SessionInfo {
	var loginUser model.SessionInfo
	session := sessions.Default(c)
	loginUser.UserId = session.Get("UserId").(int)
	loginUser.Name = session.Get("Name").(string)
	loginUser.SessionId = session.Get("SessionId").(string)
	return loginUser
}

func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
