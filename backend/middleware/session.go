package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// セッション認証
func AuthenticateBySession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessiona := sessions.Default(c)
		sessionId := sessiona.Get("SessionId")

		if sessionId == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
