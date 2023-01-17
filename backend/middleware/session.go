package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	session := sessions.Default(c)
	sessionId := session.Get("SessionId")

	if sessionId == nil {
		c.Status(http.StatusUnauthorized)
	} else {
		c.Status(http.StatusOK)
	}
}

// セッション認証
func AuthenticateBySession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get("SessionId")

		if sessionId == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
